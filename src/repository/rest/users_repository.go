package rest

import (
	"encoding/json"
	"time"

	"github.com/ThiyagoNearle/bookstore_oauth-api/src/domain/users"
	"github.com/ThiyagoNearle/bookstore_oauth-api/src/utils/errors"
	"github.com/mercadolibre/golang-restclient/rest" // Rest client library , majorly we use this librarly in production
)

// RequestBuilder is the baseline for creating requests
// There's a Default Builder that you may use for simple requests
// RequestBuilder si thread-safe, and you should store it for later re-used.

// RequestBuilder is type struct { consist many fields}	 and here we pass values for 2 fields

// BaseURL string         => Base URL to be used for each Request. The final URL will be = BaseURL + URL.

// Timeout time.Duration  =>  Complete request time out.

var (
	usersRestClient = rest.RequestBuilder{
		BaseURL: "https://api.bookstore.com",
		Timeout: 100 * time.Millisecond,
	}
)

type RestUserRepository interface {
	LoginUser(string, string) (*users.User, *errors.RestErr)
}

type usersRepository struct{}

func NewRestUsersRepository() RestUserRepository {
	return &usersRepository{}
}

func (r *usersRepository) LoginUser(email string, password string) (*users.User, *errors.RestErr) {
	request := users.UserLoginRequest{
		Email:    email,
		Password: password,
	}
	response := usersRestClient.Post("/users/login", request) // not capital POST ( this is taken from users api POST method & url is "/users/login")
	// Post take care of the json of the body we metioned

	// Response represents the response from an HTTP request.

	// The Client and Transport return Responses from servers once
	// the response headers have been received. The response body
	// is streamed on demand as the Body field is read.

	// In response we get nill situation
	if response == nil || response.Response == nil { // if we didn't give anything || if we have a timeout
		return nil, errors.NewInternalServerError("invalid restclient response when trying to login user")
	}

	// In response we get error situation  (mostly error status starts above 299 so only we use this)
	if response.StatusCode > 299 {
		var restErr errors.RestErr
		err := json.Unmarshal(response.Bytes(), &restErr) // we try convert that response error to our defined error(rest err interface), if it is converted then rest error variable hold the error
		if err != nil {                                   // if we can't able to convert it gives and and we store it in err variable that we craeted above while callling the unmarshall function
			return nil, errors.NewInternalServerError("Invalid error interface when trying to login user")
		}
		return nil, &restErr // returning address of the variable

	}

	// In response we get values situation  ( we get some values in response )

	var user users.User

	// attempting to save this values in some variable

	if err := json.Unmarshal(response.Bytes(), &user); err != nil {
		return nil, errors.NewInternalServerError("error when trying to unmarshal users login response")
	}
	return &user, nil
}
