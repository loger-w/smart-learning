package main

import (
	"log"
	"os"
	"strings"

	"smart-learning-backend/pkg/database"
	"smart-learning-backend/pkg/handlers"
	"smart-learning-backend/pkg/middleware"
	"smart-learning-backend/pkg/repositories"
	"smart-learning-backend/pkg/services"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// 載入環境變數
	if err := godotenv.Load(); err != nil {
		log.Println("⚠️ 沒有找到 .env 檔案，使用系統環境變數")
	}

	// 設置 Gin 模式
	ginMode := os.Getenv("GIN_MODE")
	if ginMode != "" {
		gin.SetMode(ginMode)
	}

	// 建立資料庫連接
	db, err := database.NewPostgresConnection()
	if err != nil {
		log.Fatalf("❌ 資料庫連接失敗: %v", err)
	}
	defer db.Close()

	// 測試資料庫連接
	if err := db.TestConnection(); err != nil {
		log.Fatalf("❌ 資料庫測試失敗: %v", err)
	}

	// 顯示連接池統計
	stats := db.GetStats()
	log.Printf("📊 連接池統計 - 最大開啟連接: %d, 開啟連接: %d, 使用中連接: %d, 閒置連接: %d",
		stats.MaxOpenConnections, stats.OpenConnections, stats.InUse, stats.Idle)

	// 初始化依賴注入
	userRepo := repositories.NewUserRepository(db.DB)
	authService := services.NewAuthService(userRepo)
	authHandler := handlers.NewAuthHandler(authService)

	// 初始化 Gin 路由器
	r := gin.Default()

	// 設置信任的代理伺服器
	trustedProxies := os.Getenv("TRUSTED_PROXIES")
	if trustedProxies == "" {
		// 開發環境：不信任任何代理（最安全）
		if err := r.SetTrustedProxies(nil); err != nil {
			log.Fatalf("❌ 設置代理信任失敗: %v", err)
		}
		log.Println("🔒 代理設置：不信任任何代理（開發模式）")
	} else {
		// 生產環境：根據環境變數設置信任的代理
		proxies := []string{}
		if trustedProxies != "none" {
			// 將逗號分隔的 IP 轉換為字串陣列
			for _, proxy := range strings.Split(trustedProxies, ",") {
				proxies = append(proxies, strings.TrimSpace(proxy))
			}
		}
		if err := r.SetTrustedProxies(proxies); err != nil {
			log.Fatalf("❌ 設置代理信任失敗: %v", err)
		}
		if len(proxies) == 0 {
			log.Println("🔒 代理設置：不信任任何代理")
		} else {
			log.Printf("🔒 代理設置：信任的代理 %v", proxies)
		}
	}

	// 添加中介軟體
	r.Use(middleware.CORSMiddleware())

	// 健康檢查端點
	r.GET("/health", func(c *gin.Context) {
		currentStats := db.GetStats()
		c.JSON(200, gin.H{
			"status":  "ok",
			"message": "Smart Learning API is running",
			"db_stats": gin.H{
				"max_open_connections": currentStats.MaxOpenConnections,
				"open_connections":     currentStats.OpenConnections,
				"in_use":               currentStats.InUse,
				"idle":                 currentStats.Idle,
			},
		})
	})

	// API 路由群組
	api := r.Group("/api/v1")
	{
		// 認證路由
		auth := api.Group("/auth")
		{
			auth.POST("/register", authHandler.Register)
			auth.POST("/login", authHandler.Login)
			auth.POST("/logout", middleware.AuthMiddleware(), authHandler.Logout)
			auth.GET("/me", middleware.AuthMiddleware(), authHandler.GetMe)
		}

		// 測試端點
		api.GET("/ping", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "pong",
			})
		})
	}

	// 獲取端口
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("🚀 伺服器啟動在端口 %s", port)
	log.Printf("🌐 健康檢查: http://localhost:%s/health", port)
	log.Printf("📡 API 端點: http://localhost:%s/api/v1/ping", port)
	log.Printf("🔐 認證端點:")
	log.Printf("   註冊: POST http://localhost:%s/api/v1/auth/register", port)
	log.Printf("   登入: POST http://localhost:%s/api/v1/auth/login", port)
	log.Printf("   登出: POST http://localhost:%s/api/v1/auth/logout", port)
	log.Printf("   用戶資料: GET http://localhost:%s/api/v1/auth/me", port)

	// 啟動伺服器
	if err := r.Run(":" + port); err != nil {
		log.Fatalf("❌ 伺服器啟動失敗: %v", err)
	}
}