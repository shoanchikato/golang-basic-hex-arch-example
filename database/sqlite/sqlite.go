package sqlite

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
	"github.com/shoanchikato/demo/hex/toy"
)

// NewSqliteToyRepository injecting the repo to the DB
func NewSqliteToyRepository(db *sql.DB) toy.Repository {
	return &toyRepository{
		db,
	}
}

type toyRepository struct {
	db *sql.DB
}

func (r *toyRepository) Create(toy *toy.Toy) error {
	stmt, err := r.db.Prepare(`INSERT INTO toys (id, name, owner, createdAt, updatedAt) VALUES (?, ?, ?, ?, ?)`)
	if err != nil {
		return nil
	}

	_, err = stmt.Exec(toy.ID, toy.Name, toy.Owner, toy.CreatedAt, toy.UpdatedAt)
	return nil
}

func (r *toyRepository) FindByID(id string) (*toy.Toy, error) {
	row, err := r.db.Query("SELECT id, name, owner, createdAt, updatedAt FROM toys WHERE id = ?", id)
	if err != nil {
		return nil, err
	}

	toy := new(toy.Toy)
	for row.Next() {
		err = row.Scan(&toy.ID, &toy.Name, &toy.Owner, &toy.CreatedAt, &toy.UpdatedAt)
		if err != nil {
			return nil, err
		}
	}
	return toy, nil
}

func (r *toyRepository) FindAll() (toys []*toy.Toy, err error) {
	rows, err := r.db.Query("SELECT id, name, owner, createdAt, updatedAt FROM toys")
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		toy := new(toy.Toy)
		err = rows.Scan(&toy.ID, &toy.Name, &toy.Owner, &toy.CreatedAt, &toy.UpdatedAt)
		if err != nil {
			return nil, err
		}
		toys = append(toys, toy)
	}
	return
}
