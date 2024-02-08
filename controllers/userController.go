package controllers

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/dhekerlaadhibi/LearnGo/jwt/database"
	helper "github.com/dhekerlaadhibi/LearnGo/jwt/helpers"
	"github.com/dhekerlaadhibi/LearnGo/jwt/models"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

var userCollection *mongo.Collection = database.openCollection(database.Client, "user")
var validate = validator.New()

func HashPassword() {

}
func VerifyPassword() {

}
func Signup() {

}
func Login() {

}
func GetUsers() {

}

func GetUser() gin.HandlerFunc {

	//userRouter incomingRoutes.GET("/users/:user_id", controller.GetUser())
	return func(c *gin.Context) {
		userId := c.Param("user_id")
		//check if user is the admin or not
		helper.MatchUserTypeTOUid(c,userId)
	}

}
