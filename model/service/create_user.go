package service

import (
	"fmt"

	"github.com/netofogagnollo/crudgo/configuration/logger"
	"github.com/netofogagnollo/crudgo/configuration/rest_err"
	"github.com/netofogagnollo/crudgo/model"
	"go.uber.org/zap"
)

func (ud *userDomainService) CreateUser(
	userDomain model.UserDomainInterface) *rest_err.RestErr {
	logger.Info("init create user model", zap.String("journey", "createUser"))

	userDomain.EncryptPassword()
	fmt.Println(userDomain.GetPassword())

	return nil
}
