package setup

import (
	"atypicaldev/splendor-go/internal/server"
	"fmt"
	"net/http"

	spv1connect "buf.build/gen/go/atypicaldev/splendorapis/connectrpc/go/atypicaldev/splendorapis/v1/splendorapisv1connect"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

type ServerOpts struct {
	Addr string
}

func Run(opts ServerOpts) {
	svc := &server.SplendorService{}
	route, handler := spv1connect.NewSplendorServiceHandler(svc)

	mux := http.NewServeMux()

	mux.Handle(route, handler)

	fmt.Printf("Starting server with address: %s\n", opts.Addr)
	if err := http.ListenAndServe(opts.Addr, h2c.NewHandler(mux, &http2.Server{})); err != nil {
		fmt.Printf("Error with server, shutting down: %v", err)
	}

}
