package repository

import (
	"atypicaldev/splendor-go/internal/data"
	"context"
	"errors"
	"log"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

type TableWithUsers struct {
	Users []data.User
	Table data.Table
}

type SplendorRepository interface {
	CreateTable(ctx context.Context, displayName string) (*data.Table, error)
	ListTables(ctx context.Context) ([]data.Table, error)
	JoinTable(ctx context.Context, tableId, userId string) (*TableWithUsers, error)
}

type splendorRepository struct {
	pool *pgxpool.Pool
}

var (
	ErrInvalidUserId                 = errors.New("Invalid user id")
	ErrInvalidTableId                = errors.New("Invalid table id")
	ErrJoiningTable                  = errors.New("Error joining table")
	ErrFetchingTableWithParticipants = errors.New("Error fetching table with participants")
)

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

func (r *splendorRepository) ListTables(ctx context.Context) ([]data.Table, error) {
	queries := data.New(r.pool)

	log.Println("Fetching tables")
	return queries.ListTables(ctx)
}

func (r *splendorRepository) JoinTable(ctx context.Context, tableId, userId string) (*TableWithUsers, error) {
	queries := data.New(r.pool)

	log.Println("Joining table")
	uId, err := uuid.Parse(userId)
	if err != nil {
		return nil, ErrInvalidUserId
	}

	tId, err := uuid.Parse(tableId)
	if err != nil {
		return nil, ErrInvalidTableId
	}

	err = queries.JoinTable(ctx, data.JoinTableParams{
		UserID:  uId,
		TableID: tId,
	})

	if err != nil {
		return nil, ErrJoiningTable
	}

	tableWithUsers := &TableWithUsers{}

	participantTable, err := queries.GetParticipants(ctx, tId)
	if err != nil || len(participantTable) <= 0 {
		return nil, ErrFetchingTableWithParticipants
	}

	for _, participant := range participantTable {
		tableWithUsers.Users = append(tableWithUsers.Users, participant.User)
	}

	tableWithUsers.Table = participantTable[0].Table
	return tableWithUsers, nil
}
