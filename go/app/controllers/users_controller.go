package controllers

import (
	. "app/models"

	"github.com/gin-gonic/gin"
)

var UsersController *_UsersController

type _UsersController struct{}

func (_ *_UsersController) Index(c *gin.Context) {
	users := User.All()
	c.JSON(200, users)
}

func (_ *_UsersController) Show(c *gin.Context) {
	user := User.Find(c.Param("id"))
	c.JSON(200, user)
}

func (_ *_UsersController) Static(c *gin.Context) {
	c.JSON(200, gin.H{
		"_id":        "57472be59cf6a4c96b088a17",
		"username":   "Josh Williams",
		"created_at": "2016-05-26T17:01:25.532Z",
		"updated_at": "2016-05-26T17:01:25.532Z",
	})
}
