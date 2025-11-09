package models

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/kshzz24/gosocial/internal/database"
)

type Post struct {
	ID           int       `json:"id"`
	Title        string    `json:"title"`
	Content      *string   `json:"content"`
	PostType     string    `json:"post_type"`
	LinkURL      *string   `json:"link_url"`
	ImageURL     *string   `json:"image_url"`
	AuthorID     int       `json:"author_id"`
	SubredditID  int       `json:"subreddit_id"`
	Upvotes      int       `json:"upvotes"`
	Downvotes    int       `json:"downvotes"`
	Score        int       `json:"score"`
	CommentCount int       `json:"comment_count"`
	IsLocked     bool      `json:"is_locked"`
	IsNSFW       bool      `json:"is_nsfw"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

// CreatePost creates a new post
func CreatePost(post *Post) (*Post, error) {
	// TODO: Implement
	query := `
		INSERT INTO posts (
			title, content, post_type, link_url, image_url,
			author_id, subreddit_id, is_locked, is_nsfw
		)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
		RETURNING id, created_at, updated_at
	`

	err := database.DB.QueryRow(
		query,
		post.Title,
		post.Content,
		post.PostType,
		post.LinkURL,
		post.ImageURL,
		post.AuthorID,
		post.SubredditID,
		post.IsLocked,
		post.IsNSFW,
	).Scan(&post.ID, &post.CreatedAt, &post.UpdatedAt)

	if err != nil {
		return nil, fmt.Errorf("failed to create post: %w", err)
	}

	// Set default values that database assigned
	post.Upvotes = 0
	post.Downvotes = 0
	post.Score = 0
	post.CommentCount = 0

	return post, nil

}

// GetPostByID retrieves a post by ID
func GetPostByID(id int) (*Post, error) {

	query := `SELECT id, title, content, post_type, link_url, image_url,
			author_id, subreddit_id, upvotes, downvotes, score, comment_count, is_locked, is_nsfw, created_at, updated_at FROM posts WHERE id = $1 `

	post := &Post{}
	err := database.DB.QueryRow(query, id).Scan(
		&post.ID,
		&post.Title,
		&post.Content,
		&post.PostType,
		&post.LinkURL,
		&post.ImageURL,
		&post.AuthorID,
		&post.SubredditID,
		&post.Upvotes,
		&post.Downvotes,
		&post.Score,
		&post.CommentCount,
		&post.IsLocked,
		&post.IsNSFW,
		&post.CreatedAt,
		&post.UpdatedAt,
	)

	if err == sql.ErrNoRows {
		return nil, nil
	}

	if err != nil {
		return nil, fmt.Errorf("failed to get post: %w", err)
	}

	return post, nil
	// TODO: Implement
}

// ListPosts retrieves posts with pagination and optional filters
func ListPosts(limit, offset int, subredditID *int) ([]*Post, error) {
	query := `
		SELECT id, title, content, post_type, link_url, image_url,
		       author_id, subreddit_id, upvotes, downvotes, score,
		       comment_count, is_locked, is_nsfw, created_at, updated_at
		FROM posts
	`

	var rows *sql.Rows
	var err error

	if subredditID != nil {
		query += `WHERE subreddit_id = $1 ORDER BY score DESC LIMIT $2 OFFSET $3`
		rows, err = database.DB.Query(query, *subredditID, limit, offset)
	} else {
		query += `ORDER BY score DESC LIMIT $1 OFFSET $2`
		rows, err = database.DB.Query(query, limit, offset)
	}

	if err != nil {
		return nil, fmt.Errorf("failed to list posts: %w", err)
	}
	defer rows.Close()

	posts := []*Post{}

	for rows.Next() {
		p := &Post{}
		err := rows.Scan(
			&p.ID,
			&p.Title,
			&p.Content,
			&p.PostType,
			&p.LinkURL,
			&p.ImageURL,
			&p.AuthorID,
			&p.SubredditID,
			&p.Upvotes,
			&p.Downvotes,
			&p.Score,
			&p.CommentCount,
			&p.IsLocked,
			&p.IsNSFW,
			&p.CreatedAt,
			&p.UpdatedAt,
		)

		if err != nil {
			return nil, fmt.Errorf("failed to scan post: %w", err)
		}

		posts = append(posts, p)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating posts: %w", err)
	}

	return posts, nil
}

// UpdatePost updates a post
func UpdatePost(post *Post) error {
	// TODO: Implement

	query := `UPDATE posts SET title=$1, content=$2, post_type=$3, link_url=$4, image_url=$5,  is_locked=$6, is_nsfw=$7,updated_at = CURRENT_TIMESTAMP where id=$8`

	rows, err := database.DB.Exec(query, post.Title, post.Content, post.PostType, post.LinkURL, post.ImageURL, post.IsLocked, post.IsNSFW, post.ID)
	// defer rows.Close()
	if err != nil {
		return fmt.Errorf("failed to get post: %w", err)
	}

	log.Println(rows.RowsAffected())
	return nil
}

// DeletePost deletes a post
func DeletePost(id int) error {
	// TODO: Implement
	query := `DELETE FROM posts where id=$1`
	rows, err := database.DB.Exec(query, id)
	// defer rows.Close()
	if err != nil {
		return fmt.Errorf("failed to get post: %w", err)
	}
	log.Println(rows.RowsAffected())
	return nil

}

// UpdatePostScore updates upvotes/downvotes/score
func UpdatePostScore(id, upvotes, downvotes int) error {
	// TODO: Implement
	query := `UPDATE posts SET upvotes=$1, downvotes=$2, score=$3 where id=$4`
	score := upvotes - downvotes
	rows, err := database.DB.Exec(query, upvotes, downvotes, score, id)
	// defer rows.Close()
	if err != nil {
		return fmt.Errorf("failed to get post: %w", err)
	}

	log.Println(rows.RowsAffected())
	return nil
}
