package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/Bilal-Z/go-ms-tutorial-nj/product-api/handlers"
)

func main()  {
	l := log.New(os.Stdout, "product-api", log.LstdFlags)

	ph := handlers.NewProduct(l)

	sm := http.NewServeMux()
	sm.Handle("/", ph)

	s := &http.Server{
		Addr: "127.0.0.1:5000",
		Handler: sm,
		IdleTimeout: 120 * time.Second,
		ReadTimeout: 1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}

	// create child routine with annonymous function/function literal
	go func(){
		l.Println("server starting on port 5000")
		err := s.ListenAndServe()
		if err != nil {
			l.Fatal(err)
		}
	}()

	// create channel
	sigChan := make(chan os.Signal, 2)
	// relay interupt and kill signals to channel
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)

	// await signal
	sig := <- sigChan
	l.Println("recieved terminate, graceful shutdown", sig)

	// timeout context
	tc, cancel := context.WithTimeout(context.Background(), 30*time.Second)

	defer cancel()

	s.Shutdown(tc)
}
