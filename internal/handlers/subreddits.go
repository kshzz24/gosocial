package handlers

import (
	"encoding/json"
	"net/http"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/kshzz24/gosocial/internal/models"
)

type CreateSubredditPayload struct {
	Name        string          `json:"name" binding:"required,min=3,max=50"`
	DisplayName string          `json:"display_name" binding:"required,max=100"`
	Description *string         `json:"description"`
	Rules       json.RawMessage `json:"rules"`
	IsNSFW      bool            `json:"is_nsfw"`
	IsPrivate   bool            `json:"is_private"`
}

type UpdateSubredditPayload struct {
	DisplayName    string          `json:"display_name"`
	Description    *string         `json:"description"`
	Rules          json.RawMessage `json:"rules"` // JSONB
	BannerImageURL *string         `json:"banner_image_url"`
	IconImageURL   *string         `json:"icon_image_url"`
	IsNSFW         bool            `json:"is_nsfw"`
	IsPrivate      bool            `json:"is_private"`
	Flairs         json.RawMessage `json:"flairs"` // JSONB
	RulesUpdatedAt *time.Time      `json:"rules_updated_at"`
}

func CreateSubreddit(c *gin.Context) {
	userID, exists := c.Get("user_id")

	if !exists || userID == "" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Authorization is required",
		})

		return
	}

	userIDInt, ok := userID.(int)
	if !ok {
		c.JSON(500, gin.H{"error": "Invalid user ID format"})
		return
	}

	var payload CreateSubredditPayload

	if err := c.BindJSON(&payload); err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	name := strings.ToLower(payload.Name)

	// Validate name format (only lowercase, numbers, underscores)
	matched, _ := regexp.MatchString("^[a-z0-9_]+$", name)
	if !matched {
		c.JSON(400, gin.H{"error": "Name can only contain lowercase letters, numbers, and underscores"})
		return
	}
	existingByName, err := models.GetSubredditByDisplayName(name)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to check subreddit name"})
		return
	}
	if existingByName != nil {
		c.JSON(409, gin.H{"error": "Subreddit name already exists"})
		return
	}

	// Check if display name already exists (NEW)
	existingByDisplayName, err := models.GetSubredditByDisplayName(payload.DisplayName)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to check display name"})
		return
	}
	if existingByDisplayName != nil {
		c.JSON(409, gin.H{"error": "Subreddit with this display name already exists"})
		return
	}

	subreddit := &models.Subreddit{
		Name:         name,
		DisplayName:  payload.DisplayName,
		Rules:        payload.Rules,
		Description:  payload.Description,
		IsNSFW:       payload.IsNSFW,
		IsPrivate:    payload.IsPrivate,
		CreatedBy:    userIDInt,
		ActiveUsers:  0,
		MembersCount: 1,
	}
	if len(subreddit.Rules) == 0 {
		subreddit.Rules = json.RawMessage(`[]`)
	}
	if len(subreddit.Flairs) == 0 {
		subreddit.Flairs = json.RawMessage(`[]`)
	}
	newSubreddit, err := models.CreateSubreddit(subreddit)
	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(201, gin.H{
		"Success": "Successfully created subreddit",
		"data":    newSubreddit,
	})

}

func GetSubreddit(c *gin.Context) {
	name := c.Param("name")

	subreddit, err := models.GetSubredditByName(name)
	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}
	if subreddit == nil {
		c.JSON(404, gin.H{"error": "Subreddit not found"})
		return
	}
	c.JSON(200, gin.H{
		"Success": "Subreddit found",
		"data":    subreddit,
	})
}

func ListSubreddits(c *gin.Context) {

	var limit, offset int

	pageStr := c.Query("page")
	if pageStr != "" {
		page, _ := strconv.Atoi(pageStr)
		perPage, _ := strconv.Atoi(c.DefaultQuery("per_page", "20"))

		if page < 1 {
			page = 1
		}

		offset = (page - 1) * perPage
		limit = perPage
	} else {
		limit, _ = strconv.Atoi(c.DefaultQuery("limit", "20"))
		offset, _ = strconv.Atoi(c.DefaultQuery("offset", "0"))
	}

	if limit < 1 {
		limit = 20
	}
	if limit > 100 {
		limit = 100
	}
	if offset < 0 {
		offset = 0
	}

	subreddits, err := models.ListSubreddits(limit, offset)

	if err != nil {
		c.JSON(500, gin.H{
			"error": "Internal server error",
		})
		return
	}
	if subreddits == nil {
		subreddits = []*models.Subreddit{}
	}

	c.JSON(200, gin.H{
		"subreddits": subreddits,
		"pagination": gin.H{
			"limit":  limit,
			"offset": offset,
			"count":  len(subreddits),
		},
	})
}

func UpdateSubreddit(c *gin.Context) {
	userID, exists := c.Get("user_id")

	if !exists || userID == "" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Authorization is required",
		})

		return
	}

	userIDInt, ok := userID.(int)
	if !ok {
		c.JSON(500, gin.H{"error": "Invalid user ID format"})
		return
	}

	subredditIDStr := c.Param("id")

	subredditID, err := strconv.Atoi(subredditIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid subreddit ID"})
		return
	}

	existingSubreddit, err := models.GetSubredditByID(subredditID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	if existingSubreddit == nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Subreddit not found",
		})
		return
	}
	if existingSubreddit.CreatedBy != userIDInt {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "You are not authorized",
		})
		return
	}

	var payload UpdateSubredditPayload

	if err = c.BindJSON(&payload); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	updatedSubreddit := &models.Subreddit{
		ID:             existingSubreddit.ID,
		DisplayName:    payload.DisplayName,
		Description:    payload.Description,
		Rules:          payload.Rules,
		BannerImageURL: payload.BannerImageURL,
		IconImageURL:   payload.IconImageURL,
		IsNSFW:         payload.IsNSFW,
		IsPrivate:      payload.IsPrivate,
		Flairs:         payload.Flairs,
		RulesUpdatedAt: payload.RulesUpdatedAt,
	}

	err = models.UpdateSubreddit(updatedSubreddit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"Success": "Subreddit Updated successfully",
	})

}

func DeleteSubreddit(c *gin.Context) {
	userID, exists := c.Get("user_id")

	if !exists || userID == "" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Authorization is required",
		})

		return
	}

	userIDInt, ok := userID.(int)
	if !ok {
		c.JSON(500, gin.H{"error": "Invalid user ID format"})
		return
	}

	subredditIDStr := c.Param("id")

	subredditID, err := strconv.Atoi(subredditIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid subreddit ID"})
		return
	}

	existingSubreddit, err := models.GetSubredditByID(subredditID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	if existingSubreddit == nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Subreddit not found",
		})
		return
	}
	if existingSubreddit.CreatedBy != userIDInt {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "You are not authorized",
		})
		return
	}

	err = models.DeleteSubreddit(existingSubreddit.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"Success": "Subreddit Deleted successfully",
	})
	return

}
