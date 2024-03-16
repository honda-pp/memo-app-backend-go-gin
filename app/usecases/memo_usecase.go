package usecases

import (
	"github.com/honda-pp/memo-app-backend-go-gin/app/interfaces"
	"github.com/honda-pp/memo-app-backend-go-gin/generated"
)

func NewMemoUsecase(memoRepository interfaces.MemoRepositoryInterface) *MemoUsecase {
	return &MemoUsecase{
		MemoRepository: memoRepository,
	}
}

type MemoUsecase struct {
	MemoRepository interfaces.MemoRepositoryInterface
}

// CreateMemo create a new memo
func (u *MemoUsecase) CreateMemo(memo generated.Memo) error {
	return u.MemoRepository.CreateMemo(memo)
}

// DeleteById delete memo by ID
func (u *MemoUsecase) DeleteById(id int) error {
	return u.MemoRepository.DeleteById(id)
}

// FindAll returns a list of memos.
func (u *MemoUsecase) FindAll() ([]generated.MemoListInner, error) {
	return u.MemoRepository.FindAll()
}

// FindById find memo by ID
func (u *MemoUsecase) FindById(id int) (*generated.Memo, error) {
	return u.MemoRepository.FindById(id)
}

// UpdateMemo update memo
func (u *MemoUsecase) UpdateMemo(memo generated.Memo) error {
	return u.MemoRepository.UpdateMemo(memo)
}
