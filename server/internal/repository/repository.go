package repository

import (
	"atypicaldev/splendor-go/internal/data"
	"context"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
)

type SplendorRepository interface {
	CreateTable(ctx context.Context, displayName string) (*data.Table, error)
}

type splendorRepository struct {
	pool *pgxpool.Pool
}

func NewRepository(pool *pgxpool.Pool) *splendorRepository {
	if err := pool.Ping(context.Background()); err != nil {
		log.Panicf("Error pinging postgres while setting up table repo: %v", err)
	}

	return &splendorRepository{pool}
}

func (r *splendorRepository) CreateTable(ctx context.Context, displayName string) (*data.Table, error) {
	conn, err := r.pool.Acquire(ctx)
	if err != nil {
		log.Printf("Issue acquiring pool when creating table: %v", err)
		return nil, err
	}
	defer conn.Release()

	queries := data.New(conn)
	log.Printf("Creating new table with name: %s", displayName)
	table, err := queries.CreateTable(ctx, displayName)
	if err != nil {
		log.Printf("Issue creating new table: %v", err)
	}

	return &table, nil
}
