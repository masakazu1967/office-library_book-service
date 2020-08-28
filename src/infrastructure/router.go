package infrastructure

import (
	"github.com/gin-gonic/gin"
	"local.package/interfaces/controllers"
)

// Router 経路
var Router *gin.Engine

func init() {
	router := gin.Default()

	writerController := controllers.NewWriterController(NewSQLHandler())

	router.POST("/writers", func(c *gin.Context) { writerController.Create(c) })
	router.GET("/writers", func(c *gin.Context) { writerController.Index(c) })
	router.GET("/writers/:id", func(c *gin.Context) { writerController.Show(c) })

	Router = router
}
