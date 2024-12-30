package repository

import (
	"atypicaldev/splendor-go/internal/data"
	"context"
	"errors"
	"log"

	"connectrpc.com/connect"
	"github.com/jackc/pgx/v5/pgxpool"
)

var (
	repositoryKey = "splendor-repository-key"
)

var (
	ErrRepositoryNotInCtx = errors.New("Repository not found in context")
)

type SplendorRepository interface {
	CreateTable(ctx context.Context, displayName string) (*data.Table, error)
}

type splendorRepository struct {
	pool *pgxpool.Pool
}

func AddTableRepository(repo SplendorRepository) connect.UnaryInterceptorFunc {
	return func(next connect.UnaryFunc) connect.UnaryFunc {
		return connect.UnaryFunc(func(
			ctx context.Context,
			req connect.AnyRequest,
		) (connect.AnyResponse, error) {
			log.Printf("Adding splendor repo during route: %s", req.Spec().Procedure)

			repoCtx := context.WithValue(ctx, repositoryKey, repo)
			return next(repoCtx, req)
		})
	}
}

func RepositoryFromCtx(ctx context.Context) (SplendorRepository, error) {
	repo := ctx.Value(repositoryKey).(SplendorRepository)
	if repo == nil {
		return nil, ErrRepositoryNotInCtx
	}

	return repo, nil
}

func NewTableRepository(pool *pgxpool.Pool) *splendorRepository {
	if err := pool.Ping(context.Background()); err != nil {
		log.Panicf("Error pinging postgres while setting up table repo: %v", err)
	}

	return &splendorRepository{}
}

func (r *splendorRepository) CreateTable(ctx context.Context, displayName string) (*data.Table, error) {
	conn, err := r.pool.Acquire(ctx)
	if err != nil {
		log.Printf("Issue acquiring pool when creating table: %v", err)
		return nil, err
	}
	defer conn.Release()

	queries := data.New(conn)
	table, err := queries.CreateTable(ctx, displayName)
	if err != nil {
		log.Printf("Issue creating new table: %v", err)
	}

	return &table, nil
}
