// Code generated by goa v3.0.2, DO NOT EDIT.
//
// userMethod HTTP server types
//
// Command:
// $ goa gen usersvr/design

package server

import (
	"unicode/utf8"
	usermethod "usersvr/gen/user_method"
	usermethodviews "usersvr/gen/user_method/views"

	goa "goa.design/goa/v3/pkg"
)

// RegisterRequestBody is the type of the "userMethod" service "register"
// endpoint HTTP request body.
type RegisterRequestBody struct {
	// email used to perform login
	Email *string `form:"email,omitempty" json:"email,omitempty" xml:"email,omitempty"`
	// Password used to perform signin
	Password *string `form:"password,omitempty" json:"password,omitempty" xml:"password,omitempty"`
}

// LoginRequestBody is the type of the "userMethod" service "login" endpoint
// HTTP request body.
type LoginRequestBody struct {
	// Username used to perform signin
	Email *string `form:"email,omitempty" json:"email,omitempty" xml:"email,omitempty"`
	// Password used to perform signin
	Password *string `form:"password,omitempty" json:"password,omitempty" xml:"password,omitempty"`
}

// ChangeInfoRequestBody is the type of the "userMethod" service "changeInfo"
// endpoint HTTP request body.
type ChangeInfoRequestBody struct {
	// Name of bottle
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// Vintage of bottle
	Email *string `form:"email,omitempty" json:"email,omitempty" xml:"email,omitempty"`
	// Description of bottle
	Icon *string `form:"icon,omitempty" json:"icon,omitempty" xml:"icon,omitempty"`
}

// ChangePasswordRequestBody is the type of the "userMethod" service
// "changePassword" endpoint HTTP request body.
type ChangePasswordRequestBody struct {
	// old password
	OldPassword *string `form:"oldPassword,omitempty" json:"oldPassword,omitempty" xml:"oldPassword,omitempty"`
}

// ForgotPasswordRequestBody is the type of the "userMethod" service
// "forgotPassword" endpoint HTTP request body.
type ForgotPasswordRequestBody struct {
	Code *string `form:"code,omitempty" json:"code,omitempty" xml:"code,omitempty"`
}

// ChangeEmailRequestBody is the type of the "userMethod" service "changeEmail"
// endpoint HTTP request body.
type ChangeEmailRequestBody struct {
	Email *string `form:"email,omitempty" json:"email,omitempty" xml:"email,omitempty"`
}

// RegisterResponseBody is the type of the "userMethod" service "register"
// endpoint HTTP response body.
type RegisterResponseBody struct {
	// code
	Code int `form:"code" json:"code" xml:"code"`
	// message
	Message *string           `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
	Data    map[string]string `form:"data,omitempty" json:"data,omitempty" xml:"data,omitempty"`
}

// ShowResponseBody is the type of the "userMethod" service "show" endpoint
// HTTP response body.
type ShowResponseBody struct {
	// ID is the unique id of the bottle.
	ID int `form:"id" json:"id" xml:"id"`
	// Name of bottle
	Name string `form:"name" json:"name" xml:"name"`
	// Vintage of bottle
	Email string `form:"email" json:"email" xml:"email"`
	// Description of bottle
	Icon string `form:"icon" json:"icon" xml:"icon"`
}

// ShowResponseBodyChangeInfo is the type of the "userMethod" service "show"
// endpoint HTTP response body.
type ShowResponseBodyChangeInfo struct {
	// ID is the unique id of the bottle.
	ID int `form:"id" json:"id" xml:"id"`
	// Name of bottle
	Name string `form:"name" json:"name" xml:"name"`
	// Description of bottle
	Icon string `form:"icon" json:"icon" xml:"icon"`
}

// LoginResponseBody is the type of the "userMethod" service "login" endpoint
// HTTP response body.
type LoginResponseBody struct {
	// JWT token
	JWT string `form:"jwt" json:"jwt" xml:"jwt"`
}

// ShowNotFoundResponseBody is the type of the "userMethod" service "show"
// endpoint HTTP response body for the "not_found" error.
type ShowNotFoundResponseBody struct {
	// Message of error
	Message string `form:"message" json:"message" xml:"message"`
	// ID of missing user
	ID string `form:"id" json:"id" xml:"id"`
}

// NewRegisterResponseBody builds the HTTP response body from the result of the
// "register" endpoint of the "userMethod" service.
func NewRegisterResponseBody(res *usermethod.ResponseResult) *RegisterResponseBody {
	body := &RegisterResponseBody{
		Code:    res.Code,
		Message: res.Message,
	}
	if res.Data != nil {
		body.Data = make(map[string]string, len(res.Data))
		for key, val := range res.Data {
			tk := key
			tv := val
			body.Data[tk] = tv
		}
	}
	return body
}

// NewShowResponseBody builds the HTTP response body from the result of the
// "show" endpoint of the "userMethod" service.
func NewShowResponseBody(res *usermethodviews.UserInfoView) *ShowResponseBody {
	body := &ShowResponseBody{
		ID:    *res.ID,
		Name:  *res.Name,
		Email: *res.Email,
		Icon:  *res.Icon,
	}
	return body
}

// NewShowResponseBodyChangeInfo builds the HTTP response body from the result
// of the "show" endpoint of the "userMethod" service.
func NewShowResponseBodyChangeInfo(res *usermethodviews.UserInfoView) *ShowResponseBodyChangeInfo {
	body := &ShowResponseBodyChangeInfo{
		ID:   *res.ID,
		Name: *res.Name,
		Icon: *res.Icon,
	}
	return body
}

// NewLoginResponseBody builds the HTTP response body from the result of the
// "login" endpoint of the "userMethod" service.
func NewLoginResponseBody(res *usermethod.Creds) *LoginResponseBody {
	body := &LoginResponseBody{
		JWT: res.JWT,
	}
	return body
}

// NewShowNotFoundResponseBody builds the HTTP response body from the result of
// the "show" endpoint of the "userMethod" service.
func NewShowNotFoundResponseBody(res *usermethod.NotFound) *ShowNotFoundResponseBody {
	body := &ShowNotFoundResponseBody{
		Message: res.Message,
		ID:      res.ID,
	}
	return body
}

// NewRegisterPayload builds a userMethod service register endpoint payload.
func NewRegisterPayload(body *RegisterRequestBody) *usermethod.RegisterPayload {
	v := &usermethod.RegisterPayload{
		Email:    *body.Email,
		Password: *body.Password,
	}
	return v
}

// NewShowPayload builds a userMethod service show endpoint payload.
func NewShowPayload(view *string, token string) *usermethod.ShowPayload {
	return &usermethod.ShowPayload{
		View:  view,
		Token: token,
	}
}

// NewLoginPayload builds a userMethod service login endpoint payload.
func NewLoginPayload(body *LoginRequestBody) *usermethod.LoginPayload {
	v := &usermethod.LoginPayload{
		Email:    *body.Email,
		Password: *body.Password,
	}
	return v
}

// NewChangeInfoUser builds a userMethod service changeInfo endpoint payload.
func NewChangeInfoUser(body *ChangeInfoRequestBody) *usermethod.User {
	v := &usermethod.User{
		Name:  *body.Name,
		Email: *body.Email,
		Icon:  *body.Icon,
	}
	return v
}

// NewChangePasswordPayload builds a userMethod service changePassword endpoint
// payload.
func NewChangePasswordPayload(body *ChangePasswordRequestBody) *usermethod.ChangePasswordPayload {
	v := &usermethod.ChangePasswordPayload{
		OldPassword: *body.OldPassword,
	}
	return v
}

// NewForgotPasswordPayload builds a userMethod service forgotPassword endpoint
// payload.
func NewForgotPasswordPayload(body *ForgotPasswordRequestBody) *usermethod.ForgotPasswordPayload {
	v := &usermethod.ForgotPasswordPayload{
		Code: *body.Code,
	}
	return v
}

// NewChangeEmailPayload builds a userMethod service changeEmail endpoint
// payload.
func NewChangeEmailPayload(body *ChangeEmailRequestBody) *usermethod.ChangeEmailPayload {
	v := &usermethod.ChangeEmailPayload{
		Email: *body.Email,
	}
	return v
}

// ValidateRegisterRequestBody runs the validations defined on
// RegisterRequestBody
func ValidateRegisterRequestBody(body *RegisterRequestBody) (err error) {
	if body.Email == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("email", "body"))
	}
	if body.Password == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("password", "body"))
	}
	return
}

// ValidateLoginRequestBody runs the validations defined on LoginRequestBody
func ValidateLoginRequestBody(body *LoginRequestBody) (err error) {
	if body.Email == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("email", "body"))
	}
	if body.Password == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("password", "body"))
	}
	return
}

// ValidateChangeInfoRequestBody runs the validations defined on
// ChangeInfoRequestBody
func ValidateChangeInfoRequestBody(body *ChangeInfoRequestBody) (err error) {
	if body.Name == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("name", "body"))
	}
	if body.Email == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("email", "body"))
	}
	if body.Icon == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("icon", "body"))
	}
	if body.Name != nil {
		if utf8.RuneCountInString(*body.Name) > 100 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("body.name", *body.Name, utf8.RuneCountInString(*body.Name), 100, false))
		}
	}
	if body.Email != nil {
		if utf8.RuneCountInString(*body.Email) > 100 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("body.email", *body.Email, utf8.RuneCountInString(*body.Email), 100, false))
		}
	}
	if body.Icon != nil {
		if utf8.RuneCountInString(*body.Icon) > 2000 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("body.icon", *body.Icon, utf8.RuneCountInString(*body.Icon), 2000, false))
		}
	}
	return
}

// ValidateChangePasswordRequestBody runs the validations defined on
// ChangePasswordRequestBody
func ValidateChangePasswordRequestBody(body *ChangePasswordRequestBody) (err error) {
	if body.OldPassword == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("oldPassword", "body"))
	}
	return
}

// ValidateForgotPasswordRequestBody runs the validations defined on
// ForgotPasswordRequestBody
func ValidateForgotPasswordRequestBody(body *ForgotPasswordRequestBody) (err error) {
	if body.Code == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("code", "body"))
	}
	return
}

// ValidateChangeEmailRequestBody runs the validations defined on
// ChangeEmailRequestBody
func ValidateChangeEmailRequestBody(body *ChangeEmailRequestBody) (err error) {
	if body.Email == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("email", "body"))
	}
	return
}
