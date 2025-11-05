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
	ID                int        `json:"id"`
	Username          string     `json:"username"`
	Email             string     `json:"email"`
	PasswordHash      string     `json:"-"` // "-" means never include in JSON
	AvatarURL         *string    `json:"avatar_url"`
	Bio               *string    `json:"bio"`
	CreatedAt         time.Time  `json:"created_at"`
	UpdatedAt         time.Time  `json:"updated_at"`
	ResetToken        *string    `json:"-"` // Add this
	ResetTokenExpires *time.Time `json:"-"` // Add this
}

func CreateUser(username, email, password string) (*User, error) {

	hashedPassword, err := utils.HashPassword(password)
	if err != nil {
		return nil, fmt.Errorf("error hashing password: %w", err)
	}

	userInsertQuery := `INSERT INTO users (username, email, password_hash)
VALUES ($1, $2, $3)
RETURNING id, username, email, avatar_url, bio, created_at, updated_at, reset_token, reset_token_expires`
	user := &User{}
	err = database.DB.QueryRow(userInsertQuery, username, email, hashedPassword).Scan(
		&user.ID,
		&user.Username,
		&user.Email,
		&user.AvatarURL,
		&user.Bio,
		&user.CreatedAt,
		&user.UpdatedAt,
		&user.ResetToken,
		&user.ResetTokenExpires,
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

func UpdatePassword(userID int, newPasswordHash string) error {

	query := `UPDATE users SET password_hash = $1, updated_at = CURRENT_TIMESTAMP WHERE id = $2`

	_, err := database.DB.Exec(query, newPasswordHash, userID)
	if err != nil {
		return fmt.Errorf("failed to update password: %w", err)
	}

	return nil
}

func SaveResetToken(userID int, token string, expiresAt time.Time) error {
	query := `UPDATE users SET reset_token = $1, reset_token_expires = $2 WHERE id = $3`
	_, err := database.DB.Exec(query, token, expiresAt, userID)
	if err != nil {
		return fmt.Errorf("failed to save reset token: %w", err)
	}
	return nil
}
func GetUserByResetToken(token string) (*User, error) {
	query := `SELECT id, username, email, password_hash, avatar_url,bio, created_at, updated_at, reset_token, reset_token_expires 
	          FROM users WHERE reset_token = $1`

	var user User
	err := database.DB.QueryRow(query, token).Scan(
		&user.ID,
		&user.Username,
		&user.Email,
		&user.PasswordHash,
		&user.AvatarURL,
		&user.Bio,
		&user.CreatedAt,
		&user.UpdatedAt,
		&user.ResetToken,
		&user.ResetTokenExpires,
	)

	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, fmt.Errorf("failed to get user: %w", err)
	}

	return &user, nil
}

func ClearResetToken(userID int) error {
	query := `UPDATE users SET reset_token = NULL, reset_token_expires = NULL WHERE id = $1`

	_, err := database.DB.Exec(query, userID)
	if err != nil {
		return fmt.Errorf("failed to clear reset token: %w", err)
	}

	return nil
}
