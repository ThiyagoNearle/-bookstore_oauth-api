package http

import (
	"net/http"

	atDomain "github.com/ThiyagoNearle/bookstore_oauth-api/src/domain/access_token"
	"github.com/ThiyagoNearle/bookstore_oauth-api/src/service/access_token"
	"github.com/ThiyagoNearle/bookstore_oauth-api/src/utils/errors"
	"github.com/gin-gonic/gin"
)

// 	"github.com/ThiyagoNearle/bookstore_oauth-api/src/service/access_token"
type AccessTokenHandler interface {
	GetById(*gin.Context)
	Create(*gin.Context)
}

type accessTokenHandler struct {
	service access_token.Service //  access_token.Service ia a new interface consist of same method, but that new interface defined in the services, so we call that service interface
	//service variable of type Service interface ( that service interface consist of some methods)
}

// So the controller(handler) function takes the services as the parameter
// The controller (NewHandler) takes the service (access_token.Service)
func NewAccessTokenHandler(service access_token.Service) AccessTokenHandler {
	return &accessTokenHandler{
		service: service,
	}
}

func (handler *accessTokenHandler) GetById(c *gin.Context) {
	accessToken := (c.Param("access_token"))

	data, err := handler.service.GetById(accessToken)
	// service already defined with package name, so we wont need that package name here
	if err != nil {
		c.JSON(err.Status, err)
		return
	}

	c.JSON(http.StatusOK, data)
}

func (handler *accessTokenHandler) Create(c *gin.Context) {
	var request atDomain.AccessTokenRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		restErr := errors.NewsBadRequestError("invalid json body")
		c.JSON(restErr.Status, restErr)
		return

	}

	accessToken, err := handler.service.Create(request)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}
	c.JSON(http.StatusOK, accessToken)

}
