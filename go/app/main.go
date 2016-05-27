package main

import (
	. "app/controllers"
	_ "app/models"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/static", UsersController.Static)
	router.GET("/users", UsersController.Index)
	router.GET("/users/:id", UsersController.Show)

	router.Run(":3000")
}
