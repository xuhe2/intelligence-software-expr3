package middlewares

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/xuhe2/intelligence-software-expr3/initializers"
	"github.com/xuhe2/intelligence-software-expr3/models"
	"gorm.io/gorm"
)

func JWTMiddleware(c *gin.Context) {
	// get the jwt from cookie
	tokenString := c.GetHeader("Authorization")
	if tokenString == "" {
		// if the jwt is empty, just go ahead
		c.Next()
		return
	} else {
		// 去除Bearer
		for strings.Contains(tokenString, "Bearer") {
			tokenString = strings.Replace(tokenString, "Bearer ", "", -1)
		}
	}
	// decode the token and check if it's valid
	// Parse takes the token string and a function for looking up the key. The latter is especially
	// useful if you use multiple keys for your application.  The standard is to use 'kid' in the
	// head of the token to identify which key to use, but the parsed token (head and claims) is provided
	// to the callback, providing flexibility.
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte(os.Getenv("JWT_SECRET")), nil
	})
	if err != nil {
		log.Printf("Error parsing token: %v", err)
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		// if the token is valid
		// if time is expired
		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			c.AbortWithStatus(http.StatusUnauthorized)
		}
		// find the user with name
		var user models.User
		result := initializers.Db.Where(models.User{Username: claims["sub"].(string)}).First(&user)
		// check if the user is exists
		// if user is not exists
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			c.AbortWithStatus(http.StatusUnauthorized)
		}
		// return the user
		c.Set("user", user)
		c.Next()
	} else {
		fmt.Println(err)
	}
}