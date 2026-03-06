package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

type Login struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func main() {
	router := gin.Default()

	router.POST("/login", func(c *gin.Context) {
		var login Login

		if err := c.BindJSON(&login); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if login.Username == "admin" && login.Password == "password" {
			c.JSON(http.StatusOK, gin.H{
				"token": "sample-jwt-token",
			})
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "invalid credentials",
			})
		}
	})

	router.Run(":8080")
}
