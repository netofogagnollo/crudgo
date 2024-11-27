package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/netofogagnollo/crudgo/configuration/logger"
	"github.com/netofogagnollo/crudgo/configuration/validation"
	"github.com/netofogagnollo/crudgo/controller/model/request"
	"github.com/netofogagnollo/crudgo/model"
	"github.com/netofogagnollo/crudgo/view"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func (uc *userControllerInterface) CreateUser(c *gin.Context) {
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

	domain := model.NewUserDomain(
		userRequest.Email,
		userRequest.Password,
		userRequest.Name,
		userRequest.Age,
	)

	if err := uc.service.CreateUser(domain); err != nil {
		c.JSON(err.Code, err)
		return
	}

	logger.Info("User CreateUser successfully",
		zap.String("joruney", "createUser"))

	c.JSON(http.StatusOK, view.ConvertDomainToResponse(
		domain,
	))

}
