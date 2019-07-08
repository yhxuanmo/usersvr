// Code generated by goa v3.0.2, DO NOT EDIT.
//
// userMethod client HTTP transport
//
// Command:
// $ goa gen usersvr/design

package client

import (
	"context"
	"net/http"

	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
)

// Client lists the userMethod service endpoint HTTP clients.
type Client struct {
	// Register Doer is the HTTP client used to make requests to the register
	// endpoint.
	RegisterDoer goahttp.Doer

	// Show Doer is the HTTP client used to make requests to the show endpoint.
	ShowDoer goahttp.Doer

	// Login Doer is the HTTP client used to make requests to the login endpoint.
	LoginDoer goahttp.Doer

	// ChangeInfo Doer is the HTTP client used to make requests to the changeInfo
	// endpoint.
	ChangeInfoDoer goahttp.Doer

	// ChangePassword Doer is the HTTP client used to make requests to the
	// changePassword endpoint.
	ChangePasswordDoer goahttp.Doer

	// ForgotPassword Doer is the HTTP client used to make requests to the
	// forgotPassword endpoint.
	ForgotPasswordDoer goahttp.Doer

	// ChangeEmail Doer is the HTTP client used to make requests to the changeEmail
	// endpoint.
	ChangeEmailDoer goahttp.Doer

	// RestoreResponseBody controls whether the response bodies are reset after
	// decoding so they can be read again.
	RestoreResponseBody bool

	scheme  string
	host    string
	encoder func(*http.Request) goahttp.Encoder
	decoder func(*http.Response) goahttp.Decoder
}

// NewClient instantiates HTTP clients for all the userMethod service servers.
func NewClient(
	scheme string,
	host string,
	doer goahttp.Doer,
	enc func(*http.Request) goahttp.Encoder,
	dec func(*http.Response) goahttp.Decoder,
	restoreBody bool,
) *Client {
	return &Client{
		RegisterDoer:        doer,
		ShowDoer:            doer,
		LoginDoer:           doer,
		ChangeInfoDoer:      doer,
		ChangePasswordDoer:  doer,
		ForgotPasswordDoer:  doer,
		ChangeEmailDoer:     doer,
		RestoreResponseBody: restoreBody,
		scheme:              scheme,
		host:                host,
		decoder:             dec,
		encoder:             enc,
	}
}

// Register returns an endpoint that makes HTTP requests to the userMethod
// service register server.
func (c *Client) Register() goa.Endpoint {
	var (
		encodeRequest  = EncodeRegisterRequest(c.encoder)
		decodeResponse = DecodeRegisterResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		req, err := c.BuildRegisterRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		err = encodeRequest(req, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.RegisterDoer.Do(req)

		if err != nil {
			return nil, goahttp.ErrRequestError("userMethod", "register", err)
		}
		return decodeResponse(resp)
	}
}

// Show returns an endpoint that makes HTTP requests to the userMethod service
// show server.
func (c *Client) Show() goa.Endpoint {
	var (
		encodeRequest  = EncodeShowRequest(c.encoder)
		decodeResponse = DecodeShowResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		req, err := c.BuildShowRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		err = encodeRequest(req, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.ShowDoer.Do(req)

		if err != nil {
			return nil, goahttp.ErrRequestError("userMethod", "show", err)
		}
		return decodeResponse(resp)
	}
}

// Login returns an endpoint that makes HTTP requests to the userMethod service
// login server.
func (c *Client) Login() goa.Endpoint {
	var (
		encodeRequest  = EncodeLoginRequest(c.encoder)
		decodeResponse = DecodeLoginResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		req, err := c.BuildLoginRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		err = encodeRequest(req, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.LoginDoer.Do(req)

		if err != nil {
			return nil, goahttp.ErrRequestError("userMethod", "login", err)
		}
		return decodeResponse(resp)
	}
}

// ChangeInfo returns an endpoint that makes HTTP requests to the userMethod
// service changeInfo server.
func (c *Client) ChangeInfo() goa.Endpoint {
	var (
		encodeRequest  = EncodeChangeInfoRequest(c.encoder)
		decodeResponse = DecodeChangeInfoResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		req, err := c.BuildChangeInfoRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		err = encodeRequest(req, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.ChangeInfoDoer.Do(req)

		if err != nil {
			return nil, goahttp.ErrRequestError("userMethod", "changeInfo", err)
		}
		return decodeResponse(resp)
	}
}

// ChangePassword returns an endpoint that makes HTTP requests to the
// userMethod service changePassword server.
func (c *Client) ChangePassword() goa.Endpoint {
	var (
		encodeRequest  = EncodeChangePasswordRequest(c.encoder)
		decodeResponse = DecodeChangePasswordResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		req, err := c.BuildChangePasswordRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		err = encodeRequest(req, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.ChangePasswordDoer.Do(req)

		if err != nil {
			return nil, goahttp.ErrRequestError("userMethod", "changePassword", err)
		}
		return decodeResponse(resp)
	}
}

// ForgotPassword returns an endpoint that makes HTTP requests to the
// userMethod service forgotPassword server.
func (c *Client) ForgotPassword() goa.Endpoint {
	var (
		encodeRequest  = EncodeForgotPasswordRequest(c.encoder)
		decodeResponse = DecodeForgotPasswordResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		req, err := c.BuildForgotPasswordRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		err = encodeRequest(req, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.ForgotPasswordDoer.Do(req)

		if err != nil {
			return nil, goahttp.ErrRequestError("userMethod", "forgotPassword", err)
		}
		return decodeResponse(resp)
	}
}

// ChangeEmail returns an endpoint that makes HTTP requests to the userMethod
// service changeEmail server.
func (c *Client) ChangeEmail() goa.Endpoint {
	var (
		encodeRequest  = EncodeChangeEmailRequest(c.encoder)
		decodeResponse = DecodeChangeEmailResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		req, err := c.BuildChangeEmailRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		err = encodeRequest(req, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.ChangeEmailDoer.Do(req)

		if err != nil {
			return nil, goahttp.ErrRequestError("userMethod", "changeEmail", err)
		}
		return decodeResponse(resp)
	}
}