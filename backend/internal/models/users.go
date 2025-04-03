package models

import (
	"backend/internal/types"
	"context"
	"errors"
	"fmt"

	"github.com/jmoiron/sqlx"
)

func CreateUser(db *sqlx.DB, user *types.User) (int, error) {
	rows, err := db.Exec("INSERT INTO users(name) VALUES(?)", user.Name)
	if err != nil {
		return 0, fmt.Errorf("failed insert user: %w", err)
	}

	lastId, err := rows.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("failed get last insert id: %w", err)
	}

	return int(lastId), nil
}

func GetUsers(db *sqlx.DB) ([]types.User, error) {
	var users []types.User
	err := db.Select(&users, "SELECT * FROM users")

	if err != nil {
		return nil, fmt.Errorf("failed to get users: %w", err)
	}

	return users, nil
}

func GetUserById(db *sqlx.DB, id int) (types.User, error) {
	var user types.User
	err := db.Get(&user, "SELECT * FROM users WHERE id = ?", id)

	if err != nil {
		return types.User{}, fmt.Errorf("failed to get user by id %d: %w", id, err)
	}

	return user, nil
}

func UpdateUser(ctx context.Context, db *sqlx.DB, id int, name string) error {
	if db == nil {
			return errors.New("db is nil")
	}

	query := `UPDATE users SET name = :name WHERE id = :id`
	
	params := map[string]interface{}{
			"id":   id,
			"name": name,
	}
	
	result, err := db.NamedExecContext(ctx, query, params)
	if err != nil {
			return fmt.Errorf("failed to update user: %w", err)
	}
	
	rowsAffected, err := result.RowsAffected()
	if err != nil {
			return fmt.Errorf("failed to get rows affected: %w", err)
	}
	
	if rowsAffected == 0 {
			return fmt.Errorf("no user found with id %d", id)
	}
	
	return nil
}

func DeleteUserById(db *sqlx.DB, id int) error {
	result, err := db.Exec("DELETE FROM users WHERE id = ?", id)
	if err != nil {
			return fmt.Errorf("failed to delete user: %w", err)
	}
	
	rowsAffected, err := result.RowsAffected()
	if err != nil {
			return fmt.Errorf("failed to get rows affected: %w", err)
	}
	
	if rowsAffected == 0 {
			return fmt.Errorf("user with id %d not found", id)
	}
	
	return nil
}