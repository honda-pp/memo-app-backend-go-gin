/*
 * Memo App API
 *
 * API description in Markdown.
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

import (
	"github.com/gin-gonic/gin"
)

type UsersMemoAppApiInterface interface {
	GetUserById(c *gin.Context)
	GetUsers(c *gin.Context)
}

func NewUsersMemoAppApi() UsersMemoAppApiInterface {
	return &UsersMemoAppApi{}
}

type UsersMemoAppApi struct {
}

// Get /user/:id
// Find user by ID 
func (api *UsersMemoAppApi) GetUserById(c *gin.Context) {
	// Your handler implementation
	c.JSON(200, gin.H{"status": "OK"})
}

// Get /users
// Returns a list of users. 
func (api *UsersMemoAppApi) GetUsers(c *gin.Context) {
	// Your handler implementation
	c.JSON(200, gin.H{"status": "OK"})
}
