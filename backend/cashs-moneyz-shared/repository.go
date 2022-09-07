package cashs_moneyz_shared

import (
	"database/sql"
)

type Repository interface {
	All() ([]interface{}, error)
	GetByName(city string) (interface{}, error)
	Update(id int64, updated interface{}) (interface{}, error)
}

type SQLiteRepository struct {
	db *sql.DB
}

func NewSQLiteRepo(db *sql.DB) *SQLiteRepository {
	return &SQLiteRepository{
		db: db,
	}
}

func (r *SQLiteRepository) All() ([]interface{}, error) {

	return nil, nil
}

func (r *SQLiteRepository) GetById(id int64) (*interface{}, error) {
	return nil, nil
}
