package db

import (
	"github.com/gocql/gocql"
	"github.com/sovernut/bookstore_oauth-api/src/clients/cassandra"
	"github.com/sovernut/bookstore_oauth-api/src/domain/access_token"
	errors "github.com/sovernut/bookstore_utils-go/rest_errors"
)

const (
	queryGetAccessToken    = "SELECT access_token,user_id,client_id,expires FROM access_tokens WHERE access_token = ?"
	queryCreateAccessToken = "INSERT INTO access_tokens(access_token,user_id,client_id,expires) values (?,?,?,?)"
	queryUpdateAccessToken = "UPDATE access_tokens SET expires=? where access_token = ?"
)

func NewRepository() DbRepository {
	return &dbRepository{}
}

type DbRepository interface {
	GetById(string) (*access_token.AccessToken, *errors.RestErr)
	Create(access_token.AccessToken) *errors.RestErr
	UpdateExpirationTime(access_token.AccessToken) *errors.RestErr
}

type dbRepository struct {
}

func (r *dbRepository) GetById(accessTokenId string) (*access_token.AccessToken, *errors.RestErr) {

	var resultAccessToken access_token.AccessToken
	if err := cassandra.GetSession().Query(queryGetAccessToken, accessTokenId).Scan(
		&resultAccessToken.AccessToken,
		&resultAccessToken.UserId,
		&resultAccessToken.ClientId,
		&resultAccessToken.Expires,
	); err != nil {
		if err == gocql.ErrNotFound {
			return nil, errors.NewNotFoundError("no access token with given id")

		}
		return nil, errors.NewInternalServerError("no access token with given id", err)
	}
	return &resultAccessToken, nil
}

func (r *dbRepository) Create(at access_token.AccessToken) *errors.RestErr {

	if err := cassandra.GetSession().Query(queryCreateAccessToken,
		at.AccessToken,
		at.UserId,
		at.ClientId,
		at.Expires).Exec(); err != nil {
		return errors.NewInternalServerError("error while creating AccessToken", err)
	}
	return nil
}

func (r *dbRepository) UpdateExpirationTime(at access_token.AccessToken) *errors.RestErr {

	if err := cassandra.GetSession().Query(queryUpdateAccessToken,
		at.Expires,
		at.AccessToken,
	).Exec(); err != nil {
		return errors.NewInternalServerError("error while Updating ExpirationTime", err)
	}
	return nil
}
