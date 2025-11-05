package handlers

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/kshzz24/gosocial/internal/models"
	"github.com/kshzz24/gosocial/internal/utils"
)

type LoginRequestBody struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type RegisterRequestBody struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type ChangePasswordInput struct {
	OldPassword string `json:"old_password" binding:"required"`
	NewPassword string `json:"new_password" binding:"required,min=6"`
}

type ForgotPasswordInput struct {
	Email string `json:"email" binding:"required,email"`
}

type ResetPasswordInput struct {
	Token    string `json:token`
	Password string `json:password`
}

func Register(c *gin.Context) {
	// 1. Define request struct with username, email, password
	// 2. Parse JSON body
	// 3. Validate input (check for errors)
	// 4. Call models.CreateUser()
	// 5. Generate JWT token
	// 6. Return user + token

	var registerBody RegisterRequestBody

	if err := c.BindJSON(&registerBody); err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	if len(registerBody.Password) < 8 {
		c.JSON(400, gin.H{
			"error": "Password must be more than 8 characters",
		})
		return
	}

	existingUser, err := models.GetUserByEmail(registerBody.Email)

	if existingUser != nil {
		c.JSON(409, gin.H{"error": "User already exists"})
		return
	}
	var newUser *models.User

	newUser, err = models.CreateUser(registerBody.Username, registerBody.Email, registerBody.Password)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	token, err := utils.GenerateJWT(newUser.ID, newUser.Username, newUser.Email)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{
		"user":  newUser,
		"token": token,
	})
}

// Login authenticates a user
func Login(c *gin.Context) {
	// 1. Define request struct with email, password
	// 2. Parse JSON body
	// 3. Get user by email
	// 4. Check password with utils.CheckPassword()
	// 5. Generate JWT token
	// 6. Return user + token

	var loginBody LoginRequestBody

	if err := c.BindJSON(&loginBody); err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	var existingUser *models.User
	existingUser, err := models.GetUserByEmail(loginBody.Email)
	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}
	if existingUser == nil {
		c.JSON(401, gin.H{
			"error": "Login or Password is incorrect",
		})
		return
	}

	isValid := utils.CheckPassword(loginBody.Password, existingUser.PasswordHash)

	if !isValid {
		c.JSON(401, gin.H{
			"error": "Login or Password is incorrect",
		})
		return
	}

	token, err := utils.GenerateJWT(existingUser.ID, existingUser.Username, existingUser.Email)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{
		"user": gin.H{
			"id":         existingUser.ID,
			"username":   existingUser.Username,
			"email":      existingUser.Email,
			"created_at": existingUser.CreatedAt,
		},
		"token": token,
	})

}

// GetMe returns current user's profile
func GetMe(c *gin.Context) {
	// 1. Get user_id from context
	// 2. Get user by ID from database
	// 3. Return user

	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(401, gin.H{"error": "Authentication info is not valid"})
		return
	}

	userIDInt, ok := userID.(int)
	if !ok {
		c.JSON(500, gin.H{"error": "Invalid user ID format"})
		return
	}

	user, err := models.GetUserByID(userIDInt)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to fetch user"})
		return
	}

	if user == nil {
		c.JSON(404, gin.H{"error": "User not found"})
		return
	}

	c.JSON(200, gin.H{
		"user": gin.H{
			"id":         user.ID,
			"username":   user.Username,
			"email":      user.Email,
			"created_at": user.CreatedAt,
		},
	})
}

func Logout(c *gin.Context) {

	// handled on client side
	c.JSON(200, gin.H{
		"message": "Logged out successfully",
	})
}

func ChangePassword(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(401, gin.H{
			"error": "You are not authorized",
		})
		return
	}

	userIDInt, ok := userID.(int)
	if !ok {
		c.JSON(500, gin.H{"error": "Invalid user ID format"})
		return
	}

	var changePasswordBody ChangePasswordInput

	if err := c.BindJSON(&changePasswordBody); err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	var exisitingUser *models.User
	exisitingUser, err := models.GetUserByID(userIDInt)

	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	isOldPasswordValid := utils.CheckPassword(changePasswordBody.OldPassword, exisitingUser.PasswordHash)

	if !isOldPasswordValid {
		c.JSON(401, gin.H{
			"error": "Old Password's dont match",
		})
		return
	}

	newPasswordHashed, err := utils.HashPassword(changePasswordBody.NewPassword)
	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = models.UpdatePassword(userIDInt, newPasswordHashed)
	if err != nil {
		c.JSON(500, gin.H{
			"error": "Failed to update password",
		})
		return
	}

	c.JSON(200, gin.H{"message": "Password updated successfully"})
}

func ForgotPassword(c *gin.Context) {
	var forgotPasswordBody ForgotPasswordInput
	if err := c.BindJSON(&forgotPasswordBody); err != nil {
		c.JSON(400, gin.H{"error": "Invalid email format"})
		return
	}

	// Get user by email
	user, err := models.GetUserByEmail(forgotPasswordBody.Email)
	if err != nil {
		c.JSON(500, gin.H{"error": "Internal server error"})
		return
	}

	fmt.Println(user, "current user")

	// If user doesn't exist, return success (security: don't reveal if email exists)
	if user == nil {
		c.JSON(200, gin.H{"message": "If that email exists, a reset link has been sent"})
		return
	}

	// Generate reset token (only if user exists)
	resetToken, err := utils.GenerateResetToken()
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to generate reset token"})
		return
	}

	// Set expiry time (1 hour from now)
	expiresAt := time.Now().Add(1 * time.Hour)

	// Save token to database
	err = models.SaveResetToken(user.ID, resetToken, expiresAt)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to save reset token"})
		return
	}

	// Send email
	err = utils.SendPasswordResetEmail(user.Email, resetToken)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to send reset email"})
		return
	}

	c.JSON(200, gin.H{"message": "If that email exists, a reset link has been sent"})
}
func ResetPassword(c *gin.Context) {
	var resetPasswordBody ResetPasswordInput
	if err := c.BindJSON(&resetPasswordBody); err != nil {
		c.JSON(400, gin.H{"error": "Invalid email format"})
		return
	}

	token := resetPasswordBody.Token

	var user *models.User

	user, err := models.GetUserByResetToken(token)
	if err != nil {
		c.JSON(500, gin.H{"error": "Internal server error"})
		return
	}
	if user == nil {
		c.JSON(400, gin.H{"error": "Invalid or expired reset token"})
		return
	}

	if time.Now().After(*user.ResetTokenExpires) {
		c.JSON(400, gin.H{"error": "Reset token has expired"})
		return
	}
	newPasswordHashed, err := utils.HashPassword(resetPasswordBody.Password)
	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}
	err = models.UpdatePassword(user.ID, newPasswordHashed)
	if err != nil {
		c.JSON(500, gin.H{
			"error": "Failed to update password",
		})
		return
	}
	err = models.ClearResetToken(user.ID)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to clear reset token"})
		return
	}
	c.JSON(200, gin.H{"message": "Password reset successfully"})
}

//vplw ighw fsuz xtzs
