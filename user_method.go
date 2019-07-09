package usersvr

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"
	usermethod "usersvr/gen/user_method"
	"usersvr/utils"
	"usersvr/utils/db/pg"
	"usersvr/utils/db/redis"
	email "usersvr/utils/email"
)

var (
	ErrPassword error = usermethod.PasswordError("password is error!!!")
)

// userMethod service example implementation.
// The example methods log the requests and return zero values.
type userMethodsrvc struct {
	logger *log.Logger
}

// NewUserMethod returns the userMethod service implementation.
func NewUserMethod(logger *log.Logger) usermethod.Service {
	return &userMethodsrvc{logger}
}

// register user
func (s *userMethodsrvc) Register(ctx context.Context, p *usermethod.RegisterPayload) (res *usermethod.ResponseResult, err error) {
	user, err :=pg.Register(p.Email, p.Password)
	if err != nil {
		return &usermethod.ResponseResult{Code:http.StatusBadRequest}, err
	}
	resStr := fmt.Sprintf("succcess create %s", user.Email)
	return &usermethod.ResponseResult{Code:http.StatusOK, Message: &resStr}, nil
}

// Show user info
func (s *userMethodsrvc) Show(ctx context.Context, p *usermethod.ShowPayload) (res *usermethod.UserInfo, view string, err error) {
	user, err := TokenToUser(p.Token)
	if err != nil {
		return &usermethod.UserInfo{}, "default", ErrInvalidToken
	}
	res = &usermethod.UserInfo{
		ID: int(user.ID),
		Name:user.Name,
		Email:user.Email,
		Icon:user.Icon,
	}
	return res, "default", nil
}

// Creates a valid JWT
func (s *userMethodsrvc) Login(ctx context.Context, p *usermethod.LoginPayload) (res *usermethod.Creds, err error) {
	user, err := pg.FindOneUser(p.Email)
	if err != nil {
		return &usermethod.Creds{}, ErrPassword
	}
	if p.Password != user.Password {
		return &usermethod.Creds{}, ErrPassword
	}

	t, err := GenToken(p.Email)
	if err != nil {
		return nil, err
	}
	return &usermethod.Creds{
		JWT:        t,
	}, nil
}

// This action is secured with the jwt scheme
//func (s *userMethodsrvc) Secure(ctx context.Context, p *usermethod.SecurePayload) (res string, err error) {
//	s.logger.Print("userMethod.secure")
//	return
//}

// Add new bottle and return its ID.
func (s *userMethodsrvc) ChangeInfo(ctx context.Context, p *usermethod.ChangeInfoPayload) (res *usermethod.UserInfo, view string, err error) {
	user, err := TokenToUser(p.Token)
	if err != nil {
		return &usermethod.UserInfo{}, "default", ErrInvalidToken
	}
	updataInfo := map[string]interface{}{
		"name": p.Name,
		"icon":p.Icon,
	}
	resUser, err := pg.UpdateUser(user, updataInfo)
	if err != nil {
		return &usermethod.UserInfo{}, "default", ErrInvalidToken
	}
	res = &usermethod.UserInfo{
		ID: int(resUser.ID),
		Name:resUser.Name,
		Email:resUser.Email,
		Icon:resUser.Icon,
	}
	return res, "changeInfo", nil
}

// Remove bottle from storage
func (s *userMethodsrvc) ChangePassword(ctx context.Context, p *usermethod.ChangePasswordPayload) (res *usermethod.ResponseResult, err error) {
	user, err := TokenToUser(p.Token)
	if err != nil {
		return &usermethod.ResponseResult{}, ErrInvalidToken
	}
	if p.OldPassword != user.Password {
		return &usermethod.ResponseResult{}, ErrPassword
	}
	updataInfo := map[string]interface{}{
		"password": p.NewPassword,
	}
	_, err = pg.UpdateUser(user, updataInfo)
	if err != nil {
		return &usermethod.ResponseResult{}, err
	}
	resMsg := "success to change password"
	res = &usermethod.ResponseResult{
		Code:http.StatusOK,
		Message: &resMsg,
	}
	return res,nil

}

// Rate bottles by IDs
func (s *userMethodsrvc) ForgotPassword(ctx context.Context, p *usermethod.ForgotPasswordPayload) (res *usermethod.ResponseResult, err error) {
	code := p.Code
	userEmail := p.Email
	user, err:= pg.FindOneUser(userEmail)
	if err != nil {
		return &usermethod.ResponseResult{}, err
	}
	redisKye := fmt.Sprintf("forgotPass%s", userEmail)
	redisCode, err := redis.RedisClient.Get(redisKye).Result()
	if err != nil {
		msg := "verify code is error"
		res := &usermethod.ResponseResult{
			Code:http.StatusBadRequest,
			Message: &msg,
		}
		return res, nil
	}
	if code != redisCode {
		msg := "verify code is error"
		res := &usermethod.ResponseResult{
			Code:http.StatusBadRequest,
			Message: &msg,
		}
		return res, nil
	}
	updateInfo := map[string]interface{}{
		"password":p.NewPassword,
	}
	_, err = pg.UpdateUser(user, updateInfo)
	if err != nil {
		return &usermethod.ResponseResult{}, err
	}
	resMsg := "success to change password"
	res = &usermethod.ResponseResult{
		Code:http.StatusOK,
		Message: &resMsg,
	}
	return res,nil
}

// Add n number of bottles and return their IDs. This is a multipart request
// and each part has field name 'bottle' and contains the encoded bottle info
// to be added.
func (s *userMethodsrvc) ChangeEmail(ctx context.Context, p *usermethod.ChangeEmailPayload) (res *usermethod.ResponseResult, err error) {
	s.logger.Print("userMethod.changeEmail")
	return
}

func (s *userMethodsrvc) SendVerifyCode(ctx context.Context, p *usermethod.SendVerifyCodePayload) (res *usermethod.ResponseResult, err error) {
	userEmail := p.Email
	_, err = pg.FindOneUser(userEmail)
	if err != nil {
		msg := "email not exist"
		res := &usermethod.ResponseResult{
			Code:http.StatusBadRequest,
			Message: &msg,
		}
		return res, nil
	}
	verifyCode := utils.GenVerifyCode()
	redisKye := fmt.Sprintf("forgotPass%s", userEmail)
	err = redis.RedisClient.Set(redisKye, verifyCode, time.Minute*5).Err()
	if err != nil {
		msg := fmt.Sprintf("fail to send verify code to your email :%s", userEmail)
		res := &usermethod.ResponseResult{
			Code: http.StatusBadRequest,
			Message: &msg,
		}
		return res, nil
	}
	sendTo := []string{userEmail}
	subject := "verify code for forgot password"
	body := fmt.Sprintf("<html><body><h3><p>您的验证码为：</p><p>%s</p></h3></body></html>", verifyCode)
	err = email.SendToMail(sendTo, subject, body)
	if err != nil {
		msg := fmt.Sprintf("fail to send verify code to your email :%s", userEmail)
		res := &usermethod.ResponseResult{
			Code: http.StatusBadRequest,
			Message: &msg,
		}
		return res, nil
	}
	msg := fmt.Sprintf("success to send verify code to your email :%s", userEmail)
	res = &usermethod.ResponseResult{
		Code: http.StatusOK,
		Message: &msg,
	}
	return res, nil
}
