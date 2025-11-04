package models

import (
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/kshzz24/gosocial/internal/database"
	"github.com/kshzz24/gosocial/internal/utils"
)

type User struct {
	ID           int       `json:"id"`
	Username     string    `json:"username"`
	Email        string    `json:"email"`
	PasswordHash string    `json:"-"` // "-" means never include in JSON
	AvatarURL    *string   `json:"avatar_url"`
	Bio          *string   `json:"bio"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

func CreateUser(username, email, password string) (*User, error) {

	hashedPassword, err := utils.HashPassword(password)
	if err != nil {
		return nil, fmt.Errorf("error hashing password: %w", err)
	}

	userInsertQuery := `INSERT INTO users (username, email, password_hash)
VALUES ($1, $2, $3)
RETURNING id, username, email, avatar_url, bio, created_at, updated_at`
	user := &User{}
	err = database.DB.QueryRow(userInsertQuery, username, email, hashedPassword).Scan(
		&user.ID,
		&user.Username,
		&user.Email,
		&user.AvatarURL,
		&user.Bio,
		&user.CreatedAt,
		&user.UpdatedAt,
	)
	if err != nil {
		return nil, fmt.Errorf("error creating user: %w", err)
	}

	return user, nil
}

func GetUserByEmail(email string) (*User, error) {

	user := &User{}
	query := `
		SELECT id, username, email, password_hash, avatar_url, bio, created_at, updated_at
		FROM users
		WHERE email = $1
	`
	err := database.DB.QueryRow(query, email).Scan(
		&user.ID,
		&user.Username,
		&user.Email,
		&user.PasswordHash,
		&user.AvatarURL,
		&user.Bio,
		&user.CreatedAt,
		&user.UpdatedAt,
	)
	if err != nil {
		return nil, fmt.Errorf("error fetching user: %w", err)
	}

	return user, nil
}

func GetUserByID(id int) (*User, error) {
	user := &User{}
	query := `
		SELECT id, username, email, password_hash, avatar_url, bio, created_at, updated_at
		FROM users
		WHERE id = $1
	`
	err := database.DB.QueryRow(query, id).Scan(
		&user.ID,
		&user.Username,
		&user.Email,
		&user.PasswordHash,
		&user.AvatarURL,
		&user.Bio,
		&user.CreatedAt,
		&user.UpdatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("user not found")
		}
		return nil, fmt.Errorf("error fetching user: %w", err)
	}

	return user, nil

}
