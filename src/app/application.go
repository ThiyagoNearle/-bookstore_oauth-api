package app

import (
	"github.com/ThiyagoNearle/bookstore_oauth-api/src/http"
	"github.com/ThiyagoNearle/bookstore_oauth-api/src/repository/db"
	"github.com/ThiyagoNearle/bookstore_oauth-api/src/repository/rest"
	"github.com/ThiyagoNearle/bookstore_oauth-api/src/service/access_token"
	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

func StartApplication() {

	//atService := access_token.NewService(db.NewRepository())
	// access_token service needs a database repository to work ( service function calling (a database repository function))
	// so atService going to interact with the database
	//atHandler := http.NewHandler(atService)

	// handler function receives service function and service function receives db_repository function
	atHandler := http.NewAccessTokenHandler(
		access_token.NewService(rest.NewRestUsersRepository(), db.NewRepository()))

	router.GET("/oauth/access_token/:access_token", atHandler.GetById)
	router.POST("/oauth/access_token", atHandler.Create)
	router.Run(":8080")

} //NewAccessTokenHandler
