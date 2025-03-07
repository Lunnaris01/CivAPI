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

	Games:
	id INTEGER PRIMARY KEY AUTOINCREMENT,
    user_id INTEGER,
    country TEXT NOT NULL,
    game_won BOOLEAN NOT NULL,
    win_condition TEXT NOT NULL,
    sharekey TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
	FOREIGN KEY(user_id) REFERENCES users(id)

*/

type Game struct {
	Country      string
	GameWon      bool
	WinCondition string
	Sharekey     string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

func (q *Queries) AddGame(ctx context.Context, userID int, country string, gameWon bool, winCondition, shareKey string) (int, error) {
	// Define the query string
	queryString := "INSERT INTO games (user_id, country, game_won, win_condition, sharekey) VALUES (?, ?, ?, ?, ?) RETURNING id"

	stmt, err := q.PrepareContext(ctx, queryString)
	if err != nil {
		return 0, fmt.Errorf("failed to prepare statement: %w", err)
	}
	defer stmt.Close() // Ensure the statement is closed after use
	// Execute the statement with the provided parameters
	game, err := stmt.ExecContext(ctx, userID, country, gameWon, winCondition, shareKey)
	if err != nil {
		return 0, fmt.Errorf("failed to execute statement: %w", err)
	}
	gameID, err := game.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("failed to execute statement: %w", err)
	}

	log.Printf("Gameid: %v", gameID)

	log.Println("Game added successfully")
	return int(gameID), nil
}

func (q *Queries) GetGamesByUserID(ctx context.Context, userID string) ([]Game, error) {
	// Define the query string
	queryString := "SELECT g.country, g.game_won, g.win_condition, g.sharekey, g.created_at FROM games g INNER JOIN users_games ug ON g.id = ug.game_id WHERE ug.user_id=?"

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
			&i.Country,
			&i.GameWon,
			&i.WinCondition,
			&i.Sharekey,
			&i.CreatedAt,
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

func (q *Queries) GetGameIdByGameCode(ctx context.Context, gameCode string) (int, error) {
	queryString := "SELECT id FROM games WHERE sharekey=?"

	stmt, err := q.PrepareContext(ctx, queryString)
	if err != nil {
		return 0, fmt.Errorf("failed to prepare statement: %w", err)
	}
	defer stmt.Close() // Ensure the statement is closed after use
	// Execute the statement with the provided parameters
	row := stmt.QueryRowContext(ctx, gameCode)
	var gameID int
	err = row.Scan(&gameID)
	if err != nil {
		return 0, fmt.Errorf("failed to fetch game ID: %w", err)
	}
	return gameID, nil

}
