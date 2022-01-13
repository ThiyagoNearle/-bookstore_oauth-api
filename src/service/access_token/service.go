package access_token

import (
	"strings"

	"github.com/ThiyagoNearle/bookstore_oauth-api/src/domain/access_token"
	"github.com/ThiyagoNearle/bookstore_oauth-api/src/repository/db"
	"github.com/ThiyagoNearle/bookstore_oauth-api/src/repository/rest"
	"github.com/ThiyagoNearle/bookstore_oauth-api/src/utils/errors"
)

// both
// SERVICE interface & REPOSITORY interface shold be same ( all methods are same)

// more over DB Repository interface methods should be same to these 2 interface methods

type Service interface {
	GetById(string) (*access_token.AccessToken, *errors.RestErr)
	Create(access_token.AccessTokenRequest) (*access_token.AccessToken, *errors.RestErr)
	UpdateExpirationTime(access_token.AccessToken) *errors.RestErr
}

type service struct {
	restUserRepo rest.RestUserRepository // Repositary this is a database repositary
	dbRepo       db.DbRepository
}

func NewService(usersRepo rest.RestUserRepository, dbRepo db.DbRepository) Service { // Service interface can hold any type
	return &service{
		restUserRepo: usersRepo,
		dbRepo:       dbRepo,
	}
}

func (s *service) GetById(accessTokenId string) (*access_token.AccessToken, *errors.RestErr) {
	accessTokenId = strings.TrimSpace(accessTokenId)
	if len(accessTokenId) == 0 {
		return nil, errors.NewsBadRequestError("invalid access token id")
	}
	accessToken, err := s.dbRepo.GetById(accessTokenId) // if you click GetById method it goes to inthis file Repository interface but internally it goes to the db repository because both consists of same methods..

	if err != nil {
		return nil, err
	}
	return accessToken, nil
}

func (s *service) Create(request access_token.AccessTokenRequest) (*access_token.AccessToken, *errors.RestErr) {

	if err := request.Validate(); err != nil {
		return nil, err
	}

	//TODO: Support both grant types: client_credentials and password

	// AUTHENTICATE THE USER AGAINST THE Users API

	user, err := s.restUserRepo.LoginUser(request.Username, request.Password)
	if err != nil { // Here we validate the access token
		return nil, err
	}

	// GENERATE A NEW ACCESS TOKEN
	at := access_token.GetNewAccessToken(user.Id)
	at.Generate()

	// Save the new access token in Cassandra:
	if err := s.dbRepo.Create(at); err != nil {
		return nil, err
	}
	return &at, nil
}

func (s *service) UpdateExpirationTime(at access_token.AccessToken) *errors.RestErr {
	if err := at.Validate(); err != nil {
		return err
	}
	return s.dbRepo.UpdateExpirationTime(at)
}
