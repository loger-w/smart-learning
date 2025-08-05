package handlers

import (
	"net/http"
	"smart-learning-backend/pkg/models"
	"smart-learning-backend/pkg/services"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type AuthHandler struct {
	authService *services.AuthService
	validator   *validator.Validate
}

func NewAuthHandler(authService *services.AuthService) *AuthHandler {
	return &AuthHandler{
		authService: authService,
		validator:   validator.New(),
	}
}

func (h *AuthHandler) Register(c *gin.Context) {
	var req models.RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		validationErrors := make(map[string][]string)
		
		if validatorErrors, ok := err.(validator.ValidationErrors); ok {
			for _, fieldError := range validatorErrors {
				field := strings.ToLower(fieldError.Field())
				var message string
				
				switch fieldError.Tag() {
				case "required":
					message = "此欄位為必填"
				case "email":
					message = "電子郵件格式不正確"
				case "min":
					if field == "password" {
						message = "密碼至少需要 8 個字符"
					} else if field == "username" {
						message = "用戶名至少需要 2 個字符"
					} else {
						message = "長度不足"
					}
				case "max":
					if field == "username" {
						message = "用戶名不能超過 20 個字符"
					} else {
						message = "長度過長"
					}
				case "alphanum":
					message = "只能包含字母和數字"
				default:
					message = "格式不正確"
				}
				
				validationErrors[field] = append(validationErrors[field], message)
			}
		}
		
		c.JSON(http.StatusBadRequest, models.APIResponse{
			Success: false,
			Message: "驗證失敗",
			Errors:  validationErrors,
		})
		return
	}
	
	authResponse, err := h.authService.Register(&req)
	if err != nil {
		if strings.Contains(err.Error(), "user already exists") {
			c.JSON(http.StatusConflict, models.APIResponse{
				Success: false,
				Message: "用戶已存在",
				Error: &models.APIError{
					Code:    "USER_ALREADY_EXISTS",
					Message: "電子郵件或用戶名已被使用",
				},
			})
			return
		}
		
		if strings.Contains(err.Error(), "passwords do not match") {
			c.JSON(http.StatusBadRequest, models.APIResponse{
				Success: false,
				Message: "驗證失敗",
				Errors: map[string][]string{
					"confirm_password": {"密碼確認不一致"},
				},
			})
			return
		}
		
		if strings.Contains(err.Error(), "username can only contain") {
			c.JSON(http.StatusBadRequest, models.APIResponse{
				Success: false,
				Message: "驗證失敗",
				Errors: map[string][]string{
					"username": {"用戶名只能包含字母、數字和底線"},
				},
			})
			return
		}
		
		c.JSON(http.StatusInternalServerError, models.APIResponse{
			Success: false,
			Message: "註冊失敗",
			Error: &models.APIError{
				Code:    "INTERNAL_SERVER_ERROR",
				Message: "伺服器內部錯誤",
			},
		})
		return
	}
	
	c.JSON(http.StatusCreated, models.APIResponse{
		Success: true,
		Message: "註冊成功",
		Data:    authResponse,
	})
}

func (h *AuthHandler) Login(c *gin.Context) {
	var req models.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		validationErrors := make(map[string][]string)
		
		if validatorErrors, ok := err.(validator.ValidationErrors); ok {
			for _, fieldError := range validatorErrors {
				field := strings.ToLower(fieldError.Field())
				var message string
				
				switch fieldError.Tag() {
				case "required":
					message = "此欄位為必填"
				case "email":
					message = "請輸入有效的電子郵件"
				case "min":
					if field == "password" {
						message = "密碼至少需要 8 個字符"
					} else {
						message = "長度不足"
					}
				default:
					message = "格式不正確"
				}
				
				validationErrors[field] = append(validationErrors[field], message)
			}
		}
		
		c.JSON(http.StatusBadRequest, models.APIResponse{
			Success: false,
			Message: "驗證失敗",
			Errors:  validationErrors,
		})
		return
	}
	
	authResponse, err := h.authService.Login(&req)
	if err != nil {
		c.JSON(http.StatusUnauthorized, models.APIResponse{
			Success: false,
			Message: "登入失敗",
			Error: &models.APIError{
				Code:    "INVALID_CREDENTIALS",
				Message: "電子郵件或密碼錯誤",
			},
		})
		return
	}
	
	c.JSON(http.StatusOK, models.APIResponse{
		Success: true,
		Message: "登入成功",
		Data:    authResponse,
	})
}

func (h *AuthHandler) Logout(c *gin.Context) {
	// 在實際應用中，這裡可以將 token 加入黑名單
	// 目前簡單返回成功響應
	c.JSON(http.StatusOK, models.APIResponse{
		Success: true,
		Message: "登出成功",
	})
}

func (h *AuthHandler) GetMe(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, models.APIResponse{
			Success: false,
			Message: "未授權",
			Error: &models.APIError{
				Code:    "UNAUTHORIZED",
				Message: "無法獲取用戶資訊",
			},
		})
		return
	}
	
	user, err := h.authService.GetUserByID(userID.(int))
	if err != nil {
		c.JSON(http.StatusNotFound, models.APIResponse{
			Success: false,
			Message: "用戶不存在",
			Error: &models.APIError{
				Code:    "USER_NOT_FOUND",
				Message: "用戶不存在",
			},
		})
		return
	}
	
	c.JSON(http.StatusOK, models.APIResponse{
		Success: true,
		Data: map[string]interface{}{
			"user": user,
		},
	})
}