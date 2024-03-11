package interfaces

import "github.com/honda-pp/memo-app-backend-go-gin/generated"

type MemoUsecaseInterface interface {
	CreateMemo(generated.Memo) error
	DeleteById(id int) error
	FindAll() ([]generated.Memo, error)
	FindById(id int) (*generated.Memo, error)
	UpdateMemo(generated.Memo) error
}

type MemoRepositoryInterface interface {
	CreateMemo(generated.Memo) error
	DeleteById(id int) error
	FindAll() ([]generated.Memo, error)
	FindById(id int) (*generated.Memo, error)
	UpdateMemo(generated.Memo) error
}
