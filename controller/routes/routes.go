package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/netofogagnollo/crud-go/crud-go/src/controller"
)

func InitRoutes(r *gin.RouterGroup) {
	r.GET("/getUserByid/:userId", controller.FindUserById)
	r.GET("/getUserByemail/:userId", controller.FindUserByEmail)
	r.POST("/createUser", controller.CreateUser)
	r.PUT("/updateUser/:userId", controller.UpdateUser)
	r.DELETE("/deleteUser/:userId", controller.DeleteUser)

}
