// Code generated by goa v3.0.2, DO NOT EDIT.
//
// userMethod HTTP client CLI support package
//
// Command:
// $ goa gen usersvr/design

package client

import (
	"encoding/json"
	"fmt"
	usermethod "usersvr/gen/user_method"
)

// BuildRegisterPayload builds the payload for the userMethod register endpoint
// from CLI flags.
func BuildRegisterPayload(userMethodRegisterBody string) (*usermethod.RegisterPayload, error) {
	var err error
	var body RegisterRequestBody
	{
		err = json.Unmarshal([]byte(userMethodRegisterBody), &body)
		if err != nil {
			return nil, fmt.Errorf("invalid JSON for body, example of valid JSON:\n%s", "'{\n      \"email\": \"123@email.com\",\n      \"password\": \"password\"\n   }'")
		}
	}
	v := &usermethod.RegisterPayload{
		Email:    body.Email,
		Password: body.Password,
	}
	return v, nil
}

// BuildShowPayload builds the payload for the userMethod show endpoint from
// CLI flags.
func BuildShowPayload(userMethodShowView string, userMethodShowToken string) (*usermethod.ShowPayload, error) {
	var view *string
	{
		if userMethodShowView != "" {
			view = &userMethodShowView
		}
	}
	var token string
	{
		token = userMethodShowToken
	}
	payload := &usermethod.ShowPayload{
		View:  view,
		Token: token,
	}
	return payload, nil
}

// BuildLoginPayload builds the payload for the userMethod login endpoint from
// CLI flags.
func BuildLoginPayload(userMethodLoginBody string) (*usermethod.LoginPayload, error) {
	var err error
	var body LoginRequestBody
	{
		err = json.Unmarshal([]byte(userMethodLoginBody), &body)
		if err != nil {
			return nil, fmt.Errorf("invalid JSON for body, example of valid JSON:\n%s", "'{\n      \"email\": \"123@email.com\",\n      \"password\": \"password\"\n   }'")
		}
	}
	v := &usermethod.LoginPayload{
		Email:    body.Email,
		Password: body.Password,
	}
	return v, nil
}

// BuildChangeInfoPayload builds the payload for the userMethod changeInfo
// endpoint from CLI flags.
func BuildChangeInfoPayload(userMethodChangeInfoBody string, userMethodChangeInfoToken string) (*usermethod.ChangeInfoPayload, error) {
	var err error
	var body ChangeInfoRequestBody
	{
		err = json.Unmarshal([]byte(userMethodChangeInfoBody), &body)
		if err != nil {
			return nil, fmt.Errorf("invalid JSON for body, example of valid JSON:\n%s", "'{\n      \"icon\": \"Ipsam velit rem temporibus.\",\n      \"name\": \"Sed dolores.\"\n   }'")
		}
	}
	var token string
	{
		token = userMethodChangeInfoToken
	}
	v := &usermethod.ChangeInfoPayload{
		Name: body.Name,
		Icon: body.Icon,
	}
	v.Token = token
	return v, nil
}

// BuildChangePasswordPayload builds the payload for the userMethod
// changePassword endpoint from CLI flags.
func BuildChangePasswordPayload(userMethodChangePasswordBody string, userMethodChangePasswordToken string) (*usermethod.ChangePasswordPayload, error) {
	var err error
	var body ChangePasswordRequestBody
	{
		err = json.Unmarshal([]byte(userMethodChangePasswordBody), &body)
		if err != nil {
			return nil, fmt.Errorf("invalid JSON for body, example of valid JSON:\n%s", "'{\n      \"newPassword\": \"new password\",\n      \"oldPassword\": \"old password\"\n   }'")
		}
	}
	var token string
	{
		token = userMethodChangePasswordToken
	}
	v := &usermethod.ChangePasswordPayload{
		OldPassword: body.OldPassword,
		NewPassword: body.NewPassword,
	}
	v.Token = token
	return v, nil
}

// BuildForgotPasswordPayload builds the payload for the userMethod
// forgotPassword endpoint from CLI flags.
func BuildForgotPasswordPayload(userMethodForgotPasswordBody string) (*usermethod.ForgotPasswordPayload, error) {
	var err error
	var body ForgotPasswordRequestBody
	{
		err = json.Unmarshal([]byte(userMethodForgotPasswordBody), &body)
		if err != nil {
			return nil, fmt.Errorf("invalid JSON for body, example of valid JSON:\n%s", "'{\n      \"code\": \"1234\",\n      \"email\": \"Iusto qui dolor.\",\n      \"newPassword\": \"Cumque est consectetur quia possimus.\"\n   }'")
		}
	}
	v := &usermethod.ForgotPasswordPayload{
		Code:        body.Code,
		NewPassword: body.NewPassword,
		Email:       body.Email,
	}
	return v, nil
}

// BuildChangeEmailPayload builds the payload for the userMethod changeEmail
// endpoint from CLI flags.
func BuildChangeEmailPayload(userMethodChangeEmailBody string, userMethodChangeEmailToken string) (*usermethod.ChangeEmailPayload, error) {
	var err error
	var body ChangeEmailRequestBody
	{
		err = json.Unmarshal([]byte(userMethodChangeEmailBody), &body)
		if err != nil {
			return nil, fmt.Errorf("invalid JSON for body, example of valid JSON:\n%s", "'{\n      \"email\": \"Architecto quo sequi veniam repellendus ab.\"\n   }'")
		}
	}
	var token string
	{
		token = userMethodChangeEmailToken
	}
	v := &usermethod.ChangeEmailPayload{
		Email: body.Email,
	}
	v.Token = token
	return v, nil
}

// BuildSendVerifyCodePayload builds the payload for the userMethod
// sendVerifyCode endpoint from CLI flags.
func BuildSendVerifyCodePayload(userMethodSendVerifyCodeBody string) (*usermethod.SendVerifyCodePayload, error) {
	var err error
	var body SendVerifyCodeRequestBody
	{
		err = json.Unmarshal([]byte(userMethodSendVerifyCodeBody), &body)
		if err != nil {
			return nil, fmt.Errorf("invalid JSON for body, example of valid JSON:\n%s", "'{\n      \"email\": \"Illo recusandae est.\"\n   }'")
		}
	}
	v := &usermethod.SendVerifyCodePayload{
		Email: body.Email,
	}
	return v, nil
}

// BuildActivatePayload builds the payload for the userMethod activate endpoint
// from CLI flags.
func BuildActivatePayload(userMethodActivateCode string) (*usermethod.ActivatePayload, error) {
	var code string
	{
		code = userMethodActivateCode
	}
	payload := &usermethod.ActivatePayload{
		Code: code,
	}
	return payload, nil
}
