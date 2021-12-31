package access_token

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestAccessTokenConstants(t *testing.T) {
	//if expirationTime != 24 {
	//	t.Error("expiration time should be 24 hours") }

	// for the above 2 lines we can use the below one line

	assert.EqualValues(t, 24, expirationTime, "expiration time should be 24 hours") //diwqjhidoooooooooooooooooooooooooooooohhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhh
}

func TestGetNewAccessToken(t *testing.T) { // T is a type passed to Test functions to manage test state and support formatted test logs.
	at := GetNewAccessToken()

	// if at.IsExpired() {
	//		t.Error("brand new access token should not be expired")}

	assert.False(t, at.IsExpired(), "brand new access token should not be expired")

	// if at.AccessToken != "" {
	//		t.Error("new access token should hot have defined access token id") }

	assert.EqualValues(t, "", at.AccessToken, "new access token should hot have defined access token id")

	// if at.UserId != 0 {
	//		t.Error("new access token should not have an associated user id")}

	assert.True(t, at.UserId == 0, "new access token should not have an associated user id")

}

func TestAccessTokenIsExpired(t *testing.T) {
	at := AccessToken{}

	//	if !at.IsExpired() {
	//		t.Error("empty access token should be expired by default")}

	assert.True(t, at.IsExpired(), "empty access token should be expired by default")

	at.Expires = time.Now().UTC().Add(3 * time.Hour).Unix() // unique syntax

	// if at.IsExpired() {
	//		t.Error("access token expiring in three hours from now should Not be expired")}

	assert.False(t, at.IsExpired(), "access token expiring in three hours from now should Not be expired")

}
