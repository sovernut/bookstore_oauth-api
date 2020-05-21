package access_token

import (
	"strings"
	"time"

	errors "github.com/sovernut/bookstore_oauth-api/src/utils/error"
)

const (
	expirationTime = 24
)

type AccessToken struct {
	AccessToken string `json:"access_token"`
	UserId      int64  `json:"user_id"`
	ClientId    int64  `json:"client_id"`
	Expires     int64  `json:"expires"`
}

func GetNewAccessToken() *AccessToken {
	return &AccessToken{
		Expires: time.Now().UTC().Add(expirationTime * time.Hour).Unix(),
	}
}

func (at AccessToken) isExpired() bool {
	now := time.Now().UTC()
	expirationTime := time.Unix(at.Expires, 0)
	return expirationTime.Before(now)
}

func (at AccessToken) Validate() *errors.RestErr {
	at.AccessToken = strings.TrimSpace(at.AccessToken)
	if at.AccessToken == "" {
		return errors.NewBadRequestError("invalid accesstoken ")
	}
	if at.UserId <= 0 {
		return errors.NewBadRequestError("invalid userId ")
	}
	if at.ClientId <= 0 {
		return errors.NewBadRequestError("invalid clientId ")
	}
	if at.Expires <= 0 {
		return errors.NewBadRequestError("invalid expiration time ")
	}
	return nil
}
