package database

import (
	"context"
	"fmt"
	"log"
	"time"
)

/*
	User:
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    username TEXT NOT NULL UNIQUE,
    display_name TEXT NOT NULL,
    hashed_password TEXT NOT NULL,
    is_contributer BOOLEAN NOT NULL DEFAULT false,
    friendscode TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP

*/

type User struct {
	ID             int
	Username       string
	HashedPassword string
	CreatedAt      time.Time
	UpdatedAt      time.Time
	IsContributer  bool
}

func (q *Queries) AddUser(ctx context.Context, username, display_name, friendscode, hashedPassword string) error {
	// Define the query string
	queryString := "INSERT INTO users (username, display_name, friendscode, hashed_password) VALUES (?, ?)"

	stmt, err := q.PrepareContext(ctx, queryString)
	if err != nil {
		return fmt.Errorf("failed to prepare statement: %w", err)
	}
	defer stmt.Close() // Ensure the statement is closed after use
	// Execute the statement with the provided parameters
	_, err = stmt.ExecContext(ctx, username, display_name, friendscode, hashedPassword)
	if err != nil {
		return fmt.Errorf("failed to execute statement: %w", err)
	}

	log.Println("User added successfully")
	return nil
}

func (q *Queries) GetUserByUsername(ctx context.Context, username string) (User, error) {
	// Define the query string
	queryString := "SELECT * FROM users WHERE username=?"

	stmt, err := q.PrepareContext(ctx, queryString)
	if err != nil {
		return User{}, fmt.Errorf("failed to prepare statement: %w", err)
	}
	defer stmt.Close() // Ensure the statement is closed after use
	// Execute the statement with the provided parameters
	var user User
	err = stmt.QueryRowContext(ctx, username).Scan(&user.ID, &user.Username, &user.HashedPassword, &user.CreatedAt, &user.UpdatedAt, &user.IsContributer)
	if err != nil {
		return User{}, fmt.Errorf("failed to fetch User from Database: %w", err)
	}

	return user, nil
}

func (q *Queries) AddUsersGamesEntry(ctx context.Context, userID, gameID int) error {
	// Define the query string
	queryString := "INSERT INTO users_games (user_id, game_id) VALUES (?, ?)"

	stmt, err := q.PrepareContext(ctx, queryString)
	if err != nil {
		return fmt.Errorf("failed to prepare statement: %w", err)
	}
	defer stmt.Close() // Ensure the statement is closed after use
	// Execute the statement with the provided parameters
	_, err = stmt.ExecContext(ctx, userID, gameID)
	if err != nil {
		return fmt.Errorf("failed to execute statement: %w", err)
	}

	log.Println("Users_Games added successfully")
	return nil
}
