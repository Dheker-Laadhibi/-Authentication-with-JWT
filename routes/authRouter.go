package routes

import (
 
controller "github.com/dhekerlaadhibi/LearnGo/jwt/controllers"
	"github.com/gin-gonic/gin"
)

func AuthRouter(incomingRoutes *gin.Engine ) {
incomingRoutes.POST("users/signup" , controller.signup())
incomingRoutes.POST("users/login" , controller.login())
}
