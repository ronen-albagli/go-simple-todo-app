package webserver

import (
	"net/http"
	"todo/adapter"
	"todo/config"
	"todo/types"

	"github.com/gin-gonic/gin"

	"todo/usecase"
)

func CreateWebServer() *gin.Engine {
	app := gin.Default()

	app.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusCreated, "{}")
	})

	app.POST("/add", func(c *gin.Context) {
		var todoInput types.TodoEvent

		if err := c.BindJSON(&todoInput); err != nil {
			panic("Failed to parse route input")
		}

		usecaseConfig := new(config.Config)
		todoCollection := usecaseConfig.GetMongoTodoGateway()

		configuration := usecase.Config{
			TodoGateway: &adapter.TodoMongoGateway{
				Collection: *todoCollection,
			},
		}

		useCase := &usecase.AddTodo{
			Config: configuration,
		}

		useCase.Do(todoInput)

		c.Status(201)
	})

	return app
}
