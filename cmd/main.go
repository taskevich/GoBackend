package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
	"main/internal/configs"
	"main/internal/database/postgresql"
	"main/internal/handler/http/docs"
	v1 "main/internal/handler/http/v1"
	"main/internal/repository"
	"main/internal/routes"
	"main/internal/service"
)

func main() {
	config, _ := configs.NewConfig(".env")
	dbConnection := postgresql.NewDatabaseConnection(config)
	dbConnection.InitModels()

	router := gin.New()
	docs.SwaggerInfo.BasePath = "/v1"

	v1Routes := router.Group("/v1")
	{
		userRepository := repository.NewUserRepository(dbConnection.Db)
		userService := service.NewUserService(userRepository)
		userHandler := v1.NewUserHandler(userService)
		routes.RegisterUserRoutes(v1Routes, userHandler)
	}
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	err := router.Run(fmt.Sprintf(":%s", config.Port))
	if err != nil {
		fmt.Println(err.Error())
		return
	}
}
