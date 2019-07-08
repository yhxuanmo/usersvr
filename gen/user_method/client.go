// Code generated by goa v3.0.2, DO NOT EDIT.
//
// userMethod client
//
// Command:
// $ goa gen usersvr/design

package usermethod

import (
	"context"

	goa "goa.design/goa/v3/pkg"
)

// Client is the "userMethod" service client.
type Client struct {
	RegisterEndpoint       goa.Endpoint
	ShowEndpoint           goa.Endpoint
	LoginEndpoint          goa.Endpoint
	ChangeInfoEndpoint     goa.Endpoint
	ChangePasswordEndpoint goa.Endpoint
	ForgotPasswordEndpoint goa.Endpoint
	ChangeEmailEndpoint    goa.Endpoint
}

// NewClient initializes a "userMethod" service client given the endpoints.
func NewClient(register, show, login, changeInfo, changePassword, forgotPassword, changeEmail goa.Endpoint) *Client {
	return &Client{
		RegisterEndpoint:       register,
		ShowEndpoint:           show,
		LoginEndpoint:          login,
		ChangeInfoEndpoint:     changeInfo,
		ChangePasswordEndpoint: changePassword,
		ForgotPasswordEndpoint: forgotPassword,
		ChangeEmailEndpoint:    changeEmail,
	}
}

// Register calls the "register" endpoint of the "userMethod" service.
// Register may return the following errors:
//	- "email_exist" (type *EmailExist): email is existed
//	- error: internal error
func (c *Client) Register(ctx context.Context, p *RegisterPayload) (res *ResponseResult, err error) {
	var ires interface{}
	ires, err = c.RegisterEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(*ResponseResult), nil
}

// Show calls the "show" endpoint of the "userMethod" service.
// Show may return the following errors:
//	- "not_found" (type *NotFound): Bottle not found
//	- error: internal error
func (c *Client) Show(ctx context.Context, p *ShowPayload) (res *UserInfo, err error) {
	var ires interface{}
	ires, err = c.ShowEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(*UserInfo), nil
}

// Login calls the "login" endpoint of the "userMethod" service.
// Login may return the following errors:
//	- "passwordError" (type PasswordError)
//	- error: internal error
func (c *Client) Login(ctx context.Context, p *LoginPayload) (res *Creds, err error) {
	var ires interface{}
	ires, err = c.LoginEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(*Creds), nil
}

// ChangeInfo calls the "changeInfo" endpoint of the "userMethod" service.
func (c *Client) ChangeInfo(ctx context.Context, p *User) (res string, err error) {
	var ires interface{}
	ires, err = c.ChangeInfoEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(string), nil
}

// ChangePassword calls the "changePassword" endpoint of the "userMethod"
// service.
// ChangePassword may return the following errors:
//	- "not_found" (type *NotFound): Bottle not found
//	- error: internal error
func (c *Client) ChangePassword(ctx context.Context, p *ChangePasswordPayload) (err error) {
	_, err = c.ChangePasswordEndpoint(ctx, p)
	return
}

// ForgotPassword calls the "forgotPassword" endpoint of the "userMethod"
// service.
func (c *Client) ForgotPassword(ctx context.Context, p *ForgotPasswordPayload) (err error) {
	_, err = c.ForgotPasswordEndpoint(ctx, p)
	return
}

// ChangeEmail calls the "changeEmail" endpoint of the "userMethod" service.
// ChangeEmail may return the following errors:
//	- "email_exist" (type *EmailExist): email is existed
//	- error: internal error
func (c *Client) ChangeEmail(ctx context.Context, p *ChangeEmailPayload) (res string, err error) {
	var ires interface{}
	ires, err = c.ChangeEmailEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(string), nil
}