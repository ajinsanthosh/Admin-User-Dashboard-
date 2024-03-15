package routes

import (
	"example/pkg/api/handler"
	"example/pkg/api/middlewire"

	"github.com/gin-gonic/gin"
)

func UserRoutes(engin *gin.RouterGroup,
	user *handler.UserHandler) {

	engin.Use(middlewire.ClearChache)

	engin.GET("/", user.HandlerGetHome)

	engin.GET("/signup", user.HandlerGetUserSignup)
	engin.POST("/signup", user.HandlerUserSignup)

	engin.GET("/login", user.HandlerGetLogin)
	engin.POST("/login", user.HandlerPostLogin)

	engin.GET("/logout", user.HandlerPostLogout)

}
