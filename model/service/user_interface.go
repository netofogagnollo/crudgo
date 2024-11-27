package service

import (
	"github.com/netofogagnollo/crudgo/configuration/rest_err"
	"github.com/netofogagnollo/crudgo/model"
)

func NewUserDomainInterface() UserDomainService {
	return &userDomainService{}
}

type userDomainService struct {
}

type UserDomainService interface {
	CreateUser(model.UserDomainInterface) *rest_err.RestErr
	UpdateUser(string, model.UserDomainInterface) *rest_err.RestErr
	FindUser(string) (*model.UserDomainInterface, *rest_err.RestErr)
	DeleteUser(string) *rest_err.RestErr
}
