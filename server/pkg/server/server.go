package server

import (
	"fmt"
	"net/http"
)

type ServerOpts struct {
	Addr string
}

func Run(opts ServerOpts) {
	mux := http.NewServeMux()

	mux.Handle("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello world"))
	}))

	fmt.Printf("Starting server with address: %s\n", opts.Addr)
	if err := http.ListenAndServe(opts.Addr, mux); err != nil {
		fmt.Printf("Error with server, shutting down: %v", err)
	}
}
