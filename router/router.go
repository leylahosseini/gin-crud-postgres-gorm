package router

import (
	"gin-crud-postgres-gorm/controller"
	_ "gin-crud-postgres-gorm/docs"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func InitRoutes(router *gin.Engine) *gin.Engine {
	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	router.GET("/users", controller.GetAllUsers)
	router.POST("/users", controller.CreateUser)
	router.GET("/users/:id", controller.GetUserByID)
	router.PUT("/users/:id", controller.UpdateUser)
	router.DELETE("/users/:id", controller.DeleteUser)

	return router
}
