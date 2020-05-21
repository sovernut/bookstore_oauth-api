package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	atDomain "github.com/sovernut/bookstore_oauth-api/src/domain/access_token"
	"github.com/sovernut/bookstore_oauth-api/src/services/access_token"
	errors "github.com/sovernut/bookstore_oauth-api/src/utils/error"
)

type AccessTokenHandler interface {
	GetById(c *gin.Context)
	Create(c *gin.Context)
}

type accessTokenHandler struct {
	service access_token.Service
}

func NewAccessTokenHandler(service access_token.Service) AccessTokenHandler {
	return &accessTokenHandler{
		service: service,
	}
}

func (h *accessTokenHandler) GetById(c *gin.Context) {
	accessToken, err := h.service.GetById(c.Param("access_token_id"))
	if err != nil {
		c.JSON(err.Status, err)
		return
	}
	c.JSON(http.StatusNotImplemented, accessToken)
}

func (h *accessTokenHandler) Create(c *gin.Context) {
	var atr atDomain.AccessTokenRequest
	if err := c.ShouldBindJSON(&atr); err != nil {
		restErr := errors.NewBadRequestError("invalid body request")
		c.JSON(restErr.Status, restErr)
		return
	}
	accessToken, err := h.service.Create(atr)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}
	c.JSON(http.StatusCreated, accessToken)
}
