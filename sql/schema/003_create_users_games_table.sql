-- +goose Up
CREATE TABLE users_games (
    user_id INTEGER,
    game_id INTEGER,
	FOREIGN KEY(user_id) REFERENCES users(id),
    FOREIGN KEY(game_id) REFERENCES games(id),
    PRIMARY KEY(user_id,game_id)
);

-- +goose Down
DROP TABLE users_games;
