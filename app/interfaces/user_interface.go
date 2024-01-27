package interfaces

import (
	"github.com/gin-gonic/gin"
	"github.com/honda-pp/memo-app-backend-go-gin/generated"
)

type UsersHandlerInterface interface {
	DeleteUserById(c *gin.Context)
	GetUserById(c *gin.Context)
	GetUserList(c *gin.Context)
}

type UserUsecaseInterface interface {
	DeleteById(id int) error
	FindAll() ([]generated.User, error)
	FindById(id int) (generated.User, error)
}

type UserRepositoryInterface interface {
	DeleteById(id int) error
	FindAll() ([]generated.User, error)
	FindById(id int) (generated.User, error)
}
