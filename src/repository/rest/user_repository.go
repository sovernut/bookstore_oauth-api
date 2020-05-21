package rest

import (
	"fmt"
	"time"

	"encoding/json"

	"github.com/mercadolibre/golang-restclient/rest"
	"github.com/sovernut/bookstore_oauth-api/src/domain/users"
	errors "github.com/sovernut/bookstore_oauth-api/src/utils/error"
)

var (
	usersRestClient = rest.RequestBuilder{
		BaseURL: "http://some.com",
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

func (r *usersRepository) LoginUser(email, password string) (*users.User, *errors.RestErr) {
	request := users.UserLoginRequest{
		Email:    email,
		Password: password,
	}
	response := usersRestClient.Post("/users/login", request)
	if response == nil || response.Response == nil {
		return nil, errors.NewInternalServerError("invalid restclient response when trying to login")
	}
	fmt.Println("Response > ", response)
	if response.StatusCode > 299 {
		var restErr errors.RestErr
		err := json.Unmarshal(response.Bytes(), &restErr)
		if err != nil {
			return nil, errors.NewInternalServerError("invalid login credentials when trying to login")
		}
		return nil, &restErr
	}

	var user users.User
	if err := json.Unmarshal(response.Bytes(), &user); err != nil {
		return nil, errors.NewInternalServerError("error when trying to unmarshal user response")
	}
	return &user, nil
}
