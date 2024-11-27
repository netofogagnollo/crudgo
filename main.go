package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/netofogagnollo/crudgo/configuration/logger"
	"github.com/netofogagnollo/crudgo/controller"
	"github.com/netofogagnollo/crudgo/controller/routes"
	"github.com/netofogagnollo/crudgo/model/service"
)

func main() {
	logger.Info("About to start application")
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loaing .env file")

	}

	//init dependecies
	service := service.NewUserDomainInterface()
	userController := controller.NewUserControllerInterface(service)

	router := gin.Default()

	routes.InitRoutes(&router.RouterGroup, userController)

	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}

}
