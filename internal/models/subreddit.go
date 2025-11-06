package models

import (
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
func GetSubredditByName(name string) (*Subreddit, error) {
	// TODO: Implement

	// 1. Query subreddit by name
	// 2. Scan all fields (including JSONB)
	// 3. Return nil if not found (sql.ErrNoRows)
	// 4. Return error for other database issues

}

// GetSubredditByID retrieves a subreddit by ID
func GetSubredditByID(id int) (*Subreddit, error) {
	// TODO: Implement
}

// ListSubreddits retrieves all subreddits with pagination
func ListSubreddits(limit, offset int) ([]*Subreddit, error) {
	// TODO: Implement
}

// UpdateSubreddit updates subreddit information
func UpdateSubreddit(subreddit *Subreddit) error {
	// TODO: Implement
}
