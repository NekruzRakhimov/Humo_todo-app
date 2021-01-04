package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"todo-app/models"
)

type AuthSqlite struct {
	db *sqlx.DB
}

func NewAuthSqlite(db *sqlx.DB) *AuthSqlite {
	return &AuthSqlite{db: db}
}

func (r *AuthSqlite) CreateUser(user models.User) (int64, error) {
	query := fmt.Sprintf("INSERT INTO %s (name, username, password_hash) VALUES (($1), ($2), ($3))", usersTable)
	row, err := r.db.Exec(query, user.Name, user.Username, user.Password)
	id, err := row.LastInsertId()
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (r *AuthSqlite) GetUser(username, password string) (models.User, error) {
	var user models.User


	query := fmt.Sprintf("SELECT * FROM %s WHERE username=($1) AND password_hash=($2)", usersTable)
	err := r.db.QueryRow(query, username, password).Scan(
		&user.Id,
		&user.Name,
		&user.Username,
		&user.Password,
	)
	return user, err
}

