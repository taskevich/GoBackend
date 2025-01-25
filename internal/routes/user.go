package routes

import (
	"github.com/gin-gonic/gin"
	v1 "main/internal/handler/http/v1"
)

func RegisterUserRoutes(router *gin.RouterGroup, handler *v1.UserHandler) {
	userRouter := router.Group("/users")
	userRouter.GET("/", handler.GetUsers)
	userRouter.GET("/:id", handler.GetUserById)
	userRouter.POST("/", handler.AddUser)
	userRouter.DELETE("/:id", handler.DeleteUserById)
}
