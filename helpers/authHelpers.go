package helpers

import (
	"errors"
	"github.com/gin-gonic/gin"
)

func MatchUserTypeTOUid(c *gin.Context, userId string) (err error) {

	userType := c.GetString("user_type")
	uid := c.GetString("uid")
	if userType == "USER" && uid != userId {
		err = errors.New("unauthorized to access the ressource ")
		return err
	}
	err = CheckUserType(c, userType)
	return err
}
