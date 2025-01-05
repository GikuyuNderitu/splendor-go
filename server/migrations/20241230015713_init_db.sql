-- +goose Up
-- +goose StatementBegin
SELECT 'Up SQL query. Init db';

CREATE TYPE GemType AS ENUM ('green', 'blue', 'red', 'brown', 'white', 'gold');
CREATE TYPE EventType AS ENUM('start', 'take_coin', 'reserve', 'purchase_card', 'points_achieved', 'game_end');

CREATE TABLE users (
	user_id UUID DEFAULT gen_random_uuid(),
	name TEXT NOT NULL,
  email TEXT NOT NULL,
  password TEXT NOT NULL,
	created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
	updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
	-- Keys
	PRIMARY KEY (user_id)
);

CREATE TABLE tables (
	table_id UUID DEFAULT gen_random_uuid(),
	display_name VARCHAR(50) NOT NULL,
	created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
	updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
	-- Keys
	PRIMARY KEY (table_id)
);

CREATE TABLE user_tables (
	id BIGSERIAL,
	user_id UUID,
	table_id UUID,
	--- Keys
	PRIMARY KEY (user_id, table_id),
	CONSTRAINT fk_user FOREIGN KEY(user_id) REFERENCES users(user_id),
	CONSTRAINT fk_table FOREIGN KEY(table_id) REFERENCES tables(table_id)
);

CREATE TABLE games (
	game_id UUID DEFAULT gen_random_uuid(),
	-- The id used to initially generate this game 
	hash_id TEXT NOT NULL,
	table_id UUID,
	game jsonb,
	-- Keys
	PRIMARY KEY (game_id),
	CONSTRAINT fk_table FOREIGN KEY(table_id) REFERENCES tables(table_id)
);

CREATE TABLE game_events (
  event_id BIGSERIAL,
  game_id UUID,
  event_type EventType,
  game jsonb,
  -- Keys
  PRIMARY KEY (event_id),
  CONSTRAINT fk_game FOREIGN KEY(game_id) REFERENCES games(game_id)
);


CREATE TABLE user_hands (
	game_id UUID,
	user_id UUID,
	nobles json[],
	-- All a player's coins. Each instance maps to a value of "1"
	coins GemType[] DEFAULT ARRAY[]::GemType[],
	owned_cards json[] DEFAULT ARRAY[]::json[],
	reserved_cards json[] DEFAULT ARRAY[]::json[],
	-- Keys
	PRIMARY KEY (user_id, game_id),
	CONSTRAINT fk_user FOREIGN KEY(user_id) REFERENCES users(user_id),
	CONSTRAINT fk_game FOREIGN KEY(game_id) REFERENCES games(game_id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'Down SQL query. Remove DB.';
DROP TABLE IF EXISTS user_tables;
DROP TABLE IF EXISTS user_hands;
DROP TABLE IF EXISTS games;
DROP TABLE IF EXISTS users;
DROP TABLE IF EXISTS tables;
DROP TYPE IF EXISTS GemType;
DROP TYPE IF EXISTS EventType;
-- +goose StatementEnd
