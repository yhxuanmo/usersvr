// Code generated by goa v3.0.2, DO NOT EDIT.
//
// userMethod endpoints
//
// Command:
// $ goa gen usersvr/design

package usermethod

import (
	"context"

	goa "goa.design/goa/v3/pkg"
	"goa.design/goa/v3/security"
)

// Endpoints wraps the "userMethod" service endpoints.
type Endpoints struct {
	Register       goa.Endpoint
	Show           goa.Endpoint
	Login          goa.Endpoint
	ChangeInfo     goa.Endpoint
	ChangePassword goa.Endpoint
	ForgotPassword goa.Endpoint
	ChangeEmail    goa.Endpoint
	SendVerifyCode goa.Endpoint
	Activate       goa.Endpoint
}

// NewEndpoints wraps the methods of the "userMethod" service with endpoints.
func NewEndpoints(s Service) *Endpoints {
	// Casting service to Auther interface
	a := s.(Auther)
	return &Endpoints{
		Register:       NewRegisterEndpoint(s),
		Show:           NewShowEndpoint(s, a.JWTAuth),
		Login:          NewLoginEndpoint(s),
		ChangeInfo:     NewChangeInfoEndpoint(s, a.JWTAuth),
		ChangePassword: NewChangePasswordEndpoint(s, a.JWTAuth),
		ForgotPassword: NewForgotPasswordEndpoint(s),
		ChangeEmail:    NewChangeEmailEndpoint(s, a.JWTAuth),
		SendVerifyCode: NewSendVerifyCodeEndpoint(s),
		Activate:       NewActivateEndpoint(s),
	}
}

// Use applies the given middleware to all the "userMethod" service endpoints.
func (e *Endpoints) Use(m func(goa.Endpoint) goa.Endpoint) {
	e.Register = m(e.Register)
	e.Show = m(e.Show)
	e.Login = m(e.Login)
	e.ChangeInfo = m(e.ChangeInfo)
	e.ChangePassword = m(e.ChangePassword)
	e.ForgotPassword = m(e.ForgotPassword)
	e.ChangeEmail = m(e.ChangeEmail)
	e.SendVerifyCode = m(e.SendVerifyCode)
	e.Activate = m(e.Activate)
}

// NewRegisterEndpoint returns an endpoint function that calls the method
// "register" of service "userMethod".
func NewRegisterEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*RegisterPayload)
		return s.Register(ctx, p)
	}
}

// NewShowEndpoint returns an endpoint function that calls the method "show" of
// service "userMethod".
func NewShowEndpoint(s Service, authJWTFn security.AuthJWTFunc) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*ShowPayload)
		var err error
		sc := security.JWTScheme{
			Name:           "jwt",
			Scopes:         []string{},
			RequiredScopes: []string{},
		}
		ctx, err = authJWTFn(ctx, p.Token, &sc)
		if err != nil {
			return nil, err
		}
		res, view, err := s.Show(ctx, p)
		if err != nil {
			return nil, err
		}
		vres := NewViewedUserInfo(res, view)
		return vres, nil
	}
}

// NewLoginEndpoint returns an endpoint function that calls the method "login"
// of service "userMethod".
func NewLoginEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*LoginPayload)
		return s.Login(ctx, p)
	}
}

// NewChangeInfoEndpoint returns an endpoint function that calls the method
// "changeInfo" of service "userMethod".
func NewChangeInfoEndpoint(s Service, authJWTFn security.AuthJWTFunc) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*ChangeInfoPayload)
		var err error
		sc := security.JWTScheme{
			Name:           "jwt",
			Scopes:         []string{},
			RequiredScopes: []string{},
		}
		ctx, err = authJWTFn(ctx, p.Token, &sc)
		if err != nil {
			return nil, err
		}
		res, view, err := s.ChangeInfo(ctx, p)
		if err != nil {
			return nil, err
		}
		vres := NewViewedUserInfo(res, view)
		return vres, nil
	}
}

// NewChangePasswordEndpoint returns an endpoint function that calls the method
// "changePassword" of service "userMethod".
func NewChangePasswordEndpoint(s Service, authJWTFn security.AuthJWTFunc) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*ChangePasswordPayload)
		var err error
		sc := security.JWTScheme{
			Name:           "jwt",
			Scopes:         []string{},
			RequiredScopes: []string{},
		}
		ctx, err = authJWTFn(ctx, p.Token, &sc)
		if err != nil {
			return nil, err
		}
		return s.ChangePassword(ctx, p)
	}
}

// NewForgotPasswordEndpoint returns an endpoint function that calls the method
// "forgotPassword" of service "userMethod".
func NewForgotPasswordEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*ForgotPasswordPayload)
		return s.ForgotPassword(ctx, p)
	}
}

// NewChangeEmailEndpoint returns an endpoint function that calls the method
// "changeEmail" of service "userMethod".
func NewChangeEmailEndpoint(s Service, authJWTFn security.AuthJWTFunc) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*ChangeEmailPayload)
		var err error
		sc := security.JWTScheme{
			Name:           "jwt",
			Scopes:         []string{},
			RequiredScopes: []string{},
		}
		ctx, err = authJWTFn(ctx, p.Token, &sc)
		if err != nil {
			return nil, err
		}
		return s.ChangeEmail(ctx, p)
	}
}

// NewSendVerifyCodeEndpoint returns an endpoint function that calls the method
// "sendVerifyCode" of service "userMethod".
func NewSendVerifyCodeEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*SendVerifyCodePayload)
		return s.SendVerifyCode(ctx, p)
	}
}

// NewActivateEndpoint returns an endpoint function that calls the method
// "activate" of service "userMethod".
func NewActivateEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*ActivatePayload)
		return s.Activate(ctx, p)
	}
}
