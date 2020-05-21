package rest

import (
	"fmt"
	"net/http"
	"os"
	"testing"

	"github.com/mercadolibre/golang-restclient/rest"
	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	fmt.Println("about to start test case")
	rest.StartMockupServer()
	os.Exit(m.Run())
}

func TestLoginUserTimeoutFromApi(t *testing.T) {
	rest.FlushMockups()
	rest.AddMockups(&rest.Mock{
		HTTPMethod:   http.MethodPost,
		URL:          "http://some.com/users/login",
		ReqBody:      `{"email":"test@gai.com","password":"1234567"}`,
		RespHTTPCode: -1,
		RespBody:     `{}`,
	})

	repo := usersRepository{}
	user, err := repo.LoginUser("test@gai.com", "1234567")
	assert.Nil(t, user)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusInternalServerError, err.Status)
	assert.Equal(t, "invalid restclient response when trying to login", err.Message)
}
func TestLoginUserInvalidErrorInterface(t *testing.T) {
	rest.FlushMockups()
	rest.AddMockups(&rest.Mock{
		HTTPMethod:   http.MethodPost,
		URL:          "http://some.com/users/login",
		ReqBody:      `{"email":"test@gai.com","password":"1234567"}`,
		RespHTTPCode: http.StatusNotFound,
		RespBody:     `{"message": "invalid login credentials when trying to login","status":404,"error":"not_found"}`,
	})

	repo := usersRepository{}
	user, err := repo.LoginUser("test@gai.com", "1234567")
	assert.Nil(t, user)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusNotFound, err.Status)
	assert.Equal(t, "invalid login credentials when trying to login", err.Message)
}
func TestLoginUserInvalidLoginCredentials(t *testing.T) {
	rest.FlushMockups()
	rest.AddMockups(&rest.Mock{
		HTTPMethod:   http.MethodPost,
		URL:          "http://some.com/users/login",
		ReqBody:      `{"email":"test@gai.com","password":"1234567"}`,
		RespHTTPCode: http.StatusNotFound,
		RespBody:     `{"message": "invalid login credentials when trying to login","status":404,"error":"not_found"}`,
	})

	repo := usersRepository{}
	user, err := repo.LoginUser("test@gai.com", "1234567")
	assert.Nil(t, user)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusNotFound, err.Status)
	assert.Equal(t, "invalid login credentials when trying to login", err.Message)
}
func TestLoginUserInvalidUserJsonReponse(t *testing.T) {
	rest.FlushMockups()
	rest.AddMockups(&rest.Mock{
		HTTPMethod:   http.MethodPost,
		URL:          "http://some.com/users/login",
		ReqBody:      `{"email":"test@gai.com","password":"1234567"}`,
		RespHTTPCode: http.StatusOK,
		RespBody:     `{"id" :"1","first_name":"t","last_name":"t","email":"test@hm.com"}`,
	})

	repo := usersRepository{}
	user, err := repo.LoginUser("test@gai.com", "1234567")
	assert.Nil(t, user)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusInternalServerError, err.Status)
	assert.Equal(t, "error when trying to unmarshal user response", err.Message)
}
func TestLoginUserNoError(t *testing.T) {
	rest.FlushMockups()
	rest.AddMockups(&rest.Mock{
		HTTPMethod:   http.MethodPost,
		URL:          "http://some.com/users/login",
		ReqBody:      `{"email":"test@gai.com","password":"1234567"}`,
		RespHTTPCode: http.StatusOK,
		RespBody:     `{"id" :1,"first_name":"t","last_name":"tt","email":"test@hm.com"}`,
	})

	repo := usersRepository{}
	user, err := repo.LoginUser("test@gai.com", "1234567")
	assert.Nil(t, err)
	assert.NotNil(t, user)
	assert.EqualValues(t, 1, user.Id)
	assert.EqualValues(t, "t", user.FirstName)
	assert.EqualValues(t, "tt", user.LastName)
	assert.EqualValues(t, "test@hm.com", user.Email)
}
