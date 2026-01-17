package main

import (
	"github.com/gin-gonic/gin"
	"github.com/ricejson/apollo-backend/controller"
	"github.com/ricejson/apollo-backend/models/mongodb"
	_ "github.com/ricejson/apollo-backend/models/mongodb"
	"github.com/ricejson/apollo-backend/repository"
	"github.com/ricejson/apollo-backend/repository/dao"
	"github.com/ricejson/apollo-backend/service/toggle"
)

func main() {
	server := gin.Default()
	col := mongodb.GetClient().Database("apollo").Collection("toggles")
	toggleDAO := dao.NewMongoToggleDAO(col)
	toggleRepository := repository.NewDefaultToggleRepository(toggleDAO)
	toggleService := toggle.NewHTTPToggleService(toggleRepository)
	toggleController := controller.NewToggleController(toggleService)
	toggleController.RegisterServices(server)
	if err := server.Run(":8991"); err != nil {
		panic(err)
	}
}
