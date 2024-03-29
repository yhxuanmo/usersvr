// Code generated by goa v3.0.2, DO NOT EDIT.
//
// userMethod gRPC client types
//
// Command:
// $ goa gen usersvr/design

package client

import (
	"unicode/utf8"
	user_methodpb "usersvr/gen/grpc/user_method/pb"
	usermethod "usersvr/gen/user_method"
	usermethodviews "usersvr/gen/user_method/views"

	goa "goa.design/goa/v3/pkg"
)

// NewRegisterRequest builds the gRPC request type from the payload of the
// "register" endpoint of the "userMethod" service.
func NewRegisterRequest(payload *usermethod.RegisterPayload) *user_methodpb.RegisterRequest {
	message := &user_methodpb.RegisterRequest{
		Email:    payload.Email,
		Password: payload.Password,
	}
	return message
}

// NewRegisterResult builds the result type of the "register" endpoint of the
// "userMethod" service from the gRPC response type.
func NewRegisterResult(message *user_methodpb.RegisterResponse) *usermethod.ResponseResult {
	result := &usermethod.ResponseResult{
		Code: int(message.Code),
	}
	if message.Message_ != "" {
		result.Message = &message.Message_
	}
	if message.Data != nil {
		result.Data = make(map[string]string, len(message.Data))
		for key, val := range message.Data {
			tk := key
			tv := val
			result.Data[tk] = tv
		}
	}
	return result
}

// NewShowRequest builds the gRPC request type from the payload of the "show"
// endpoint of the "userMethod" service.
func NewShowRequest() *user_methodpb.ShowRequest {
	message := &user_methodpb.ShowRequest{}
	return message
}

// NewShowResult builds the result type of the "show" endpoint of the
// "userMethod" service from the gRPC response type.
func NewShowResult(message *user_methodpb.ShowResponse) *usermethodviews.UserInfoView {
	result := &usermethodviews.UserInfoView{
		Name:     &message.Name,
		Email:    &message.Email,
		Icon:     &message.Icon,
		Activate: &message.Activate,
	}
	idptr := int(message.Id)
	result.ID = &idptr
	if message.Password != "" {
		result.Password = &message.Password
	}
	return result
}

// NewShowNotFoundError builds the error type of the "show" endpoint of the
// "userMethod" service from the gRPC error response type.
func NewShowNotFoundError(message *user_methodpb.ShowNotFoundError) *usermethod.NotFound {
	er := &usermethod.NotFound{
		Message: message.Message_,
		ID:      message.Id,
	}
	return er
}

// NewLoginRequest builds the gRPC request type from the payload of the "login"
// endpoint of the "userMethod" service.
func NewLoginRequest(payload *usermethod.LoginPayload) *user_methodpb.LoginRequest {
	message := &user_methodpb.LoginRequest{
		Email:    payload.Email,
		Password: payload.Password,
	}
	return message
}

// NewLoginResult builds the result type of the "login" endpoint of the
// "userMethod" service from the gRPC response type.
func NewLoginResult(message *user_methodpb.LoginResponse) *usermethod.Creds {
	result := &usermethod.Creds{
		JWT: message.Jwt,
	}
	return result
}

// NewChangeInfoRequest builds the gRPC request type from the payload of the
// "changeInfo" endpoint of the "userMethod" service.
func NewChangeInfoRequest(payload *usermethod.ChangeInfoPayload) *user_methodpb.ChangeInfoRequest {
	message := &user_methodpb.ChangeInfoRequest{
		Name: payload.Name,
		Icon: payload.Icon,
	}
	return message
}

// NewChangeInfoResult builds the result type of the "changeInfo" endpoint of
// the "userMethod" service from the gRPC response type.
func NewChangeInfoResult(message *user_methodpb.ChangeInfoResponse) *usermethodviews.UserInfoView {
	result := &usermethodviews.UserInfoView{
		Name:     &message.Name,
		Email:    &message.Email,
		Icon:     &message.Icon,
		Activate: &message.Activate,
	}
	idptr := int(message.Id)
	result.ID = &idptr
	if message.Password != "" {
		result.Password = &message.Password
	}
	return result
}

// NewChangePasswordRequest builds the gRPC request type from the payload of
// the "changePassword" endpoint of the "userMethod" service.
func NewChangePasswordRequest(payload *usermethod.ChangePasswordPayload) *user_methodpb.ChangePasswordRequest {
	message := &user_methodpb.ChangePasswordRequest{
		OldPassword: payload.OldPassword,
		NewPassword: payload.NewPassword,
	}
	return message
}

// NewChangePasswordResult builds the result type of the "changePassword"
// endpoint of the "userMethod" service from the gRPC response type.
func NewChangePasswordResult(message *user_methodpb.ChangePasswordResponse) *usermethod.ResponseResult {
	result := &usermethod.ResponseResult{
		Code: int(message.Code),
	}
	if message.Message_ != "" {
		result.Message = &message.Message_
	}
	if message.Data != nil {
		result.Data = make(map[string]string, len(message.Data))
		for key, val := range message.Data {
			tk := key
			tv := val
			result.Data[tk] = tv
		}
	}
	return result
}

// NewForgotPasswordRequest builds the gRPC request type from the payload of
// the "forgotPassword" endpoint of the "userMethod" service.
func NewForgotPasswordRequest(payload *usermethod.ForgotPasswordPayload) *user_methodpb.ForgotPasswordRequest {
	message := &user_methodpb.ForgotPasswordRequest{
		Code:        payload.Code,
		NewPassword: payload.NewPassword,
		Email:       payload.Email,
	}
	return message
}

// NewForgotPasswordResult builds the result type of the "forgotPassword"
// endpoint of the "userMethod" service from the gRPC response type.
func NewForgotPasswordResult(message *user_methodpb.ForgotPasswordResponse) *usermethod.ResponseResult {
	result := &usermethod.ResponseResult{
		Code: int(message.Code),
	}
	if message.Message_ != "" {
		result.Message = &message.Message_
	}
	if message.Data != nil {
		result.Data = make(map[string]string, len(message.Data))
		for key, val := range message.Data {
			tk := key
			tv := val
			result.Data[tk] = tv
		}
	}
	return result
}

// NewChangeEmailRequest builds the gRPC request type from the payload of the
// "changeEmail" endpoint of the "userMethod" service.
func NewChangeEmailRequest(payload *usermethod.ChangeEmailPayload) *user_methodpb.ChangeEmailRequest {
	message := &user_methodpb.ChangeEmailRequest{
		Email: payload.Email,
	}
	return message
}

// NewChangeEmailResult builds the result type of the "changeEmail" endpoint of
// the "userMethod" service from the gRPC response type.
func NewChangeEmailResult(message *user_methodpb.ChangeEmailResponse) *usermethod.ResponseResult {
	result := &usermethod.ResponseResult{
		Code: int(message.Code),
	}
	if message.Message_ != "" {
		result.Message = &message.Message_
	}
	if message.Data != nil {
		result.Data = make(map[string]string, len(message.Data))
		for key, val := range message.Data {
			tk := key
			tv := val
			result.Data[tk] = tv
		}
	}
	return result
}

// NewSendVerifyCodeRequest builds the gRPC request type from the payload of
// the "sendVerifyCode" endpoint of the "userMethod" service.
func NewSendVerifyCodeRequest(payload *usermethod.SendVerifyCodePayload) *user_methodpb.SendVerifyCodeRequest {
	message := &user_methodpb.SendVerifyCodeRequest{
		Email: payload.Email,
	}
	return message
}

// NewSendVerifyCodeResult builds the result type of the "sendVerifyCode"
// endpoint of the "userMethod" service from the gRPC response type.
func NewSendVerifyCodeResult(message *user_methodpb.SendVerifyCodeResponse) *usermethod.ResponseResult {
	result := &usermethod.ResponseResult{
		Code: int(message.Code),
	}
	if message.Message_ != "" {
		result.Message = &message.Message_
	}
	if message.Data != nil {
		result.Data = make(map[string]string, len(message.Data))
		for key, val := range message.Data {
			tk := key
			tv := val
			result.Data[tk] = tv
		}
	}
	return result
}

// NewActivateRequest builds the gRPC request type from the payload of the
// "activate" endpoint of the "userMethod" service.
func NewActivateRequest(payload *usermethod.ActivatePayload) *user_methodpb.ActivateRequest {
	message := &user_methodpb.ActivateRequest{
		Code: payload.Code,
	}
	return message
}

// NewActivateResult builds the result type of the "activate" endpoint of the
// "userMethod" service from the gRPC response type.
func NewActivateResult(message *user_methodpb.ActivateResponse) *usermethod.ResponseResult {
	result := &usermethod.ResponseResult{
		Code: int(message.Code),
	}
	if message.Message_ != "" {
		result.Message = &message.Message_
	}
	if message.Data != nil {
		result.Data = make(map[string]string, len(message.Data))
		for key, val := range message.Data {
			tk := key
			tv := val
			result.Data[tk] = tv
		}
	}
	return result
}

// ValidateShowResponse runs the validations defined on ShowResponse.
func ValidateShowResponse(message *user_methodpb.ShowResponse) (err error) {
	if utf8.RuneCountInString(message.Name) > 100 {
		err = goa.MergeErrors(err, goa.InvalidLengthError("message.name", message.Name, utf8.RuneCountInString(message.Name), 100, false))
	}
	if utf8.RuneCountInString(message.Email) > 100 {
		err = goa.MergeErrors(err, goa.InvalidLengthError("message.email", message.Email, utf8.RuneCountInString(message.Email), 100, false))
	}
	if utf8.RuneCountInString(message.Icon) > 2000 {
		err = goa.MergeErrors(err, goa.InvalidLengthError("message.icon", message.Icon, utf8.RuneCountInString(message.Icon), 2000, false))
	}
	return
}

// ValidateChangeInfoResponse runs the validations defined on
// ChangeInfoResponse.
func ValidateChangeInfoResponse(message *user_methodpb.ChangeInfoResponse) (err error) {
	if utf8.RuneCountInString(message.Name) > 100 {
		err = goa.MergeErrors(err, goa.InvalidLengthError("message.name", message.Name, utf8.RuneCountInString(message.Name), 100, false))
	}
	if utf8.RuneCountInString(message.Email) > 100 {
		err = goa.MergeErrors(err, goa.InvalidLengthError("message.email", message.Email, utf8.RuneCountInString(message.Email), 100, false))
	}
	if utf8.RuneCountInString(message.Icon) > 2000 {
		err = goa.MergeErrors(err, goa.InvalidLengthError("message.icon", message.Icon, utf8.RuneCountInString(message.Icon), 2000, false))
	}
	return
}
