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
    user_id INTEGER,
    country TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
	FOREIGN KEY(user_id) REFERENCES users(id)

*/

type Game struct {
	ID        int
	UserID    int
	Country   string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (q *Queries) AddGame(ctx context.Context, userID int, country string) error {
	// Define the query string
	queryString := "INSERT INTO games (user_id, country) VALUES (?, ?)"

	stmt, err := q.PrepareContext(ctx, queryString)
	if err != nil {
		return fmt.Errorf("failed to prepare statement: %w", err)
	}
	defer stmt.Close() // Ensure the statement is closed after use
	// Execute the statement with the provided parameters
	_, err = stmt.ExecContext(ctx, userID, country)
	if err != nil {
		return fmt.Errorf("failed to execute statement: %w", err)
	}

	log.Println("User added successfully")
	return nil
}

func (q *Queries) GetGamesByUserID(ctx context.Context, userID string) ([]Game, error) {
	// Define the query string
	queryString := "SELECT * FROM games WHERE user_id=?"

	stmt, err := q.PrepareContext(ctx, queryString)
	if err != nil {
		return nil, fmt.Errorf("failed to prepare statement: %w", err)
	}
	defer stmt.Close() // Ensure the statement is closed after use
	// Execute the statement with the provided parameters
	var games []Game
	rows, err := stmt.QueryContext(ctx, userID)
	if err != nil {
		return nil, fmt.Errorf("failed to load data from database: %w", err)
	}
	for rows.Next() {
		var i Game
		if err := rows.Scan(
			&i.ID,
			&i.UserID,
			&i.Country,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		games = append(games, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return games, nil
}
