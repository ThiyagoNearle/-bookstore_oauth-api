package access_token

import (
	"time"
)

const (
	expirationTime = 24 // accesstoken expiration time 24 hours
)

type AccessToken struct {
	AccessToken string `json:"access_Token"`
	UserId      int64  `json:"userId"`
	ClientId    int64  `json:"clientId"`
	Expires     int64  `json:"expires"`
}

// web frontend - Client-Id:123,
// Android App - Client-Id:234,

func GetNewAccessToken() AccessToken {
	return AccessToken{
		Expires: time.Now().UTC().Add(expirationTime * time.Hour).Unix(), // unix is the time stamp of the current time, it gives result in seconds
	}
}

func (at AccessToken) IsExpired() bool { // this function always execute the whether it is true otherwise it throws error
	now := time.Now().UTC()
	expirationTime := time.Unix(at.Expires, 0) // convert that seconds into datetime

	// if you pass 0 to Expires then expirationTime is 1970-01-01 05:30:00 +0530 IST ( that means starting of that count)
	return expirationTime.Before(now) // it means [ expirationTime < now  ]  if the condition is ok, it gives true ohtherwise it throws error
	// Before reports whether the time  t (expirationTime) is before u (now)

	// WE CAN WRITE THE ABOVE 3 LINES OF CODE IN THIS BELOW 1 LINE

	// return time.Unix(at.Expires, 0).Before(time.Now()).UTC())
}
