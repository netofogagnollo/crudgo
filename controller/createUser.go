package controller

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/netofogagnollo/crud-go/crud-go/src/configuration/validation"
	"github.com/netofogagnollo/crud-go/crud-go/src/controller/model/request"
)

func CreateUser(c *gin.Context) {
	var userRequest request.UserRequest

	if err := c.ShouldBindBodyWithJSON(&userRequest); err != nil {
		restErr := validation.ValidateUserError(err)
		c.JSON(restErr.Code, restErr)
		return
	}
	fmt.Println(userRequest)
}
