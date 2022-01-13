package access_token

import (
	"fmt"
	"strings"
	"time"

	"github.com/ThiyagoNearle/bookstore_oauth-api/src/utils/crypto_utils"
	"github.com/ThiyagoNearle/bookstore_oauth-api/src/utils/errors"
)

const (
	expirationTime             = 24
	grantTypePassword          = "password"
	grandTypeClientCredentials = "client_credentials"
)

type AccessTokenRequest struct {
	GrantType string `json:"grant_type"`
	Scope     string `json:"scope"`

	// Used for password grant type
	Username string `json:"username"`
	Password string `json:"password"`

	// Used for client_credentials grant type
	ClientId     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
}

func (at *AccessTokenRequest) Validate() *errors.RestErr {
	switch at.GrantType {
	case grantTypePassword:
		break

	case grandTypeClientCredentials:
		break

	default:
		return errors.NewsBadRequestError("invalid grant_type parameter")
	}

	//TODO: Validate parameters for each grant_type
	return nil
}

type AccessToken struct {
	AccessToken string `json:"access_token"`
	UserId      int64  `json:"user_id"`             // the user will give username, password to user api and if both are same in the records, it will give userId corresponding to the row
	ClientId    int64  `json:"client_id,omitempty"` // based on platform whether it is web frontend - Client-Id:123, ( need less expiration time ) // Android App - Client-Id:234, (need more expiration time)
	Expires     int64  `json:"expires"`
}

func (at *AccessToken) Validate() *errors.RestErr {
	at.AccessToken = strings.TrimSpace(at.AccessToken)
	if at.AccessToken == "" {
		return errors.NewsBadRequestError("invalid access token id")
	}
	if at.UserId <= 0 { // if we didnt enter anything or skip UserId, it takes as a 0
		return errors.NewsBadRequestError("invalid user id")
	}
	if at.ClientId <= 0 { // if we didnt enter anything or skip Client, it takes as a 0
		return errors.NewsBadRequestError("invalid client id")
	}
	if at.Expires <= 0 { // if we didnt enter anything or skip Expires, it takes as a 0
		return errors.NewsBadRequestError("invalid expiration time")
	}
	return nil

}

func GetNewAccessToken(userId int64) AccessToken {
	return AccessToken{
		UserId:  userId,
		Expires: time.Now().UTC().Add(expirationTime * time.Hour).Unix(), // unix is the time stamp of the current time, it converts date into seconds
	}
}

func (at AccessToken) IsExpired() bool { // this function always execute the whether it is true otherwise it throws error
	//now := time.Now().UTC()
	//expirationTime := time.Unix(at.Expires, 0) // convert that seconds into datetime

	// if you pass 0 to Expires then expirationTime is 1970-01-01 05:30:00 +0530 IST ( that means starting of that count)
	//return expirationTime.Before(now) // it means [ expirationTime < now  ]  if the condition is ok, it gives true ohtherwise it throws error
	// Before reports whether the time  t (expirationTime) is before u (now)

	// WE CAN WRITE THE ABOVE 3 LINES OF CODE IN THIS BELOW 1 LINE

	return time.Unix(at.Expires, 0).Before(time.Now().UTC())

	// time.Unix(at.Expires, 0) // convert that seconds into datetime
}

func (at *AccessToken) Generate() {
	at.AccessToken = crypto_utils.GetMd5(fmt.Sprintf("at-%d-%d-ran", at.UserId, at.Expires))
}
