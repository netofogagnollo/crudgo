package controller

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/netofogagnollo/crudgo/configuration/logger"
	"github.com/netofogagnollo/crudgo/configuration/validation"
	"github.com/netofogagnollo/crudgo/controller/model/request"
	"github.com/netofogagnollo/crudgo/model"
	"go.uber.org/zap/zapcore"
)

var (
	UserDomainInterface model.UserDomainInterface
)

func CreateUser(c *gin.Context) {
	logger.Info("Init CreateUser controller",
		zapcore.Field{
			Key:    "joruney",
			String: "createUser",
		},
	)

	var userRequest request.UserRequest

	if err := c.ShouldBindBodyWithJSON(&userRequest); err != nil {
		logger.Error("Error tring to validate user info", err,
			zapcore.Field{
				Key:    "joruney",
				String: "createUser",
			},
		)
		restErr := validation.ValidateUserError(err)
		c.JSON(restErr.Code, restErr)
		return
	}

	domain := model.NewUserDomain(userRequest.Email, userRequest.Password, userRequest.Name, userRequest.Age)
	if err := domain.CreateUser(); err != nil {
		c.JSON(err.Code, err)
		return
	}

	logger.Info("User CreateUser successfully",
		zapcore.Field{
			Key:    "joruney",
			String: "createUser",
		},
	)
	fmt.Println(userRequest)
	// response := response.UserResponse{
	// 	ID:    "test",
	// 	Email: userRequest.Email,
	// 	Name:  userRequest.Name,
	// 	Age:   userRequest.Age,
	// }
	// c.JSON(http.statusOK, response)
}
