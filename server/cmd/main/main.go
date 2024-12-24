package main

import "atypicaldev/splendor-go/pkg/server"

func main() {
	server.Run(server.ServerOpts{Addr: ":8080"})
}
