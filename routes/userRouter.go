package routes

import(
	controller "github.com/dhekerlaadhibi/LearnGo/jwt/controllers"
	"github.com/dhekerlaadhibi/LearnGo/jwt/middleware"
	"github.com/gin-gonic/gin"
)

func UserRoutes(incomingRoutes *gin.Engine){
	incomingRoutes.Use(middleware.Authenticate())
	incomingRoutes.GET("/users", controller.GetUsers())
	incomingRoutes.GET("/users/:user_id", controller.GetUser())
}