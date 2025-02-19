package database

import (
	"context"
	"fmt"
	"log"
)

func (q *Queries) AddUser(ctx context.Context, username, hashedPassword string) error {
	// Define the query string
	queryString := "INSERT INTO users (username, hashed_password) VALUES (?, ?)"

	stmt, err := q.PrepareContext(ctx, queryString)
	if err != nil {
		return fmt.Errorf("failed to prepare statement: %w", err)
	}
	defer stmt.Close() // Ensure the statement is closed after use
	// Execute the statement with the provided parameters
	_, err = stmt.ExecContext(ctx, username, hashedPassword)
	if err != nil {
		return fmt.Errorf("failed to execute statement: %w", err)
	}

	log.Println("User added successfully")
	return nil
}
