// Code generated by goa v3.0.2, DO NOT EDIT.
//
// HTTP request path constructors for the userMethod service.
//
// Command:
// $ goa gen usersvr/design

package server

// RegisterUserMethodPath returns the URL path to the userMethod service register HTTP endpoint.
func RegisterUserMethodPath() string {
	return "/user/register"
}

// ShowUserMethodPath returns the URL path to the userMethod service show HTTP endpoint.
func ShowUserMethodPath() string {
	return "/user"
}

// LoginUserMethodPath returns the URL path to the userMethod service login HTTP endpoint.
func LoginUserMethodPath() string {
	return "/user/login"
}

// ChangeInfoUserMethodPath returns the URL path to the userMethod service changeInfo HTTP endpoint.
func ChangeInfoUserMethodPath() string {
	return "/user/change/info"
}

// ChangePasswordUserMethodPath returns the URL path to the userMethod service changePassword HTTP endpoint.
func ChangePasswordUserMethodPath() string {
	return "/user/change/password"
}

// ForgotPasswordUserMethodPath returns the URL path to the userMethod service forgotPassword HTTP endpoint.
func ForgotPasswordUserMethodPath() string {
	return "/user/forgot/password"
}

// ChangeEmailUserMethodPath returns the URL path to the userMethod service changeEmail HTTP endpoint.
func ChangeEmailUserMethodPath() string {
	return "/user/change/email"
}

// SendVerifyCodeUserMethodPath returns the URL path to the userMethod service sendVerifyCode HTTP endpoint.
func SendVerifyCodeUserMethodPath() string {
	return "/user/send/code"
}
