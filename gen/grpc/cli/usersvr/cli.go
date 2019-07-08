// Code generated by goa v3.0.2, DO NOT EDIT.
//
// usersvr gRPC client CLI support package
//
// Command:
// $ goa gen usersvr/design

package cli

import (
	"flag"
	"fmt"
	"os"
	usermethodc "usersvr/gen/grpc/user_method/client"

	goa "goa.design/goa/v3/pkg"
	grpc "google.golang.org/grpc"
)

// UsageCommands returns the set of commands and sub-commands using the format
//
//    command (subcommand1|subcommand2|...)
//
func UsageCommands() string {
	return `user-method (register|show|login|change-info|change-password|forgot-password|change-email)
`
}

// UsageExamples produces an example of a valid invocation of the CLI tool.
func UsageExamples() string {
	return os.Args[0] + ` user-method register --message '{
      "email": "123@email.com",
      "password": "password"
   }'` + "\n" +
		""
}

// ParseEndpoint returns the endpoint and payload as specified on the command
// line.
func ParseEndpoint(cc *grpc.ClientConn, opts ...grpc.CallOption) (goa.Endpoint, interface{}, error) {
	var (
		userMethodFlags = flag.NewFlagSet("user-method", flag.ContinueOnError)

		userMethodRegisterFlags       = flag.NewFlagSet("register", flag.ExitOnError)
		userMethodRegisterMessageFlag = userMethodRegisterFlags.String("message", "", "")

		userMethodShowFlags     = flag.NewFlagSet("show", flag.ExitOnError)
		userMethodShowViewFlag  = userMethodShowFlags.String("view", "", "")
		userMethodShowTokenFlag = userMethodShowFlags.String("token", "REQUIRED", "")

		userMethodLoginFlags       = flag.NewFlagSet("login", flag.ExitOnError)
		userMethodLoginMessageFlag = userMethodLoginFlags.String("message", "", "")

		userMethodChangeInfoFlags       = flag.NewFlagSet("change-info", flag.ExitOnError)
		userMethodChangeInfoMessageFlag = userMethodChangeInfoFlags.String("message", "", "")

		userMethodChangePasswordFlags       = flag.NewFlagSet("change-password", flag.ExitOnError)
		userMethodChangePasswordMessageFlag = userMethodChangePasswordFlags.String("message", "", "")

		userMethodForgotPasswordFlags       = flag.NewFlagSet("forgot-password", flag.ExitOnError)
		userMethodForgotPasswordMessageFlag = userMethodForgotPasswordFlags.String("message", "", "")

		userMethodChangeEmailFlags       = flag.NewFlagSet("change-email", flag.ExitOnError)
		userMethodChangeEmailMessageFlag = userMethodChangeEmailFlags.String("message", "", "")
	)
	userMethodFlags.Usage = userMethodUsage
	userMethodRegisterFlags.Usage = userMethodRegisterUsage
	userMethodShowFlags.Usage = userMethodShowUsage
	userMethodLoginFlags.Usage = userMethodLoginUsage
	userMethodChangeInfoFlags.Usage = userMethodChangeInfoUsage
	userMethodChangePasswordFlags.Usage = userMethodChangePasswordUsage
	userMethodForgotPasswordFlags.Usage = userMethodForgotPasswordUsage
	userMethodChangeEmailFlags.Usage = userMethodChangeEmailUsage

	if err := flag.CommandLine.Parse(os.Args[1:]); err != nil {
		return nil, nil, err
	}

	if flag.NArg() < 2 { // two non flag args are required: SERVICE and ENDPOINT (aka COMMAND)
		return nil, nil, fmt.Errorf("not enough arguments")
	}

	var (
		svcn string
		svcf *flag.FlagSet
	)
	{
		svcn = flag.Arg(0)
		switch svcn {
		case "user-method":
			svcf = userMethodFlags
		default:
			return nil, nil, fmt.Errorf("unknown service %q", svcn)
		}
	}
	if err := svcf.Parse(flag.Args()[1:]); err != nil {
		return nil, nil, err
	}

	var (
		epn string
		epf *flag.FlagSet
	)
	{
		epn = svcf.Arg(0)
		switch svcn {
		case "user-method":
			switch epn {
			case "register":
				epf = userMethodRegisterFlags

			case "show":
				epf = userMethodShowFlags

			case "login":
				epf = userMethodLoginFlags

			case "change-info":
				epf = userMethodChangeInfoFlags

			case "change-password":
				epf = userMethodChangePasswordFlags

			case "forgot-password":
				epf = userMethodForgotPasswordFlags

			case "change-email":
				epf = userMethodChangeEmailFlags

			}

		}
	}
	if epf == nil {
		return nil, nil, fmt.Errorf("unknown %q endpoint %q", svcn, epn)
	}

	// Parse endpoint flags if any
	if svcf.NArg() > 1 {
		if err := epf.Parse(svcf.Args()[1:]); err != nil {
			return nil, nil, err
		}
	}

	var (
		data     interface{}
		endpoint goa.Endpoint
		err      error
	)
	{
		switch svcn {
		case "user-method":
			c := usermethodc.NewClient(cc, opts...)
			switch epn {
			case "register":
				endpoint = c.Register()
				data, err = usermethodc.BuildRegisterPayload(*userMethodRegisterMessageFlag)
			case "show":
				endpoint = c.Show()
				data, err = usermethodc.BuildShowPayload(*userMethodShowViewFlag, *userMethodShowTokenFlag)
			case "login":
				endpoint = c.Login()
				data, err = usermethodc.BuildLoginPayload(*userMethodLoginMessageFlag)
			case "change-info":
				endpoint = c.ChangeInfo()
				data, err = usermethodc.BuildChangeInfoPayload(*userMethodChangeInfoMessageFlag)
			case "change-password":
				endpoint = c.ChangePassword()
				data, err = usermethodc.BuildChangePasswordPayload(*userMethodChangePasswordMessageFlag)
			case "forgot-password":
				endpoint = c.ForgotPassword()
				data, err = usermethodc.BuildForgotPasswordPayload(*userMethodForgotPasswordMessageFlag)
			case "change-email":
				endpoint = c.ChangeEmail()
				data, err = usermethodc.BuildChangeEmailPayload(*userMethodChangeEmailMessageFlag)
			}
		}
	}
	if err != nil {
		return nil, nil, err
	}

	return endpoint, data, nil
}

// user-methodUsage displays the usage of the user-method command and its
// subcommands.
func userMethodUsage() {
	fmt.Fprintf(os.Stderr, `The storage service makes it possible to view, add or remove wine bottles.
Usage:
    %s [globalflags] user-method COMMAND [flags]

COMMAND:
    register: register user
    show: Show user info
    login: Creates a valid JWT
    change-info: Add new bottle and return its ID.
    change-password: Remove bottle from storage
    forgot-password: Rate bottles by IDs
    change-email: Add n number of bottles and return their IDs. This is a multipart request and each part has field name 'bottle' and contains the encoded bottle info to be added.

Additional help:
    %s user-method COMMAND --help
`, os.Args[0], os.Args[0])
}
func userMethodRegisterUsage() {
	fmt.Fprintf(os.Stderr, `%s [flags] user-method register -message JSON

register user
    -message JSON: 

Example:
    `+os.Args[0]+` user-method register --message '{
      "email": "123@email.com",
      "password": "password"
   }'
`, os.Args[0])
}

func userMethodShowUsage() {
	fmt.Fprintf(os.Stderr, `%s [flags] user-method show -view STRING -token STRING

Show user info
    -view STRING: 
    -token STRING: 

Example:
    `+os.Args[0]+` user-method show --view "tiny" --token "Velit rem temporibus possimus qui."
`, os.Args[0])
}

func userMethodLoginUsage() {
	fmt.Fprintf(os.Stderr, `%s [flags] user-method login -message JSON

Creates a valid JWT
    -message JSON: 

Example:
    `+os.Args[0]+` user-method login --message '{
      "email": "123@email.com",
      "password": "password"
   }'
`, os.Args[0])
}

func userMethodChangeInfoUsage() {
	fmt.Fprintf(os.Stderr, `%s [flags] user-method change-info -message JSON

Add new bottle and return its ID.
    -message JSON: 

Example:
    `+os.Args[0]+` user-method change-info --message '{
      "email": "po2",
      "icon": "Red wine blend with an emphasis on the Cabernet Franc grape and including other Bordeaux grape varietals and some Syrah",
      "name": "Blue\'s Cuvee"
   }'
`, os.Args[0])
}

func userMethodChangePasswordUsage() {
	fmt.Fprintf(os.Stderr, `%s [flags] user-method change-password -message JSON

Remove bottle from storage
    -message JSON: 

Example:
    `+os.Args[0]+` user-method change-password --message '{
      "newPassword": "new password",
      "oldPassword": "old password"
   }'
`, os.Args[0])
}

func userMethodForgotPasswordUsage() {
	fmt.Fprintf(os.Stderr, `%s [flags] user-method forgot-password -message JSON

Rate bottles by IDs
    -message JSON: 

Example:
    `+os.Args[0]+` user-method forgot-password --message '{
      "code": "1234",
      "newPassword": "Id voluptatem quibusdam nulla autem."
   }'
`, os.Args[0])
}

func userMethodChangeEmailUsage() {
	fmt.Fprintf(os.Stderr, `%s [flags] user-method change-email -message JSON

Add n number of bottles and return their IDs. This is a multipart request and each part has field name 'bottle' and contains the encoded bottle info to be added.
    -message JSON: 

Example:
    `+os.Args[0]+` user-method change-email --message '{
      "email": "Dolorem repudiandae blanditiis ab molestiae quidem ut."
   }'
`, os.Args[0])
}