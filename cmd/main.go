package main

import (
	"fmt"
	"local/github/demo-app/internal/handlers"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/braintree/manners"
)

func main() {
	fmt.Println("App started...")

	httpAddr := os.Getenv("DEMO_APP_ADDR")
	if httpAddr == "" {
		log.Fatal("DEMO_APP_ADDR must be set and non-empty")
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/", handlers.HelloHandler)

	httpServer := manners.NewServer()
	httpServer.Addr = httpAddr
	httpServer.Handler = handlers.LoggingHandler(mux)

	errChan := make(chan error, 10)

	go func() {
		errChan <- httpServer.ListenAndServe()
	}()

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)

	for {
		select {
		case err := <-errChan:
			if err != nil {
				log.Fatal(err)
			}
		case s := <-signalChan:
			log.Println(fmt.Sprintf("Captured %v. Exiting gracefully...", s))
			httpServer.BlockingClose()
			os.Exit(0)
		}
	}
}
