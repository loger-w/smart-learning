package repositories

import (
	"database/sql"
	"fmt"
	"smart-learning-backend/pkg/models"

	"github.com/lib/pq"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) CreateUser(user *models.User) error {
	query := `
		INSERT INTO users (email, username, password_hash, learning_level, avatar_url)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id, created_at, updated_at
	`
	
	err := r.db.QueryRow(
		query,
		user.Email,
		user.Username,
		user.PasswordHash,
		user.LearningLevel,
		user.AvatarURL,
	).Scan(&user.ID, &user.CreatedAt, &user.UpdatedAt)
	
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			if pqErr.Code == "23505" { // unique_violation
				return fmt.Errorf("user already exists")
			}
		}
		return fmt.Errorf("failed to create user: %w", err)
	}
	
	return nil
}

func (r *UserRepository) GetUserByEmail(email string) (*models.User, error) {
	user := &models.User{}
	query := `
		SELECT id, email, username, password_hash, learning_level, avatar_url, created_at, updated_at
		FROM users
		WHERE email = $1
	`
	
	err := r.db.QueryRow(query, email).Scan(
		&user.ID,
		&user.Email,
		&user.Username,
		&user.PasswordHash,
		&user.LearningLevel,
		&user.AvatarURL,
		&user.CreatedAt,
		&user.UpdatedAt,
	)
	
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("user not found")
		}
		return nil, fmt.Errorf("failed to get user: %w", err)
	}
	
	return user, nil
}

func (r *UserRepository) GetUserByID(id int) (*models.User, error) {
	user := &models.User{}
	query := `
		SELECT id, email, username, password_hash, learning_level, avatar_url, created_at, updated_at
		FROM users
		WHERE id = $1
	`
	
	err := r.db.QueryRow(query, id).Scan(
		&user.ID,
		&user.Email,
		&user.Username,
		&user.PasswordHash,
		&user.LearningLevel,
		&user.AvatarURL,
		&user.CreatedAt,
		&user.UpdatedAt,
	)
	
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("user not found")
		}
		return nil, fmt.Errorf("failed to get user: %w", err)
	}
	
	return user, nil
}

func (r *UserRepository) CheckUserExists(email, username string) (bool, error) {
	var count int
	query := `SELECT COUNT(*) FROM users WHERE email = $1 OR username = $2`
	
	err := r.db.QueryRow(query, email, username).Scan(&count)
	if err != nil {
		return false, fmt.Errorf("failed to check user existence: %w", err)
	}
	
	return count > 0, nil
}