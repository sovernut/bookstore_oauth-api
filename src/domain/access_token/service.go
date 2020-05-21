package access_token

import (
	"strings"

	errors "github.com/sovernut/bookstore_oauth-api/src/utils/error"
)

type Repository interface {
	GetById(string) (*AccessToken, *errors.RestErr)
	Create(AccessToken) *errors.RestErr
	UpdateExpirationTime(AccessToken) *errors.RestErr
}

type Service interface {
	GetById(string) (*AccessToken, *errors.RestErr)
	Create(AccessToken) *errors.RestErr
	UpdateExpirationTime(AccessToken) *errors.RestErr
}

type service struct {
	repository Repository
}

func NewService(repo Repository) Service {
	return &service{
		repository: repo,
	}
}

func (s *service) GetById(accessTokenId string) (*AccessToken, *errors.RestErr) {
	accessTokenId = strings.TrimSpace(accessTokenId)
	if len(accessTokenId) == 0 {
		return nil, errors.NewBadRequestError("accesstoken length = 0")
	}
	accessToken, err := s.repository.GetById(accessTokenId)
	if err != nil {
		return nil, err
	}
	return accessToken, nil
}

func (s *service) Create(at AccessToken) *errors.RestErr {
	if err := at.Validate(); err != nil {
		return err
	}

	if err := s.repository.Create(at); err != nil {
		return err
	}
	return nil
}

func (s *service) UpdateExpirationTime(at AccessToken) *errors.RestErr {
	if err := at.Validate(); err != nil {
		return err
	}

	if err := s.repository.UpdateExpirationTime(at); err != nil {
		return err
	}
	return s.repository.UpdateExpirationTime(at)
}
