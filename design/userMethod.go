package design

import . "goa.design/goa/v3/dsl"

var _ = Service("userMethod", func() {
	Description("The storage service makes it possible to view, add or remove wine bottles.")

	HTTP(func() {
		Path("/user")
	})

	Error("unauthorized", String, "Credentials are invalid")

	Method("register", func() {
		Description("register user")
		Payload(func() {
			Field(1, "email", String, "email used to perform login", func() {
				Example("123@email.com")
			})
			Field(2, "password", String, "Password used to perform signin", func() {
				Example("password")
			})
			Required("email", "password")
		})
		Error("email_exist", EmailExist, "email is existed")
		Result(ResponseResult)
		HTTP(func() {
			POST("/register")
			Response(StatusOK)
		})
		GRPC(func() {
			Response(CodeOK)
		})
	})

	Method("show", func() {
		Description("Show user info")
		Security(JWTAuth)
		Payload(func() {
			//Field(1, "id", String, "ID of bottle to show")
			TokenField(1, "token", String, func() {
				Description("JWT used for authentication")
			})

			Field(2, "view", String, "View to render", func() {
				Enum("default", "tiny")
			})
			Required("token")
		})
		Result(UserInfo)
		Error("not_found", NotFound, "Bottle not found")
		HTTP(func() {
			GET("/")
			Param("view")
			Response(StatusOK)
			Response("not_found", StatusNotFound)
		})
		GRPC(func() {
			Metadata(func() {
				Attribute("view")
			})
			Response(CodeOK)
			Response("not_found", CodeNotFound)
		})
	})

	Method("login", func() {
		Description("Creates a valid JWT")

		// The signin endpoint is secured via basic auth
		//Security(JWTAuth)

		Payload(func() {
			Description("Credentials used to authenticate to retrieve JWT token")
			Field(1, "email", String, "Username used to perform signin", func() {
				Example("123@email.com")
			})
			Field(2, "password", String, "Password used to perform signin", func() {
				Example("password")
			})
			Required("email", "password")
			//TokenField(1, "token", String, func() {
			//	Description("JWT used for authentication")
			//})
			//Required("token")
		})
		Error("passwordError", String, "password is error")
		Result(Creds)

		HTTP(func() {
			POST("/login")
			// Use Authorization header to provide basic auth value.
			Response(StatusOK)
		})

		GRPC(func() {
			Response(CodeOK)
		})
	})

	//Method("secure", func() {
	//	Description("This action is secured with the jwt scheme")
	//
	//	Security(JWTAuth, func() { // Use JWT to auth requests to this endpoint.
	//		Scope("api:read") // Enforce presence of "api:read" scope in JWT claims.
	//	})
	//
	//	Payload(func() {
	//		Field(1, "fail", Boolean, func() {
	//			Description("Whether to force auth failure even with a valid JWT")
	//		})
	//		TokenField(2, "token", String, func() {
	//			Description("JWT used for authentication")
	//		})
	//		Required("token")
	//	})
	//
	//	Result(String)
	//
	//	Error("invalid-scopes", String, "Token scopes are invalid")
	//
	//	HTTP(func() {
	//		GET("/secure")
	//		Param("fail")
	//		Response(StatusOK)
	//		Response("invalid-scopes", StatusForbidden)
	//	})
	//
	//	GRPC(func() {
	//		Response(CodeOK)
	//		Response("invalid-scopes", CodeUnauthenticated)
	//	})
	//})

	Method("changeInfo", func() {
		Description("Add new bottle and return its ID.")
		Payload(func() {
			TokenField(1, "token", String, func() {
				Description("JWT used for authentication")
			})

			Field(2, "name", String, "name of user")
			Field(3, "icon", String, "icon of user")
			Required("token", "name", "icon")
		})
		Security(JWTAuth)
		Result(UserInfo)
		HTTP(func() {
			POST("/change/info")
			Response(StatusOK)
		})
		GRPC(func() {
			Response(CodeOK)
		})
	})

	Method("changePassword", func() {
		Description("Remove bottle from storage")
		Security(JWTAuth)
		Payload(func() {
			TokenField(1, "token", String, func() {
				Description("JWT used for authentication")
			})
			Field(2,"oldPassword", String, "old password", func() {
				Example("old password")
			})
			Field(3, "newPassword", String, "new password", func() {
				Example("new password")
			})
			Required("token", "oldPassword", "newPassword")
		})
		Error("not_found", NotFound, "Bottle not found")
		Result(ResponseResult)
		HTTP(func() {
			POST("/change/password")
			Response(StatusOK)
		})
		GRPC(func() {
			Response(CodeOK)
		})
	})

	Method("forgotPassword", func() {
		Description("Rate bottles by IDs")
		Payload(func() {
			Field(1, "code", String, func() {
				Example("1234")
			})
			Field(2,"newPassword", String)
			Field(3, "email", String)
			Required("code", "newPassword", "email")
		})
		Result(ResponseResult)
		HTTP(func() {
			POST("/forgot/password")
			Response(StatusOK)
		})
		GRPC(func() {
			Response(CodeOK)
		})
	})

	Method("changeEmail", func() {
		Description("Add n number of bottles and return their IDs. This is a multipart request and each part has field name 'bottle' and contains the encoded bottle info to be added.")
		Payload(func() {
			TokenField(1, "token", String, func() {
				Description("JWT used for authentication")
			})
			Field(2,"email",String)
			Required("token", "email")
		})
		Security(JWTAuth)
		Error("email_exist", EmailExist, "email is existed")
		Result(ResponseResult)
		HTTP(func() {
			POST("/change/email")
			Response(StatusOK)
		})
		GRPC(func() {
			Response(CodeOK)
		})
	})

	Method("sendVerifyCode", func() {
		Description("send verify code to email")
		Payload(func() {
			Field(1, "email", String)
			Required("email")
		})
		Result(ResponseResult)
		HTTP(func() {
			POST("/send/code")
			Response(StatusOK)
		})
		GRPC(func() {
			Response(CodeOK)
		})
	})

	Method("activate", func() {
		Description("activate")
		Payload(func() {
			Field(1, "code", String)
			Required("code")
		})
		Result(ResponseResult)
		HTTP(func() {
			GET("/activate/{code}")
			Response(StatusOK)
		})
		GRPC(func() {
			Response(CodeOK)
		})
	})
})

var Creds = Type("Creds", func() {
	Field(1, "jwt", String, "JWT token", func() {
		Example("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiYWRtaW4iOnRydWV9.TJVA95OrM7E2cBab30RMHrHDcEfxjoYZgeFONFh7HgQ")
	})
	//Field(2, "api_key", String, "API Key", func() {
	//	Example("abcdef12345")
	//})
	//Field(3, "oauth_token", String, "OAuth2 token", func() {
	//	Example("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiYWRtaW4iOnRydWV9.TJVA95OrM7E2cBab30RMHrHDcEfxjoYZgeFONFh7HgQ")
	//})
	//Required("jwt", "api_key", "oauth_token")
	Required("jwt")
})

var ResponseResult = Type("ResponseResult", func() {
	Field(1, "code", Int, "code", func() {
		Example(200)
	})
	Field(2, "message", String, "message", func() {
		Example("bad request")
	})
	Field(3, "data", MapOf(String, String))
	Required("code")
})