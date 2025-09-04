package main

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type PasswordInput struct {
	Password string `json:"password" binding:"required"`
}

func main() {
	r := gin.Default()

	var storedH = make(map[string]string)

	r.POST("/hash-password", func(c *gin.Context) {
		var input PasswordInput
		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
		if err != nil {
			c.JSON(500, gin.H{"error": "Hashing failed"})
			return
		}

		storedH["jerin"] = string(hashedPassword)
		c.JSON(200, gin.H{"message": "Password hashed", "password": string(hashedPassword)})
	})

	r.POST("/verifypassword", func(c *gin.Context) {
		var input PasswordInput
		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		storedHash, exists := storedH["jerin"]
		if !exists {
			c.JSON(404, gin.H{"error": "User not found"})
			return
		}

		err := bcrypt.CompareHashAndPassword([]byte(storedHash), []byte(input.Password))
		if err != nil {
			c.JSON(401, gin.H{"error": "Invalid password"})
			return
		}

		c.JSON(200, gin.H{"message": "Password verified"})
	})

	r.Run(":8080")
}
