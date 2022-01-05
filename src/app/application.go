package app

import (
	"fmt"

	"github.com/ThiyagoNearle/bookstore_oauth-api/src/domain/access_token"
	"github.com/ThiyagoNearle/bookstore_oauth-api/src/http"
	"github.com/ThiyagoNearle/bookstore_oauth-api/src/repository/db"
	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

func StartApplication() {
	// handler using service and use database repository
	atService := access_token.NewService(db.NewRepository())
	atHandler := http.NewHandler(atService)
	fmt.Println("atHandler", atHandler)

	router.GET("/oauth/access_token/:access_token_id", atHandler.GetById)
	router.POST("/oauth/access_token", atHandler.Create)
	router.Run(":8080")

} //NewAccessTokenHandler
