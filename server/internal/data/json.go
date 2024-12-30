package data

import (
	"database/sql"

	"github.com/jackc/pgx/v5/pgtype"
)

type Requirement struct {
	Value pgtype.Int2 `json:"value"`
	Gem   Gemtype     `json:"gem"`
}

type Noble struct {
	Requirements []Requirement `json:"requirements"`
}

type Card struct {
	Requirements []Requirement  `json:"requirements"`
	Hidden       bool           `json:"hidden"`
	Url          sql.NullString `json:"url"`
}

type GameData struct {
	LowDeck    []Card  `json:"low_deck"`
	LowVisible [4]Card `json:"low_visible"`
	MidDeck    []Card  `json:"mid_deck"`
	MidVisible [4]Card `json:"mid_visible"`
	HighDeck   []Card  `json:"high_deck"`
	Highisible [4]Card `json:"high_visible"`
}
