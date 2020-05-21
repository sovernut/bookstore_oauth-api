package access_token

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestAcessTokenConstants(t *testing.T) {
	assert.EqualValues(t, 24, expirationTime, "default expiration time is 24")
}

func TestGetNewAccessToken(t *testing.T) {
	at := GetNewAccessToken()
	assert.False(t, at.isExpired(), "newAccessToken is expired!!!!")
	assert.EqualValues(t, "", at.AccessToken, "newAccessToken should not have defined access token id")
	assert.True(t, at.UserId == 0, "newAccessToken should not have accosicated user id")
}

func TestAcessTokenIsExpired(t *testing.T) {
	at := AccessToken{}

	assert.True(t, at.isExpired(), "empty access token should expired by default")

	at.Expires = time.Now().UTC().Add(3 * time.Hour).Unix()
	assert.False(t, at.isExpired(), "access token created 3 hour should'nt expired by default")

}
