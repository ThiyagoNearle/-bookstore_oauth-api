package access_token

//  In Go, a test function must always use the following signature:
// func TestXxx(*testing.T)

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestAccessTokenConstants(t *testing.T) {
	//if expirationTime != 24 {
	//t.Error("expiration time should be 24 hours") }    // // Error is equivalent to Log followed by Fail , it means it takes the error message and log the error message and and finaly it fails

	// for the above 2 lines we can use the below one line
	//fmt.Printf("type of t is %T", t) // type of t is *testing.TPASS

	assert.EqualValues(t, 24, expirationTime, "expiration time should be 24 hours") // it check both or equal, if not, it throws error as we have given here as a parameter
}

// Next we have our function TestSum(). All tests in Go should be written in the format of func TestXxx(*testing.T) where Xxx can be any charater or number,
// and the first character needs to be an uppercase character, or a number.

func TestGetNewAccessToken(t *testing.T) { // T is a type passed to Test functions to manage test state and support formatted test logs.
	at := GetNewAccessToken()
	//fmt.Println("at", at)

	if at.IsExpired() {
		t.Error("brand new access token should not be expired")
	}

	//t.Fail(), which marks the current function as failed without halting its execution.

	assert.False(t, at.IsExpired(), "brand new access token should not be expired")

	// if at.AccessToken != "" {
	//		t.Error("new access token should hot have defined access token id") }

	// expected value == actual value, if not it throws error and the error message is below
	assert.EqualValues(t, "", at.AccessToken, "new access token should hot have defined access token id")

	// if at.UserId != 0 {
	//		t.Error("new access token should not have an associated user id")}

	assert.True(t, at.UserId == 0, "new access token should not have an associated user id")

}

func TestAccessTokenIsExpired(t *testing.T) {
	at := AccessToken{} // creating new variable  => all values in at are 0
	//expirationTimescheck := time.Unix(at.Expires, 0)
	//fmt.Println("expirationTimescheck", expirationTimescheck)
	//fmt.Println("\nemptyat", at) //> all values are 0

	//	if !at.IsExpired() {
	//		t.Error("empty access token should be expired by default")}

	assert.True(t, at.IsExpired(), "empty access token should be expired by default")

	at.Expires = time.Now().UTC().Add(3 * time.Hour).Unix() // unique syntax

	// if at.IsExpired() {
	//		t.Error("access token expiring in three hours from now should Not be expired")}

	assert.False(t, at.IsExpired(), "access token expiring in three hours from now should Not be expired")

}
