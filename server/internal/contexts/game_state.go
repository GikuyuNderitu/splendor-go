package contexts

import (
	"atypicaldev/splendor-go/internal/data"
	"atypicaldev/splendor-go/internal/repository"
	"context"
	"errors"
	"log"
)

var (
	UnhandledGameEventType = errors.New("Unhandled GameEventType")
)

type GameEventType string

const (
	GameStart GameEventType = "start"
)

type GameEvent struct {
	Type GameEventType
}

type GameStateDispatcher interface {
	ReportEvent(ctx context.Context, event GameEvent) (*data.Game, error)
}

type dispatcher struct {
	repo repository.SplendorRepository
}

func New(repo repository.SplendorRepository) *dispatcher {
	return &dispatcher{repo}
}

func (d *dispatcher) ReportEvent(ctx context.Context, event GameEvent) (*data.Game, error) {
	if event.Type == GameStart {
		game, err := d.startGame(ctx)
		if err != nil {
			log.Printf("Error starting game: %v", err)
			return nil, err
		}

		return game, nil
	}

	return nil, UnhandledGameEventType
}

func (d *dispatcher) startGame(ctx context.Context) (*data.Game, error) {
	return nil, errors.New("StartGame Unimplemented")
}
