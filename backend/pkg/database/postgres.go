package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	_ "github.com/lib/pq"
)

type DB struct {
	*sql.DB
}

// NewPostgresConnection 建立 PostgreSQL 連接
func NewPostgresConnection() (*DB, error) {
	// 從環境變數構建連接字串
	databaseURL := os.Getenv("DATABASE_URL")
	if databaseURL == "" {
		// 如果沒有 DATABASE_URL，則從個別環境變數構建
		host := os.Getenv("DB_HOST")
		port := os.Getenv("DB_PORT")
		user := os.Getenv("DB_USER")
		password := os.Getenv("DB_PASSWORD")
		dbname := os.Getenv("DB_NAME")
		sslmode := os.Getenv("DB_SSL_MODE")

		if sslmode == "" {
			sslmode = "disable"
		}

		databaseURL = fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s",
			user, password, host, port, dbname, sslmode)
	}

	// 開啟資料庫連接
	db, err := sql.Open("postgres", databaseURL)
	if err != nil {
		return nil, fmt.Errorf("failed to open database: %w", err)
	}

	// 設置連接池參數
	db.SetMaxOpenConns(30)                 // 最大開啟連接數
	db.SetMaxIdleConns(5)                  // 最大閒置連接數
	db.SetConnMaxLifetime(time.Hour)       // 連接最大生命週期
	db.SetConnMaxIdleTime(time.Minute * 30) // 連接最大閒置時間

	// 測試連接
	if err := db.Ping(); err != nil {
		db.Close()
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	log.Println("✅ 資料庫連接成功建立")

	return &DB{DB: db}, nil
}

// TestConnection 測試資料庫連接並顯示版本資訊
func (db *DB) TestConnection() error {
	var version string
	err := db.QueryRow("SELECT version()").Scan(&version)
	if err != nil {
		return fmt.Errorf("查詢失敗: %w", err)
	}

	log.Printf("🗄️ 資料庫版本: %s", version)
	return nil
}

// Close 關閉資料庫連接
func (db *DB) Close() {
	if db.DB != nil {
		db.DB.Close()
		log.Println("🔚 資料庫連接已關閉")
	}
}

// GetStats 獲取連接池統計資訊
func (db *DB) GetStats() sql.DBStats {
	return db.DB.Stats()
}

// SimpleConnection 簡單連接範例 (用於測試)
func SimpleConnection() error {
	databaseURL := os.Getenv("DATABASE_URL")
	if databaseURL == "" {
		return fmt.Errorf("DATABASE_URL 環境變數未設置")
	}

	db, err := sql.Open("postgres", databaseURL)
	if err != nil {
		return fmt.Errorf("連接資料庫失敗: %w", err)
	}
	defer db.Close()

	// 簡單查詢測試連接
	var version string
	if err := db.QueryRow("SELECT version()").Scan(&version); err != nil {
		return fmt.Errorf("查詢失敗: %w", err)
	}

	log.Println("🔗 簡單連接成功:", version)
	return nil
}