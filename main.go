package main

import (
	"gin-crud-postgres-gorm/middelware""
	"gin-crud-postgres-gorm/router"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.Use(middelware.Logger())
	router.InitRoutes(r)
	r.Run(":8080")
}
