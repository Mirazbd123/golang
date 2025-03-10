package helpers

import (
	"errors"

	"github.com/gin-gonic/gin"
)

func CheckUserType(c *gin.Context, role string) (err error) {
	userType := c.GetString("user_type")
	err = nil
	if userType != role {
		err = errors.New("Invalid user type")
		return err
	}
	return err
}

func MatchUserTypeToUid(c *gin.Context, UserId string) (err error) {
	userType := c.GetString("user_type")
	uid := c.GetString("uid")
	err = nil

	if userType == "user" && UserId != uid {
		err = errors.New("Unauthorized access")
		return err
	}
	err = CheckUserType(c, userType)
	return err
}
