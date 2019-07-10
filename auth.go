package usersvr

import (
	"context"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"time"
	usermethod "usersvr/gen/user_method"
	"usersvr/utils"
	"usersvr/utils/db/model"
	"usersvr/utils/db/pg"

	"goa.design/goa/v3/security"
)

var (
	// ErrUnauthorized is the error returned by Login when the request credentials
	// are invalid.
	ErrUnauthorized error = usermethod.Unauthorized("invalid username and password combination")

	// ErrInvalidToken is the error returned when the JWT token is invalid.
	ErrInvalidToken error = usermethod.Unauthorized("invalid token")

	// ErrInvalidTokenScopes is the error returned when the scopes provided in
	// the JWT token claims are invalid.
	//ErrInvalidTokenScopes error = usermethod.InvalidScopes("invalid scopes in token")

	// Key is the key used in JWT authentication
	Key = []byte(utils.Config.Jwt.Secret)
)

// JWTAuth implements the authorization logic for service "userMethod" for the
// "jwt" security scheme.
func (s *userMethodsrvc) JWTAuth(ctx context.Context, token string, scheme *security.JWTScheme) (context.Context, error) {
	//
	// TBD: add authorization logic.
	//
	// In case of authorization failure this function should return
	// one of the generated error structs, e.g.:
	//
	//    return ctx, myservice.MakeUnauthorizedError("invalid token")
	//
	// Alternatively this function may return an instance of
	// goa.ServiceError with a Name field value that matches one of
	// the design error names, e.g:
	//
	//    return ctx, goa.PermanentError("unauthorized", "invalid token")
	//
	type CustomClaims struct {
		jwt.StandardClaims
		//scopes []string
		UserEmail string
	}

	// authorize request
	// 1. parse JWT token, token key is hardcoded to "secret" in this example
	t, err := jwt.ParseWithClaims(token, &CustomClaims{}, func(_ *jwt.Token) (interface{}, error) { return []byte(Key), nil })
	if err != nil {
		return ctx, ErrInvalidToken
	}
	fmt.Println(t)
	return ctx, nil
}


func GenToken(Email string)(string, error){
	type CustomClaims struct {
		jwt.StandardClaims
		//scopes []string
		UserEmail string
	}
	claims := CustomClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour*24*7).Unix(),
		},
		Email,

	}
	//token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
	//	"nbf":    time.Now().Unix(),
	//	"iat":    time.Now().Unix(),
	//	//"scopes": []string{"api:read", "api:write"},
	//	"userEmail": Email,
	//})
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString(Key)
	if err != nil {
		return "", err
	}
	return t, nil
}

func TokenToUser(Token string)(*model.User, error){
	type CustomClaims struct {
		jwt.StandardClaims
		//scopes []string
		UserEmail string
	}
	var user *model.User
	token, err := jwt.ParseWithClaims(Token, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(Key), nil
	})
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		user,err = pg.FindOneUser(claims.UserEmail)
		if err != nil {
			return user, err
		} else {
			return user, nil
		}
	} else {
		return user, err
	}
}