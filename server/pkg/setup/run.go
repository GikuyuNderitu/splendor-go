package setup

import (
	"atypicaldev/splendor-go/internal/repository"
	"atypicaldev/splendor-go/internal/server"
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"

	spv1connect "buf.build/gen/go/atypicaldev/splendorapis/connectrpc/go/atypicaldev/splendorapis/v1/splendorapisv1connect"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

type ServerOpts struct {
	Addr string
}

func Config() *pgxpool.Config {
	const defaultMaxConns = int32(4)
	const defaultMinConns = int32(0)
	const defaultMaxConnLifetime = time.Hour
	const defaultMaxConnIdleTime = time.Minute * 30
	const defaultHealthCheckPeriod = time.Minute
	const defaultConnectTimeout = time.Second * 5

	// Your own Database URL
	var DATABASE_URL string = os.Getenv("POSTGRES_URI")

	dbConfig, err := pgxpool.ParseConfig(DATABASE_URL)
	if err != nil {
		log.Fatal("Failed to create a config, error: ", err)
	}

	dbConfig.MaxConns = defaultMaxConns
	dbConfig.MinConns = defaultMinConns
	dbConfig.MaxConnLifetime = defaultMaxConnLifetime
	dbConfig.MaxConnIdleTime = defaultMaxConnIdleTime
	dbConfig.HealthCheckPeriod = defaultHealthCheckPeriod
	dbConfig.ConnConfig.ConnectTimeout = defaultConnectTimeout

	dbConfig.BeforeAcquire = func(ctx context.Context, c *pgx.Conn) bool {
		log.Println("Before acquiring the connection pool to the database!!")
		return true
	}

	dbConfig.AfterRelease = func(c *pgx.Conn) bool {
		log.Println("After releasing the connection pool to the database!!")
		return true
	}

	dbConfig.BeforeClose = func(c *pgx.Conn) {
		log.Println("Closed the connection pool to the database!!")
	}

	return dbConfig
}

func Run(opts ServerOpts) {
	pool, err := pgxpool.NewWithConfig(context.Background(), Config())
	if err != nil {
		panic(err)
	}

	defer pool.Close()

	if err = pool.Ping(context.Background()); err != nil {
		log.Printf("Error pinging db")
	}
	var greeting string

	if err != nil {
		panic("error reading simple select")
	}

	log.Println(greeting)

	repo := repository.NewRepository(pool)
	svc := &server.SplendorService{
		Repo: repo,
	}
	route, handler := spv1connect.NewSplendorServiceHandler(svc)

	mux := http.NewServeMux()

	mux.Handle(route, handler)

	fmt.Printf("Starting server with address: %s\n", opts.Addr)
	if err := http.ListenAndServe(opts.Addr, h2c.NewHandler(mux, &http2.Server{})); err != nil {
		fmt.Printf("Error with server, shutting down: %v", err)
	}

}
