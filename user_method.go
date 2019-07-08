package usersvr

import (
	"context"
	"fmt"
	"log"
	"net/http"
	usermethod "usersvr/gen/user_method"
	"usersvr/utils/db/pg"
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
func (s *userMethodsrvc) ChangeInfo(ctx context.Context, p *usermethod.User) (res string, err error) {
	s.logger.Print("userMethod.changeInfo")
	return
}

// Remove bottle from storage
func (s *userMethodsrvc) ChangePassword(ctx context.Context, p *usermethod.ChangePasswordPayload) (err error) {
	s.logger.Print("userMethod.changePassword")
	return
}

// Rate bottles by IDs
func (s *userMethodsrvc) ForgotPassword(ctx context.Context, p *usermethod.ForgotPasswordPayload) (err error) {
	s.logger.Print("userMethod.forgotPassword")
	return
}

// Add n number of bottles and return their IDs. This is a multipart request
// and each part has field name 'bottle' and contains the encoded bottle info
// to be added.
func (s *userMethodsrvc) ChangeEmail(ctx context.Context, p *usermethod.ChangeEmailPayload) (res string, err error) {
	s.logger.Print("userMethod.changeEmail")
	return
}
