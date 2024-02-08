package routes

import (
 
controller "github.com/dhekerlaadhibi/LearnGo/jwt/controllers"
	"github.com/gin-gonic/gin"
)

func AuthRouter(incomingRoutes *gin.Engine ) {
	//here routes signup and sign in dont need protectionn
incomingRoutes.POST("users/signup" , controller.Signup())
incomingRoutes.POST("users/login" , controller.Login())
}


