package storage

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
	"github.com/spellsaif/golang-betterme-app/internal/models"
)

type Sqlite struct {
	Db *sql.DB
}

func New() (*Sqlite, error) {
	db, err := sql.Open("sqlite3", "database.db")

	if err != nil {
		return nil, err
	}

	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS users (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			username VARCHAR(255),
			email VARCHAR(255),
			password VARCHAR(255)
		)
	`)

	if err != nil {
		return nil, err
	}

	return &Sqlite{Db: db}, nil
}

func (s *Sqlite) CreateUser(user *models.User) (int64, error) {
	query := `INSERT INTO users (username, email, password) VALUES (?, ?, ?)`

	result, err := s.Db.Exec(query, user.Username, user.Email, user.Password)

	if err != nil {
		return 0, err
	}

	//getting id of newly inserted user

	id, err := result.LastInsertId()

	if err != nil {
		return 0, err
	}

	return id, nil
}

func (s *Sqlite) FindUserByUsername(username string) (*models.User, error) {
	query := `SELECT username, password FROM users WHERE username = ?`

	user := &models.User{}

	row := s.Db.QueryRow(query, username)

	err := row.Scan(&user.Username, &user.Password)

	if err == sql.ErrNoRows {
		return nil, err
	}

	if err != nil {
		return nil, err
	}

	return user, nil
}
