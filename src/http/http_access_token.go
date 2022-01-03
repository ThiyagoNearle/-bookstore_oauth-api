package http

import (
	"net/http"

	"github.com/ThiyagoNearle/bookstore_oauth-api/src/domain/access_token"
	"github.com/gin-gonic/gin"
)

type AccessTokenHandler interface {
	GetById(*gin.Context)
}

type accessTokenHandler struct {
	service access_token.Service
}

func NewHandler(service access_token.Service) AccessTokenHandler {
	return &accessTokenHandler{
		service: service,
	}
}

func (handler *accessTokenHandler) GetById(c *gin.Context) {
	accessTokenId := (c.Param("access_token_id"))

	accessToken, err := handler.service.GetById(accessTokenId)

	if err != nil {
		c.JSON(err.Status, err)
		return
	}

	c.JSON(http.StatusNotImplemented, accessToken)
}
