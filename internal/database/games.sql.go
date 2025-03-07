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
	ID           int
	UserID       int
	Country      string
	GameWon      bool
	WinCondition string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

func (q *Queries) AddGame(ctx context.Context, userID int, country string, gameWon bool, winCondition string) (int, error) {
	// Define the query string
	queryString := "INSERT INTO games (user_id, country, game_won, win_condition) VALUES (?, ?, ?, ?) RETURNING id"

	stmt, err := q.PrepareContext(ctx, queryString)
	if err != nil {
		return 0, fmt.Errorf("failed to prepare statement: %w", err)
	}
	defer stmt.Close() // Ensure the statement is closed after use
	// Execute the statement with the provided parameters
	game, err := stmt.ExecContext(ctx, userID, country, gameWon, winCondition)
	if err != nil {
		return 0, fmt.Errorf("failed to execute statement: %w", err)
	}
	gameID, err := game.LastInsertId()
	log.Printf("Gameid: %v", gameID)

	log.Println("Game added successfully")
	return int(gameID), nil
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
			&i.GameWon,
			&i.WinCondition,
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
