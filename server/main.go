package main

import (
	"toy-chess/domain/users"
	config "toy-chess/infra/config"
	"toy-chess/infra/database"
	"toy-chess/lib"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
)

func init() {
	lib.GlobalValidatorInstance = validator.New()
}

func main() {
	//Let's go

	config.InitGlobalConfig("server/local-config.json")

	database.InitMongoConnection(config.GlobalConfigInstance.MongoConnectionString)

	//Setup Gin Router and Routes
	router := gin.Default()

	//User Routes
	router.GET("/users", users.GetUsersRoute)
	router.GET("/users/:id", users.GetUserByIdRoute)
	router.POST("/users", users.CreateUserRoute)

	router.Run("localhost:8080")
}
