package app

import (
	"github.com/gin-gonic/gin"
	http "github.com/sovernut/bookstore_oauth-api/src/http/access_token"
	"github.com/sovernut/bookstore_oauth-api/src/repository/db"
	"github.com/sovernut/bookstore_oauth-api/src/repository/rest"
	"github.com/sovernut/bookstore_oauth-api/src/services/access_token"
)

var (
	router = gin.Default()
)

func StartApplication() {
	atHandler := http.NewAccessTokenHandler(
		access_token.NewService(rest.NewRestUsersRepository(), db.NewRepository()))
	router.GET("/oauth/acess_token/:access_token_id", atHandler.GetById)
	router.POST("/oauth/access_token/", atHandler.Create)
	router.Run(":8080")
}
