package main

import (
	"os"

	routes "github.com/dhekerlaadhibi/LearnGo/jwt/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}
	//creation de router
	router := gin.New()
	router.Use(gin.Logger())
	routes.AuthRouter(router)
	routes.UserRoutes(router)
	/* lorsque le serveur reçoit une requête GET sur la route "/api-1", il renvoie une réponse JSON  */
	router.GET("/api-1", func(c *gin.Context) {
		c.JSON(200, gin.H{"success": "access granted for api- 1"})
	})
	router.GET("/api-2", func(c *gin.Context) {
		c.JSON(200, gin.H{"success": "access granted for api -2"})
	})
	router.Run(":" + port)
}
