package design

import . "goa.design/goa/v3/dsl"

var _ = API("usersvr", func() {
	Title("User Service")
	Description("HTTP service for managing user info")

	Server("usersvr", func() {
		Description("cellar hosts the storage, sommelier and swagger services.")
		Services("userMethod")
		Host("localhost", func() {
			Description("default host")
			URI("http://localhost:8000/user")
			URI("grpc://localhost:8080/user")
		})
		//Host("goa.design", func() {
		//	Description("public host")
		//	URI("https://goa.design/cellar")
		//	URI("grpcs://goa.design/cellar")
		//})
	})
})

var UserInfo = ResultType("application/vnd.cellar.stored-bottle", func() {
	Description("A StoredBottle describes a bottle retrieved by the storage service.")
	Reference(User)
	TypeName("UserInfo")

	Attributes(func() {
		Attribute("id", Int, "ID is the unique id of the bottle.", func() {
			Example(1)
			Meta("rpc:tag", "8")
		})
		Field(2, "name")
		Field(3, "email")
		Field(4, "icon")
		Field(5, "password")
		//Field(6, "description")
		//Field(7, "rating")
	})

	View("default", func() {
		Attribute("id", Int)
		Attribute("name")
		Attribute("email")
		Attribute("icon")
	})

	View("changeInfo", func() {
		Attribute("id", Int)
		Attribute("name")
		Attribute("icon")
	})

	Required("id", "name", "icon")
})

var User = Type("User", func() {
	Description("Bottle describes a bottle of wine to be stored.")
	Attribute("name", String, "Name of bottle", func() {
		MaxLength(100)
		Example("Blue's Cuvee")
		Meta("rpc:tag", "1")
	})
	Attribute("email", String, "Vintage of bottle", func() {
		MaxLength(100)
		Meta("rpc:tag", "2")
	})
	Attribute("icon", String, "Description of bottle", func() {
		MaxLength(2000)
		Example("Red wine blend with an emphasis on the Cabernet Franc grape and including other Bordeaux grape varietals and some Syrah")
		Meta("rpc:tag", "3")
	})
	Required("name", "email", "icon")
})

var NotFound = Type("NotFound", func() {
	Description("NotFound is the type returned when attempting to show or delete a user that does not exist.")
	Attribute("message", String, "Message of error", func() {
		Meta("struct:error:name")
		Example("user 1 not found")
		Meta("rpc:tag", "1")
	})
	Field(2, "id", String, "ID of missing user")
	Required("message", "id")
})

var PasswordError = Type("PasswordError", func() {
	Description("NotFound is the type returned when attempting to show or delete a user that does not exist.")
	Attribute("message", String, "Message of error", func() {
		Meta("struct:error:name")
		Example("password error")
		Meta("rpc:tag", "1")
	})
	Required("message")
})

var EmailExist = Type("EmailExist", func() {
	Description("EmailExist is the type returned when user register with the email that is existed")
	Attribute("message", String, "Message of error", func() {
		Meta("struct:error:name")
		Example("email 123@email.com is existed")
		Meta("rpc:tag", "1")
	})
	Field(2, "email", String, "email of user")
	Required("message", "email")
})

var JWTAuth = JWTSecurity("jwt", func() {
	Description(`Secures endpoint by requiring a valid JWT token retrieved via the signin endpoint. Supports scopes "api:read" and "api:write".`)
	//Scope("api:read", "Read-only access")
	//Scope("api:write", "Read and write access")
})

var BasicAuth = BasicAuthSecurity("basic", func() {
	Description("Basic authentication used to authenticate security principal during signin")
	Scope("api:read", "Read-only access")
})