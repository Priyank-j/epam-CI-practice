package main

import (
	"fmt"
	"log"
	"os"
	"time"
)

func main() {
	listenAddr := "5000"

	// create a logger, router and server
	logger := log.New(os.Stdout, "http: ", log.LstdFlags)
	router := newRouter()
	server := newServer(
		listenAddr,
		(middlewares{logging(logger), tracing(func() string { return fmt.Sprintf("%d", time.Now().UnixNano()) })}).apply(router),
		logger,
	)

	// run our server
	if err := server.run(); err != nil {
		log.Fatal(err)
	}
}
