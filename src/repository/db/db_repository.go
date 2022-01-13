package db

import (
	"github.com/ThiyagoNearle/bookstore_oauth-api/src/clients/mysql"
	"github.com/ThiyagoNearle/bookstore_oauth-api/src/domain/access_token"
	"github.com/ThiyagoNearle/bookstore_oauth-api/src/utils/errors"
)

const (
	queryGetAccessToken    = "SELECT access_token, user_id, client_id, expires FROM access_tokens WHERE access_token=?;"
	queryCreateAccessToken = "INSERT INTO access_tokens(access_token, user_id, client_id, expires) VALUES(?,?,?,?);"
	queryUpdateExpires     = "UPDATE access_tokens SET expires=? WHERE access_token=?;"
)

func NewRepository() DbRepository {
	return &dbRepository{}
}

// This BRepository interface methods should be same to those 2 service package interface methods.

type DbRepository interface {
	GetById(string) (*access_token.AccessToken, *errors.RestErr)
	Create(access_token.AccessToken) *errors.RestErr
	UpdateExpirationTime(access_token.AccessToken) *errors.RestErr
}

type dbRepository struct {
}

//
func (r *dbRepository) GetById(accessToken string) (*access_token.AccessToken, *errors.RestErr) {

	var result access_token.AccessToken

	stmt, err := mysql.Client.Prepare(queryGetAccessToken)
	if err != nil {
		return nil, errors.NewInternalServerError("database error")
	}
	defer stmt.Close()

	values := stmt.QueryRow(accessToken) // values holds the row// here id is just a name, in this we pass access token
	if getErr := values.Scan(&result.AccessToken, &result.UserId, &result.ClientId, &result.Expires); getErr != nil {
		return nil, errors.NewInternalServerError("database error")
	}

	return &result, nil

}

func (r *dbRepository) Create(at access_token.AccessToken) *errors.RestErr {
	stmt, err := mysql.Client.Prepare(queryCreateAccessToken)
	if err != nil {
		return errors.NewInternalServerError("database error")
	}
	defer stmt.Close()

	_, saveErr := stmt.Exec( // inorder to save in database use Exec(pass the values for ?)
		at.AccessToken, // _ holds the row with values after the execution, here we dont need that because after save in to database, we dont need anything.
		at.UserId,
		at.ClientId,
		at.Expires) // passing values from struct to database
	if saveErr != nil { // stmt.Exec(value) insert values to the database and while executing database automatically generate some numbers for id column
		return errors.NewInternalServerError("database error")
	}
	return nil
}

func (r *dbRepository) UpdateExpirationTime(at access_token.AccessToken) *errors.RestErr {
	stmt, err := mysql.Client.Prepare(queryUpdateExpires)
	if err != nil {
		return errors.NewInternalServerError("database error")
	}
	defer stmt.Close()

	_, err = stmt.Exec(at.AccessToken, at.Expires)
	if err != nil {
		return errors.NewInternalServerError("database error")
	}
	return nil

}

// CASSANDRA SESSION
/*
func (r *dbRepository) GetById(id string) (*access_token.AccessToken, *errors.RestErr) {

	// TODO: IMPLEMENT get access token from CassandraDB

	var result access_token.AccessToken

	if err := cassandra.GetSession().Query(queryGetAccessToken, id).Scan(
		&result.AccessToken,
		&result.UserId,
		&result.ClientId,
		&result.Expires,
	); err != nil {
		if err == gocql.ErrNotFound {
			return nil, errors.NewNotFoundError("no acces token found for the given id")
		}

		return nil, errors.NewInternalServerError(err.Error())

	}

	return &result, nil
}

func (r *dbRepository) Create(at access_token.AccessToken) *errors.RestErr {

	if err := cassandra.GetSession().Query(queryCreateAccessToken,
		at.AccessToken,
		at.UserId,
		at.ClientId,
		at.Expires).Exec(); err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	return nil
}

func (r *dbRepository) UpdateExpirationTime(at access_token.AccessToken) *errors.RestErr {

	if err := cassandra.GetSession().Query(queryUpdateExpires,
		at.AccessToken,
		at.Expires).Exec(); err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	return nil
}


*/
