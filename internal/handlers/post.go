package handlers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kshzz24/gosocial/internal/models"
)

type CreatePostPayload struct {
	Title       string  `json:"title"`
	Content     *string `json:"content"`
	PostType    string  `json:"post_type"`
	LinkURL     *string `json:"link_url"`
	ImageURL    *string `json:"image_url"`
	IsLocked    bool    `json:"is_locked"`
	IsNSFW      bool    `json:"is_nsfw"`
	SubredditID int     `json:"subreddit_id"`
}

// type Post struct {
// 	ID           int       `json:"id"`
// 	Title        string    `json:"title"`
// 	Content      *string   `json:"content"`
// 	PostType     string    `json:"post_type"`
// 	LinkURL      *string   `json:"link_url"`
// 	ImageURL     *string   `json:"image_url"`
// 	AuthorID     int       `json:"author_id"`
// 	SubredditID  int       `json:"subreddit_id"`
// 	Upvotes      int       `json:"upvotes"`
// 	Downvotes    int       `json:"downvotes"`
// 	Score        int       `json:"score"`
// 	CommentCount int       `json:"comment_count"`
// 	IsLocked     bool      `json:"is_locked"`
// 	IsNSFW       bool      `json:"is_nsfw"`
// 	CreatedAt    time.Time `json:"created_at"`
// 	UpdatedAt    time.Time `json:"updated_at"`
// }

func CreatePost(c *gin.Context) {

	userID, exists := c.Get("user_id")
	if !exists || userID == "" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Authorization is required",
		})

		return
	}

	userID_int, ok := userID.(int)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Error while parsing token",
		})

		return
	}

	var payload CreatePostPayload

	if err := c.BindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON payload"})
		return
	}

	newPost := &models.Post{}
	newPost.AuthorID = userID_int
	newPost.Title = payload.Title
	newPost.Content = payload.Content
	newPost.LinkURL = payload.LinkURL
	newPost.ImageURL = payload.ImageURL
	newPost.PostType = payload.PostType
	newPost.IsLocked = payload.IsLocked
	newPost.IsNSFW = payload.IsNSFW
	newPost.SubredditID = payload.SubredditID
	newPost.Upvotes = 0
	newPost.Downvotes = 0
	newPost.CommentCount = 0
	newPost.Score = 0

	newPost, err := models.CreatePost(newPost)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Something went wrong"})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Post Created Successfully"})

}
