package model

import (
	"fmt"

	"github.com/netofogagnollo/crudgo/configuration/logger"
	"github.com/netofogagnollo/crudgo/configuration/rest_err"
	"go.uber.org/zap"
)

func (ud *UserDomain) CreateUser() *rest_err.RestErr {
	logger.Info("init create user model", zap.String("journey", "createUser"))

	ud.EncryptPassword()
	fmt.Println(ud)

	return nil
}
