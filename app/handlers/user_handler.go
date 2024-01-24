package handlers

import (
	"github.com/gin-gonic/gin"
)

type UsersHandler struct{}

// Get /user/:id
// Find user by ID
func (api *UsersHandler) GetUserById(c *gin.Context) {
	// Your handler implementation
	c.JSON(200, gin.H{"status": "OK"})
}

// Get /users
// Returns a list of users.
func (api *UsersHandler) GetUsers(c *gin.Context) {
	// Your handler implementation
	c.JSON(200, gin.H{"status": "OK"})
}
