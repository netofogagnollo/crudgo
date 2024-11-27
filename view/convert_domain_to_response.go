package view

import (
	"github.com/netofogagnollo/crudgo/controller/model/response"
	"github.com/netofogagnollo/crudgo/model"
)

func ConvertDomainToResponse(
	userDomain model.UserDomainInterface,
) response.UserResponse {
	return response.UserResponse{
		ID:    "",
		Email: userDomain.GetEmail(),
		Name:  userDomain.GetName(),
		Age:   userDomain.GetAge(),
	}
}
