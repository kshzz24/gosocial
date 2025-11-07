package models

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"time"

	"github.com/kshzz24/gosocial/internal/database"
)

type Subreddit struct {
	ID             int             `json:"id"`
	Name           string          `json:"name"`
	DisplayName    string          `json:"display_name"`
	Description    *string         `json:"description"`
	Rules          json.RawMessage `json:"rules"` // JSONB
	BannerImageURL *string         `json:"banner_image_url"`
	IconImageURL   *string         `json:"icon_image_url"`
	IsNSFW         bool            `json:"is_nsfw"`
	IsPrivate      bool            `json:"is_private"`
	CreatedBy      int             `json:"created_by"`
	MembersCount   int             `json:"members_count"`
	ActiveUsers    int             `json:"active_users"`
	Flairs         json.RawMessage `json:"flairs"` // JSONB
	RulesUpdatedAt *time.Time      `json:"rules_updated_at"`
	CreatedAt      time.Time       `json:"created_at"`
	UpdatedAt      time.Time       `json:"updated_at"`
}

// CreateSubreddit creates a new subreddit
func CreateSubreddit(subreddit *Subreddit) (*Subreddit, error) {

	query := `
	INSERT INTO subreddits (
	  name, display_name, description, rules,
	  banner_image_url, icon_image_url,
	  is_nsfw, is_private, created_by,
	  members_count, active_users,
	  flairs
	)
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12)
	RETURNING id, created_at, updated_at;
	`

	rules := subreddit.Rules
	if len(rules) == 0 {
		rules = json.RawMessage(`[]`)
	}

	flairs := subreddit.Flairs
	if len(flairs) == 0 {
		flairs = json.RawMessage(`[]`)
	}
	err := database.DB.QueryRow(
		query,
		subreddit.Name,
		subreddit.DisplayName,
		subreddit.Description,
		rules,
		subreddit.BannerImageURL,
		subreddit.IconImageURL,
		subreddit.IsNSFW,
		subreddit.IsPrivate,
		subreddit.CreatedBy,
		subreddit.MembersCount,
		subreddit.ActiveUsers,
		flairs,
	).Scan(&subreddit.ID, &subreddit.CreatedAt, &subreddit.UpdatedAt)

	if err != nil {
		return nil, fmt.Errorf("failed to insert subreddit: %w", err)
	}
	subreddit.Rules = rules
	subreddit.Flairs = flairs
	return subreddit, nil

}

// GetSubredditByName retrieves a subreddit by its name
func GetSubredditByDisplayName(name string) (*Subreddit, error) {
	// TODO: Implement
	query := `
		SELECT id, name, display_name, description, rules,
		       banner_image_url, icon_image_url, is_nsfw, is_private,
		       created_by, members_count, active_users, flairs,
		       rules_updated_at, created_at, updated_at
		FROM subreddits
		WHERE display_name = $1
	`
	// 1. Query subreddit by name
	// 2. Scan all fields (including JSONB)
	// 3. Return nil if not found (sql.ErrNoRows)
	// 4. Return error for other database issues
	subreddit := &Subreddit{}
	err := database.DB.QueryRow(query, name).Scan(
		&subreddit.ID,
		&subreddit.Name,
		&subreddit.DisplayName,
		&subreddit.Description,
		&subreddit.Rules,
		&subreddit.BannerImageURL,
		&subreddit.IconImageURL,
		&subreddit.IsNSFW,
		&subreddit.IsPrivate,
		&subreddit.CreatedBy,
		&subreddit.MembersCount,
		&subreddit.ActiveUsers,
		&subreddit.Flairs,
		&subreddit.RulesUpdatedAt,
		&subreddit.CreatedAt,
		&subreddit.UpdatedAt,
	)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, fmt.Errorf("error while fetching subreddit")
	}
	return subreddit, nil
}
func GetSubredditByName(name string) (*Subreddit, error) {
	// TODO: Implement
	query := `
		SELECT id, name, display_name, description, rules,
		       banner_image_url, icon_image_url, is_nsfw, is_private,
		       created_by, members_count, active_users, flairs,
		       rules_updated_at, created_at, updated_at
		FROM subreddits
		WHERE name = $1
	`
	// 1. Query subreddit by name
	// 2. Scan all fields (including JSONB)
	// 3. Return nil if not found (sql.ErrNoRows)
	// 4. Return error for other database issues
	subreddit := &Subreddit{}
	err := database.DB.QueryRow(query, name).Scan(
		&subreddit.ID,
		&subreddit.Name,
		&subreddit.DisplayName,
		&subreddit.Description,
		&subreddit.Rules,
		&subreddit.BannerImageURL,
		&subreddit.IconImageURL,
		&subreddit.IsNSFW,
		&subreddit.IsPrivate,
		&subreddit.CreatedBy,
		&subreddit.MembersCount,
		&subreddit.ActiveUsers,
		&subreddit.Flairs,
		&subreddit.RulesUpdatedAt,
		&subreddit.CreatedAt,
		&subreddit.UpdatedAt,
	)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, fmt.Errorf("error while fetching subreddit")
	}
	return subreddit, nil
}

// GetSubredditByID retrieves a subreddit by ID
func GetSubredditByID(id int) (*Subreddit, error) {
	// TODO: Implement
	query := `
		SELECT id, name, display_name, description, rules,
		       banner_image_url, icon_image_url, is_nsfw, is_private,
		       created_by, members_count, active_users, flairs,
		       rules_updated_at, created_at, updated_at
		FROM subreddits
		WHERE id = $1
	`
	// 1. Query subreddit by name
	// 2. Scan all fields (including JSONB)
	// 3. Return nil if not found (sql.ErrNoRows)
	// 4. Return error for other database issues
	subreddit := &Subreddit{}
	err := database.DB.QueryRow(query, id).Scan(
		&subreddit.ID,
		&subreddit.Name,
		&subreddit.DisplayName,
		&subreddit.Description,
		&subreddit.Rules,
		&subreddit.BannerImageURL,
		&subreddit.IconImageURL,
		&subreddit.IsNSFW,
		&subreddit.IsPrivate,
		&subreddit.CreatedBy,
		&subreddit.MembersCount,
		&subreddit.ActiveUsers,
		&subreddit.Flairs,
		&subreddit.RulesUpdatedAt,
		&subreddit.CreatedAt,
		&subreddit.UpdatedAt,
	)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, fmt.Errorf("error while fetching subreddit")
	}
	return subreddit, nil
}

// ListSubreddits retrieves all subreddits with pagination
func ListSubreddits(limit, offset int) ([]*Subreddit, error) {
	query := `
		SELECT id, name, display_name, description, rules,
		       banner_image_url, icon_image_url, is_nsfw, is_private,
		       created_by, members_count, active_users, flairs,
		       rules_updated_at, created_at, updated_at
		FROM subreddits
		ORDER BY members_count DESC
		LIMIT $1 OFFSET $2
	`

	rows, err := database.DB.Query(query, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("failed to list subreddits: %w", err)
	}
	defer rows.Close()

	subreddits := []*Subreddit{}

	for rows.Next() {
		s := &Subreddit{} // Create new instance (not nil!)

		err := rows.Scan(
			&s.ID,
			&s.Name,
			&s.DisplayName,
			&s.Description,
			&s.Rules,
			&s.BannerImageURL,
			&s.IconImageURL,
			&s.IsNSFW,
			&s.IsPrivate,
			&s.CreatedBy,
			&s.MembersCount,
			&s.ActiveUsers,
			&s.Flairs,
			&s.RulesUpdatedAt,
			&s.CreatedAt,
			&s.UpdatedAt,
		)

		if err != nil {
			return nil, fmt.Errorf("failed to scan subreddit: %w", err)
		}

		subreddits = append(subreddits, s)
	}

	// Check for errors during iteration
	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating subreddits: %w", err)
	}

	return subreddits, nil
}

// UpdateSubreddit updates subreddit information
func UpdateSubreddit(subreddit *Subreddit) error {
	query := `
		UPDATE subreddits 
		SET display_name = $1,
		    description = $2,
		    rules = $3,
		    banner_image_url = $4,
		    icon_image_url = $5,
		    is_nsfw = $6,
		    is_private = $7,
		    flairs = $8,
		    rules_updated_at = $9,
		    updated_at = CURRENT_TIMESTAMP
		WHERE id = $10
	`

	// Set defaults for JSONB if empty
	rules := subreddit.Rules
	if len(rules) == 0 {
		rules = json.RawMessage(`[]`)
	}

	flairs := subreddit.Flairs
	if len(flairs) == 0 {
		flairs = json.RawMessage(`[]`)
	}

	_, err := database.DB.Exec(
		query,
		subreddit.DisplayName,
		subreddit.Description,
		rules,
		subreddit.BannerImageURL,
		subreddit.IconImageURL,
		subreddit.IsNSFW,
		subreddit.IsPrivate,
		flairs,
		subreddit.RulesUpdatedAt,
		subreddit.ID,
	)

	if err != nil {
		return fmt.Errorf("failed to update subreddit: %w", err)
	}

	return nil
}
func DeleteSubreddit(id int) error {
	query := `DELETE FROM subreddits WHERE id = $1 `

	// Set defaults for JSONB if empty

	_, err := database.DB.Exec(
		query,
		id,
	)

	if err != nil {
		return fmt.Errorf("failed to delete subreddit: %w", err)
	}

	return nil
}
