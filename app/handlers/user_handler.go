package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/honda-pp/memo-app-backend-go-gin/generated"
)

func NewUsersHandler() generated.UsersHandlerInterface {
	return &UsersHandler{}
}

type UsersHandler struct{}

// Delete /user/:id
// Delete user by ID
func (api *UsersHandler) DeleteUserById(c *gin.Context) {
	// Your handler implementation
	c.JSON(200, gin.H{"status": "OK"})
}

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
