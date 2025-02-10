package routes

import (
	"github.com/gin-gonic/gin"
	controller "github.com/mirazbd123/golang-jwt/controllers"
	"github.com/mirazbd123/golang-jwt/middleware"
)

func UserRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.Use(middleware.Authenticate())
	incomingRoutes.GET("/users", controller.GetUsers())
	incomingRoutes.GET("/user/:user_id", controller.GetUser())
}
