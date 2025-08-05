package middleware

import (
	"net/http"
	"smart-learning-backend/pkg/models"
	"smart-learning-backend/pkg/utils"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, models.APIResponse{
				Success: false,
				Message: "未授權",
				Error: &models.APIError{
					Code:    "MISSING_TOKEN",
					Message: "Authorization header is required",
				},
			})
			c.Abort()
			return
		}

		tokenString, err := utils.ExtractTokenFromHeader(authHeader)
		if err != nil {
			c.JSON(http.StatusUnauthorized, models.APIResponse{
				Success: false,
				Message: "未授權",
				Error: &models.APIError{
					Code:    "INVALID_TOKEN_FORMAT",
					Message: "Invalid authorization header format",
				},
			})
			c.Abort()
			return
		}

		claims, err := utils.ValidateJWT(tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, models.APIResponse{
				Success: false,
				Message: "未授權",
				Error: &models.APIError{
					Code:    "INVALID_TOKEN",
					Message: "Token 無效或已過期",
				},
			})
			c.Abort()
			return
		}

		// 將用戶資訊存儲在上下文中
		c.Set("user_id", claims.UserID)
		c.Set("email", claims.Email)
		c.Set("username", claims.Username)

		c.Next()
	}
}