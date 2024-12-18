package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/xuhe2/intelligence-software-expr3/initializers"
	"github.com/xuhe2/intelligence-software-expr3/models"
)

func Register(c *gin.Context) {
	var body struct{
		Username string `json:"username" binding:"required"`
		Password string `json:"password" binding:"required"`
	}
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
	}
	// TODO: register user
	user := models.User{
		Username: body.Username,
		Password: body.Password,
	}
	result := initializers.Db.Create(&user)
	if result.Error != nil {
		c.JSON(400, gin.H{"error": result.Error.Error()})
		return
	}
	c.JSON(200, gin.H{"message": "success"})
}