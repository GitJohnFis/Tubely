package database

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
	"time"
)

// User struct - expand fields as needed
type User struct {
	ID       string
	Password string
	Email    string
}

// CreateRefreshTokenParams struct - expand fields as needed
type CreateRefreshTokenParams struct {
	UserID    string
	Token     string
	ExpiresAt time.Time
}

type Client struct {
	db *sql.DB
}

func NewClient(pathToDB string) (Client, error) {
	db, err := sql.Open("sqlite3", pathToDB)
	if err != nil {
		return Client{}, err
	}
	c := Client{db}
	err = c.autoMigrate()
	if err != nil {
		return Client{}, err
	}
	return c, nil
}

func (c *Client) autoMigrate() error {
	userTable := `
	CREATE TABLE IF NOT EXISTS users (
		id TEXT PRIMARY KEY,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		password TEXT NOT NULL,
		email TEXT UNIQUE NOT NULL
	);
	`
	_, err := c.db.Exec(userTable)
	if err != nil {
		return err
	}
	refreshTokenTable := `
	CREATE TABLE IF NOT EXISTS refresh_tokens (
		token TEXT PRIMARY KEY,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		revoked_at TIMESTAMP,
		user_id TEXT NOT NULL,
		expires_at TIMESTAMP NOT NULL,
		FOREIGN KEY(user_id) REFERENCES users(id)
	);
	`
	_, err = c.db.Exec(refreshTokenTable)
	if err != nil {
		return err
	}

	videoTable := `
	CREATE TABLE IF NOT EXISTS videos (
		id TEXT PRIMARY KEY,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		title TEXT NOT NULL,
		description TEXT,
		thumbnail_url TEXT,
		video_url TEXT,
		user_id TEXT,
		FOREIGN KEY(user_id) REFERENCES users(id)
	);
	`
	_, err = c.db.Exec(videoTable)
	if err != nil {
		return err
	}
	return nil
}

func (c Client) Reset() error {
	if _, err := c.db.Exec("DELETE FROM refresh_tokens"); err != nil {
		return fmt.Errorf("failed to reset table refresh_tokens: %w", err)
	}
	if _, err := c.db.Exec("DELETE FROM users"); err != nil {
		return fmt.Errorf("failed to reset table users: %w", err)
	}
	if _, err := c.db.Exec("DELETE FROM videos"); err != nil {
		return fmt.Errorf("failed to reset table videos: %w", err)
	}
	return nil
}

// Stub: Implement GetUserByEmail as used in handler_login.go
func (c *Client) GetUserByEmail(email string) (User, error) {
	var user User
	row := c.db.QueryRow("SELECT id, password, email FROM users WHERE email = ?", email)
	err := row.Scan(&user.ID, &user.Password, &user.Email)
	if err != nil {
		return User{}, err
	}
	return user, nil
}

// Stub: Implement CreateRefreshToken as used in handler_login.go
func (c *Client) CreateRefreshToken(params CreateRefreshTokenParams) (string, error) {
	_, err := c.db.Exec(
		`INSERT INTO refresh_tokens (token, user_id, expires_at) VALUES (?, ?, ?)`,
		params.Token, params.UserID, params.ExpiresAt,
	)
	return params.Token, err
}
