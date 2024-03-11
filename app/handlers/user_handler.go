package handlers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/honda-pp/memo-app-backend-go-gin/app/interfaces"
	"github.com/honda-pp/memo-app-backend-go-gin/generated"
)

func NewUsersHandler(userUsecase interfaces.UserUsecaseInterface) *UsersHandler {
	return &UsersHandler{
		UserUsecase: userUsecase,
	}
}

type UsersHandler struct {
	UserUsecase interfaces.UserUsecaseInterface
}

// Create /user
// Create a new user
func (h *UsersHandler) CreateUser(c *gin.Context) {
	var user generated.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := h.UserUsecase.CreateUser(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}
	c.JSON(200, gin.H{"status": "OK"})
}

// Delete /user/:id
// Delete user by ID
func (h *UsersHandler) DeleteUserById(c *gin.Context) {
	userIDStr := c.Param("id")
	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}
	h.UserUsecase.DeleteById(userID)
	c.JSON(200, gin.H{"status": "OK"})
}

// Get /user/:id
// Find user by ID
func (h *UsersHandler) GetUserById(c *gin.Context) {
	userIDStr := c.Param("id")
	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}
	user, err := h.UserUsecase.FindById(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get user"})
		return
	}
	c.JSON(200, gin.H{"user": user})
}

// Get /user-list
// Returns a list of users.
func (h *UsersHandler) GetUserList(c *gin.Context) {
	users, err := h.UserUsecase.FindAll()

	if err != nil {
		log.Fatal("Failed to get user list: " + err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get user list"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"userList": users})
}
