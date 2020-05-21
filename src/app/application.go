package app

import (
	"github.com/gin-gonic/gin"
	"github.com/sovernut/bookstore_oauth-api/src/domain/access_token"
	http "github.com/sovernut/bookstore_oauth-api/src/http/access_token"
	"github.com/sovernut/bookstore_oauth-api/src/repository/db"
)

var (
	router = gin.Default()
)

func StartApplication() {
	atHandler := http.NewHandler(access_token.NewService(db.NewRepository()))
	router.GET("/oauth/acess_token/:access_token_id", atHandler.GetById)
	router.POST("/oauth/acess_token/", atHandler.Create)
	router.Run(":8080")
}
