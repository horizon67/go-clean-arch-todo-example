package infrastructure

import (
	gin "github.com/gin-gonic/gin"

	"github.com/horizon67/go-clean-arch-todo-example/app/interfaces/controllers"
)

var Router *gin.Engine

func init() {
	router := gin.Default()
	router.Use(gin.Logger())

	todoController := controllers.NewTodoController(NewSqlHandler())

	router.GET("/todos", func(c *gin.Context) { todoController.Index(c) })
	router.GET("/todos/:id", func(c *gin.Context) { todoController.Show(c) })
	router.POST("/todos", func(c *gin.Context) { todoController.Create(c) })
	router.PUT("/todos/:id", func(c *gin.Context) { todoController.Save(c) })
	router.DELETE("/todos/:id", func(c *gin.Context) { todoController.Delete(c) })

	Router = router
}
