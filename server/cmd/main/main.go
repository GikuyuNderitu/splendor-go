package main

import "atypicaldev/splendor-go/pkg/setup"

func main() {
	setup.Run(setup.ServerOpts{Addr: ":8080"})
}
