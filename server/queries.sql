-- name: CreateUser :one
INSERT INTO users (
	name
) VALUES (
	$1
)
RETURNING *;

-- name: GetUser :one
SELECT * FROM users
WHERE user_id = $1 LIMIT 1;

-- name: ListTables :many
SELECT * FROM tables;

-- name: GetParticipants :many
SELECT
	u.user_id,
	u.name
FROM users AS u
JOIN user_tables AS ut ON u.user_id = ut.user_id
JOIN tables AS t ON ut.table_id = t.table_id
WHERE t.table_id = $1;

-- name: CreateGame :one
INSERT INTO games (
	hash_id, table_id, game
) VALUES (
	$1, $2, $3
)
RETURNING *;

-- name: AddPlayer :one
INSERT INTO user_hands (
	game_id, user_id
) VALUES (
	$1, $2
)
RETURNING *;

-- name: GetGame :one
SELECT * FROM games
WHERE table_id = $1 LIMIT 1;

-- name: GetPlayers :many
SELECT
	u.user_id,
	u.name
FROM games AS g
JOIN user_hands AS uh ON g.game_id = uh.game_id
JOIN users AS u ON uh.user_id = u.user_id
WHERE g.table_id = $1;

-- name: UpdateGame :exec
UPDATE games
SET game = $2
WHERE game_id = $1;

-- name: UpdateUserNoble :exec
UPDATE user_hands
SET nobles = $3
WHERE user_id = $1 AND game_id = $2;

-- name: UpdateUserReserved :exec
UPDATE user_hands
SET reserved_cards = $3
WHERE user_id = $1 AND game_id = $2;

-- name: UpdateUserOwnedCards :exec
UPDATE user_hands
SET owned_cards = $3
WHERE user_id = $1 AND game_id = $2;

-- name: UpdateUserCoins :exec
UPDATE user_hands
SET coins = $3
WHERE user_id = $1 AND game_id = $2;

