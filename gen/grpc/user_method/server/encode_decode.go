// Code generated by goa v3.0.2, DO NOT EDIT.
//
// userMethod gRPC server encoders and decoders
//
// Command:
// $ goa gen usersvr/design

package server

import (
	"context"
	"strings"
	user_methodpb "usersvr/gen/grpc/user_method/pb"
	usermethod "usersvr/gen/user_method"
	usermethodviews "usersvr/gen/user_method/views"

	goagrpc "goa.design/goa/v3/grpc"
	goa "goa.design/goa/v3/pkg"
	"google.golang.org/grpc/metadata"
)

// EncodeRegisterResponse encodes responses from the "userMethod" service
// "register" endpoint.
func EncodeRegisterResponse(ctx context.Context, v interface{}, hdr, trlr *metadata.MD) (interface{}, error) {
	result, ok := v.(*usermethod.ResponseResult)
	if !ok {
		return nil, goagrpc.ErrInvalidType("userMethod", "register", "*usermethod.ResponseResult", v)
	}
	resp := NewRegisterResponse(result)
	return resp, nil
}

// DecodeRegisterRequest decodes requests sent to "userMethod" service
// "register" endpoint.
func DecodeRegisterRequest(ctx context.Context, v interface{}, md metadata.MD) (interface{}, error) {
	var (
		message *user_methodpb.RegisterRequest
		ok      bool
	)
	{
		if message, ok = v.(*user_methodpb.RegisterRequest); !ok {
			return nil, goagrpc.ErrInvalidType("userMethod", "register", "*user_methodpb.RegisterRequest", v)
		}
	}
	var payload *usermethod.RegisterPayload
	{
		payload = NewRegisterPayload(message)
	}
	return payload, nil
}

// EncodeShowResponse encodes responses from the "userMethod" service "show"
// endpoint.
func EncodeShowResponse(ctx context.Context, v interface{}, hdr, trlr *metadata.MD) (interface{}, error) {
	vres, ok := v.(*usermethodviews.UserInfo)
	if !ok {
		return nil, goagrpc.ErrInvalidType("userMethod", "show", "*usermethodviews.UserInfo", v)
	}
	result := vres.Projected
	(*hdr).Append("goa-view", vres.View)
	resp := NewShowResponse(result)
	return resp, nil
}

// DecodeShowRequest decodes requests sent to "userMethod" service "show"
// endpoint.
func DecodeShowRequest(ctx context.Context, v interface{}, md metadata.MD) (interface{}, error) {
	var (
		view  *string
		token string
		err   error
	)
	{
		if vals := md.Get("view"); len(vals) > 0 {
			view = &vals[0]
		}
		if view != nil {
			if !(*view == "default" || *view == "tiny") {
				err = goa.MergeErrors(err, goa.InvalidEnumValueError("view", *view, []interface{}{"default", "tiny"}))
			}
		}
		if vals := md.Get("authorization"); len(vals) == 0 {
			err = goa.MergeErrors(err, goa.MissingFieldError("authorization", "metadata"))
		} else {
			token = vals[0]
		}
	}
	if err != nil {
		return nil, err
	}
	var payload *usermethod.ShowPayload
	{
		payload = NewShowPayload(view, token)
		if strings.Contains(payload.Token, " ") {
			// Remove authorization scheme prefix (e.g. "Bearer")
			cred := strings.SplitN(payload.Token, " ", 2)[1]
			payload.Token = cred
		}
	}
	return payload, nil
}

// EncodeLoginResponse encodes responses from the "userMethod" service "login"
// endpoint.
func EncodeLoginResponse(ctx context.Context, v interface{}, hdr, trlr *metadata.MD) (interface{}, error) {
	result, ok := v.(*usermethod.Creds)
	if !ok {
		return nil, goagrpc.ErrInvalidType("userMethod", "login", "*usermethod.Creds", v)
	}
	resp := NewLoginResponse(result)
	return resp, nil
}

// DecodeLoginRequest decodes requests sent to "userMethod" service "login"
// endpoint.
func DecodeLoginRequest(ctx context.Context, v interface{}, md metadata.MD) (interface{}, error) {
	var (
		message *user_methodpb.LoginRequest
		ok      bool
	)
	{
		if message, ok = v.(*user_methodpb.LoginRequest); !ok {
			return nil, goagrpc.ErrInvalidType("userMethod", "login", "*user_methodpb.LoginRequest", v)
		}
	}
	var payload *usermethod.LoginPayload
	{
		payload = NewLoginPayload(message)
	}
	return payload, nil
}

// EncodeChangeInfoResponse encodes responses from the "userMethod" service
// "changeInfo" endpoint.
func EncodeChangeInfoResponse(ctx context.Context, v interface{}, hdr, trlr *metadata.MD) (interface{}, error) {
	vres, ok := v.(*usermethodviews.UserInfo)
	if !ok {
		return nil, goagrpc.ErrInvalidType("userMethod", "changeInfo", "*usermethodviews.UserInfo", v)
	}
	result := vres.Projected
	(*hdr).Append("goa-view", vres.View)
	resp := NewChangeInfoResponse(result)
	return resp, nil
}

// DecodeChangeInfoRequest decodes requests sent to "userMethod" service
// "changeInfo" endpoint.
func DecodeChangeInfoRequest(ctx context.Context, v interface{}, md metadata.MD) (interface{}, error) {
	var (
		token string
		err   error
	)
	{
		if vals := md.Get("authorization"); len(vals) == 0 {
			err = goa.MergeErrors(err, goa.MissingFieldError("authorization", "metadata"))
		} else {
			token = vals[0]
		}
	}
	if err != nil {
		return nil, err
	}
	var (
		message *user_methodpb.ChangeInfoRequest
		ok      bool
	)
	{
		if message, ok = v.(*user_methodpb.ChangeInfoRequest); !ok {
			return nil, goagrpc.ErrInvalidType("userMethod", "changeInfo", "*user_methodpb.ChangeInfoRequest", v)
		}
	}
	var payload *usermethod.ChangeInfoPayload
	{
		payload = NewChangeInfoPayload(message, token)
		if strings.Contains(payload.Token, " ") {
			// Remove authorization scheme prefix (e.g. "Bearer")
			cred := strings.SplitN(payload.Token, " ", 2)[1]
			payload.Token = cred
		}
	}
	return payload, nil
}

// EncodeChangePasswordResponse encodes responses from the "userMethod" service
// "changePassword" endpoint.
func EncodeChangePasswordResponse(ctx context.Context, v interface{}, hdr, trlr *metadata.MD) (interface{}, error) {
	result, ok := v.(*usermethod.ResponseResult)
	if !ok {
		return nil, goagrpc.ErrInvalidType("userMethod", "changePassword", "*usermethod.ResponseResult", v)
	}
	resp := NewChangePasswordResponse(result)
	return resp, nil
}

// DecodeChangePasswordRequest decodes requests sent to "userMethod" service
// "changePassword" endpoint.
func DecodeChangePasswordRequest(ctx context.Context, v interface{}, md metadata.MD) (interface{}, error) {
	var (
		token string
		err   error
	)
	{
		if vals := md.Get("authorization"); len(vals) == 0 {
			err = goa.MergeErrors(err, goa.MissingFieldError("authorization", "metadata"))
		} else {
			token = vals[0]
		}
	}
	if err != nil {
		return nil, err
	}
	var (
		message *user_methodpb.ChangePasswordRequest
		ok      bool
	)
	{
		if message, ok = v.(*user_methodpb.ChangePasswordRequest); !ok {
			return nil, goagrpc.ErrInvalidType("userMethod", "changePassword", "*user_methodpb.ChangePasswordRequest", v)
		}
	}
	var payload *usermethod.ChangePasswordPayload
	{
		payload = NewChangePasswordPayload(message, token)
		if strings.Contains(payload.Token, " ") {
			// Remove authorization scheme prefix (e.g. "Bearer")
			cred := strings.SplitN(payload.Token, " ", 2)[1]
			payload.Token = cred
		}
	}
	return payload, nil
}

// EncodeForgotPasswordResponse encodes responses from the "userMethod" service
// "forgotPassword" endpoint.
func EncodeForgotPasswordResponse(ctx context.Context, v interface{}, hdr, trlr *metadata.MD) (interface{}, error) {
	result, ok := v.(*usermethod.ResponseResult)
	if !ok {
		return nil, goagrpc.ErrInvalidType("userMethod", "forgotPassword", "*usermethod.ResponseResult", v)
	}
	resp := NewForgotPasswordResponse(result)
	return resp, nil
}

// DecodeForgotPasswordRequest decodes requests sent to "userMethod" service
// "forgotPassword" endpoint.
func DecodeForgotPasswordRequest(ctx context.Context, v interface{}, md metadata.MD) (interface{}, error) {
	var (
		message *user_methodpb.ForgotPasswordRequest
		ok      bool
	)
	{
		if message, ok = v.(*user_methodpb.ForgotPasswordRequest); !ok {
			return nil, goagrpc.ErrInvalidType("userMethod", "forgotPassword", "*user_methodpb.ForgotPasswordRequest", v)
		}
	}
	var payload *usermethod.ForgotPasswordPayload
	{
		payload = NewForgotPasswordPayload(message)
	}
	return payload, nil
}

// EncodeChangeEmailResponse encodes responses from the "userMethod" service
// "changeEmail" endpoint.
func EncodeChangeEmailResponse(ctx context.Context, v interface{}, hdr, trlr *metadata.MD) (interface{}, error) {
	result, ok := v.(*usermethod.ResponseResult)
	if !ok {
		return nil, goagrpc.ErrInvalidType("userMethod", "changeEmail", "*usermethod.ResponseResult", v)
	}
	resp := NewChangeEmailResponse(result)
	return resp, nil
}

// DecodeChangeEmailRequest decodes requests sent to "userMethod" service
// "changeEmail" endpoint.
func DecodeChangeEmailRequest(ctx context.Context, v interface{}, md metadata.MD) (interface{}, error) {
	var (
		token string
		err   error
	)
	{
		if vals := md.Get("authorization"); len(vals) == 0 {
			err = goa.MergeErrors(err, goa.MissingFieldError("authorization", "metadata"))
		} else {
			token = vals[0]
		}
	}
	if err != nil {
		return nil, err
	}
	var (
		message *user_methodpb.ChangeEmailRequest
		ok      bool
	)
	{
		if message, ok = v.(*user_methodpb.ChangeEmailRequest); !ok {
			return nil, goagrpc.ErrInvalidType("userMethod", "changeEmail", "*user_methodpb.ChangeEmailRequest", v)
		}
	}
	var payload *usermethod.ChangeEmailPayload
	{
		payload = NewChangeEmailPayload(message, token)
		if strings.Contains(payload.Token, " ") {
			// Remove authorization scheme prefix (e.g. "Bearer")
			cred := strings.SplitN(payload.Token, " ", 2)[1]
			payload.Token = cred
		}
	}
	return payload, nil
}

// EncodeSendVerifyCodeResponse encodes responses from the "userMethod" service
// "sendVerifyCode" endpoint.
func EncodeSendVerifyCodeResponse(ctx context.Context, v interface{}, hdr, trlr *metadata.MD) (interface{}, error) {
	result, ok := v.(*usermethod.ResponseResult)
	if !ok {
		return nil, goagrpc.ErrInvalidType("userMethod", "sendVerifyCode", "*usermethod.ResponseResult", v)
	}
	resp := NewSendVerifyCodeResponse(result)
	return resp, nil
}

// DecodeSendVerifyCodeRequest decodes requests sent to "userMethod" service
// "sendVerifyCode" endpoint.
func DecodeSendVerifyCodeRequest(ctx context.Context, v interface{}, md metadata.MD) (interface{}, error) {
	var (
		message *user_methodpb.SendVerifyCodeRequest
		ok      bool
	)
	{
		if message, ok = v.(*user_methodpb.SendVerifyCodeRequest); !ok {
			return nil, goagrpc.ErrInvalidType("userMethod", "sendVerifyCode", "*user_methodpb.SendVerifyCodeRequest", v)
		}
	}
	var payload *usermethod.SendVerifyCodePayload
	{
		payload = NewSendVerifyCodePayload(message)
	}
	return payload, nil
}

// EncodeActivateResponse encodes responses from the "userMethod" service
// "activate" endpoint.
func EncodeActivateResponse(ctx context.Context, v interface{}, hdr, trlr *metadata.MD) (interface{}, error) {
	result, ok := v.(*usermethod.ResponseResult)
	if !ok {
		return nil, goagrpc.ErrInvalidType("userMethod", "activate", "*usermethod.ResponseResult", v)
	}
	resp := NewActivateResponse(result)
	return resp, nil
}

// DecodeActivateRequest decodes requests sent to "userMethod" service
// "activate" endpoint.
func DecodeActivateRequest(ctx context.Context, v interface{}, md metadata.MD) (interface{}, error) {
	var (
		message *user_methodpb.ActivateRequest
		ok      bool
	)
	{
		if message, ok = v.(*user_methodpb.ActivateRequest); !ok {
			return nil, goagrpc.ErrInvalidType("userMethod", "activate", "*user_methodpb.ActivateRequest", v)
		}
	}
	var payload *usermethod.ActivatePayload
	{
		payload = NewActivatePayload(message)
	}
	return payload, nil
}
