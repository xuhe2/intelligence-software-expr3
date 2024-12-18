package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/xuhe2/intelligence-software-expr3/initializers"
	"github.com/xuhe2/intelligence-software-expr3/models"
	"github.com/xuhe2/intelligence-software-expr3/utils"
)

func Login(c *gin.Context) {
	var body struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	if err := c.BindJSON(&body); err != nil {
		c.JSON(400, gin.H{
			"error": "Invalid request body",
		})
		return
	}
	var user models.User
	if err := initializers.Db.Where(&models.User{Username: body.Username, Password: body.Password}).First(&user).Error; err != nil {
		c.JSON(401, gin.H{
			"error": "Invalid username or password",
		})
		return
	}
	jwt, err := utils.CreateJWT(user.Username)
	if err != nil {
		c.JSON(401, gin.H{
			"error": "Invalid username or password",
		})
		return
	}
	c.JSON(200, gin.H{
		"jwt": jwt,
	})
}
