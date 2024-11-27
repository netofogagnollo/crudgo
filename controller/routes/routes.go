package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/netofogagnollo/crudgo/controller"
)

func InitRoutes(r *gin.RouterGroup, userController controller.UserControllerInterface) {

	r.GET("/getUserByid/:userId", userController.FindUserById)
	r.GET("/getUserByemail/:userId", userController.FindUserByEmail)
	r.POST("/createUser", userController.CreateUser)
	r.PUT("/updateUser/:userId", userController.UpdateUser)
	r.DELETE("/deleteUser/:userId", userController.DeleteUser)

}
