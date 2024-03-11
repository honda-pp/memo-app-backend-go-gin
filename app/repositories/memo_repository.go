package repositories

import (
	"database/sql"

	"github.com/honda-pp/memo-app-backend-go-gin/generated"
)

type MemoRepository struct {
	DB *sql.DB
}

func NewMemoRepository(db *sql.DB) *MemoRepository {
	return &MemoRepository{
		DB: db,
	}
}

// CreateMemo create a new memo
func (r *MemoRepository) CreateMemo(memo generated.Memo) error {
	query := "INSERT INTO memo (title, content, user_id) VALUES ($1, $2, $2)"
	_, err := r.DB.Exec(query, memo.Title, memo.Content, memo.UserId)
	if err != nil {
		return err
	}
	return nil
}

// DeleteById delete memo by ID
func (r *MemoRepository) DeleteById(id int) error {
	query := "DELETE FROM memo WHERE id = $1"
	_, err := r.DB.Exec(query, id)
	if err != nil {
		return err
	}
	return nil
}

// FindAll returns a list of memos.
func (r *MemoRepository) FindAll() ([]generated.Memo, error) {
	query := "SELECT id, title, content FROM memo"

	memos := []generated.Memo{}

	rows, err := r.DB.Query(query)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		memo := generated.Memo{}
		err = rows.Scan(&memo.Id, &memo.Title, &memo.Content)
		if err != nil {
			return nil, err
		}
		memos = append(memos, memo)
	}
	return memos, nil
}

// FindById find memo by ID
func (r *MemoRepository) FindById(id int) (*generated.Memo, error) {
	query := "SELECT id, title, content, user_id FROM memo WHERE id = $1"

	memo := generated.Memo{}

	err := r.DB.QueryRow(query, id).Scan(
		&memo.Id,
		&memo.Title,
		&memo.Content,
		&memo.UserId,
	)
	if err != nil {
		return nil, err
	}
	return &memo, nil
}

// UpdateMemo update memo
func (r *MemoRepository) UpdateMemo(memo generated.Memo) error {
	query := "UPDATE memo SET title = $1, content = $2 WHERE id = $3"
	_, err := r.DB.Exec(query, memo.Title, memo.Content, memo.Id)
	if err != nil {
		return err
	}
	return nil
}
