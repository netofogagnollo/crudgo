package service

import (
	"github.com/netofogagnollo/crudgo/configuration/rest_err"
	"github.com/netofogagnollo/crudgo/model"
)

func (ud *userDomainService) FindUser(string) (
	*model.UserDomainInterface,
	*rest_err.RestErr,
) {
	return nil, nil
}
