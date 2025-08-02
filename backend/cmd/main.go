package main

import (
	"log"
	"os"

	"smart-learning-backend/pkg/database"

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

	// 初始化 Gin 路由器
	r := gin.Default()

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

	// 啟動伺服器
	if err := r.Run(":" + port); err != nil {
		log.Fatalf("❌ 伺服器啟動失敗: %v", err)
	}
}