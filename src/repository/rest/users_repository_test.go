package rest

import (
	"fmt"
	"net/http"
	"os"
	"testing"

	"github.com/mercadolibre/golang-restclient/rest"
	"github.com/stretchr/testify/assert"
)

// Entry point for every testing is TestMain fuction (*testing.M)

func TestMain(m *testing.M) {
	fmt.Println("TEST FLOW STARTED FROM HERE ONLY")
	rest.StartMockupServer()
	os.Exit(m.Run()) // m.Run returns integer

}

func TestLoginUserTimeoutFromApi(t *testing.T) {
	rest.FlushMockups() // this line removes any mock we might have from the previous test cases // so we are starting with fresh  mocup system.
	rest.AddMockups(&rest.Mock{
		HTTPMethod: http.MethodPost,
		URL:        "https://api.bookstore.com/users/login",
		ReqBody:    `{"email":"email@gmail.com", "password":"the-password"}`,
		// if we do above post request against the url & body
		// in output we will receive a statuscode & some body as usual
		//(But in this output, we want the status code as -1 & don't want any body ( so empty body))
		RespHTTPCode: -1,
		RespBody:     `{}`, // body should be in json
	})

	repository := usersRepository{}

	user, err := repository.LoginUser("email@gmail.com", "the-password")

	assert.Nil(t, user)                                               // user should be nil
	assert.NotNil(t, err)                                             // error should be not nill
	assert.EqualValues(t, http.StatusInternalServerError, err.Status) // both should be same 500, 500
	assert.EqualValues(t, "invalid restclient response when trying to login user", err.Message)

}

func TestLoginUserInvalidErrorInterface(t *testing.T) {
	rest.FlushMockups() // this line removes any mock we might have from the previous test cases // so we are starting with fresh  mocup system.
	rest.AddMockups(&rest.Mock{
		HTTPMethod: http.MethodPost,
		URL:        "https://api.bookstore.com/users/login",
		ReqBody:    `{"email":"email@gmail.com", "password":"the-password"}`,
		// we defining OUTPUT
		RespHTTPCode: http.StatusNotFound,                                                              // 404
		RespBody:     `{"message": "Invalid login credentials", "status": "404","error": "not_found"}`, // customized error output
	})

	repository := usersRepository{}

	user, err := repository.LoginUser("email@gmail.com", "the-password")

	assert.Nil(t, user)                                               // user should be nil
	assert.NotNil(t, err)                                             // error should be not nill
	assert.EqualValues(t, http.StatusInternalServerError, err.Status) // both should be same 500, 500
	assert.EqualValues(t, "Invalid error interface when trying to login user", err.Message)
}

func TestLoginUserInvalidLoginCredential(t *testing.T) {
	rest.FlushMockups() // this line removes any mock we might have from the previous test cases // so we are starting with fresh  mocup system.
	rest.AddMockups(&rest.Mock{
		HTTPMethod: http.MethodPost,
		URL:        "https://api.bookstore.com/users/login",
		ReqBody:    `{"email":"email@gmail.com", "password":"the-password"}`,
		// we defining OUTPUT
		RespHTTPCode: http.StatusNotFound,
		RespBody:     `{"message": "Invalid login credentials", "status": 404,"error": "not_found"}`,
	})

	repository := usersRepository{}

	user, err := repository.LoginUser("email@gmail.com", "the-password")

	assert.Nil(t, user)                                    // user should be nil
	assert.NotNil(t, err)                                  // error should be not nill
	assert.EqualValues(t, http.StatusNotFound, err.Status) // both should be same 500, 500
	assert.EqualValues(t, "Invalid Login credentials", err.Message)

}

func TestLoginUserInvalidUserJsonResponse(t *testing.T) {
	rest.FlushMockups() // this line removes any mock we might have from the previous test cases // so we are starting with fresh  mocup system.
	rest.AddMockups(&rest.Mock{
		HTTPMethod: http.MethodPost,
		URL:        "https://api.bookstore.com/users/login",
		ReqBody:    `{"email":"email@gmail.com", "password":"the-password"}`,
		// we defining OUTPUT
		RespHTTPCode: http.StatusOK,
		RespBody:     `{"id":"1", "first_name":"vic", "last_name":"nesh","email":"vic12@gmail.com"}`,
	})

	repository := usersRepository{}

	user, err := repository.LoginUser("email@gmail.com", "the-password")

	assert.Nil(t, user)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusInternalServerError, err.Status)
	assert.EqualValues(t, "error when trying to unmarshall users login response", err.Message)

}

func TestLoginUserDetails(t *testing.T) {
	rest.FlushMockups() // this line removes any mock we might have from the previous test cases // so we are starting with fresh  mocup system.
	rest.AddMockups(&rest.Mock{
		HTTPMethod: http.MethodPost,
		URL:        "https://api.bookstore.com/users/login",
		ReqBody:    `{"email":"email@gmail.com", "password":"the-password"}`,
		// we defining OUTPUT
		RespHTTPCode: http.StatusOK,
		RespBody:     `{"id":1, "first_name":"vic", "last_name":"nesh","email":"vic12@gmail.com"}`,
	})

	repository := usersRepository{}

	user, err := repository.LoginUser("email@gmail.com", "the-password")

	assert.Nil(t, err)
	assert.NotNil(t, user)
	assert.EqualValues(t, 1, user.Id)
	assert.EqualValues(t, "vic", user.FirstName)
	assert.EqualValues(t, "nesh", user.LastName)
	assert.EqualValues(t, "vic12@gmail.com", user.Email)

}
