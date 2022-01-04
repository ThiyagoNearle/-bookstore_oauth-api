package access_token

import (
	"strings"

	"github.com/ThiyagoNearle/bookstore_oauth-api/src/utils/errors"
)

type Repositary interface {
	GetById(string) (*AccessToken, *errors.RestErr)
	Create(AccessToken) *errors.RestErr
	UpdateExpirationTime(AccessToken) *errors.RestErr // it is having one method
}

type Service interface {
	GetById(string) (*AccessToken, *errors.RestErr)
	Create(AccessToken) *errors.RestErr
	UpdateExpirationTime(AccessToken) *errors.RestErr
}

type service struct {
	repositary Repositary // Repositary this is a database repositary
}

func NewService(repo Repositary) Service { // Service interface can hold any type
	return &service{
		repositary: repo,
	}
}

func (s *service) GetById(accessTokenId string) (*AccessToken, *errors.RestErr) {
	accessTokenId = strings.TrimSpace(accessTokenId)
	if len(accessTokenId) == 0 {
		return nil, errors.NewsBadRequestError("invalid access token id")
	}
	accessToken, err := s.repositary.GetById(accessTokenId)

	if err != nil {
		return nil, err
	}
	return accessToken, nil
}

func (s *service) Create(at AccessToken) *errors.RestErr {
	if err := at.Validate(); err != nil {
		return err
	}
	return s.repositary.Create(at)
}
func (s *service) UpdateExpirationTime(at AccessToken) *errors.RestErr {
	if err := at.Validate(); err != nil {
		return err
	}
	return s.repositary.UpdateExpirationTime(at)
}