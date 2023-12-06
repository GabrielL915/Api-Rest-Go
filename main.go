package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/GabrielL915/Api-Rest-Go/db"
	"github.com/GabrielL915/Api-Rest-Go/handler"
)

func main() {
	addr := ":8080"
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatal("listener error: ", err.Error())
	}
	dbUser, dbPassword, dbName := os.Getenv("POSTGRES_USER"), os.Getenv("POSTGRES_PASSWORD"), os.Getenv("POSTGRES_DB")
	dbInstance, err := db.Initialize(dbUser, dbPassword, dbName)
	if err != nil {
		log.Fatal("database error: ", err.Error())
	}
	defer dbInstance.Conn.Close()

	httpHandler := handler.NewHandler(dbInstance)
	server := &http.Server{
		Handler: httpHandler,
	}

	go func() {
		server.Serve(listener)
	}()

	defer Stop(server)
	log.Printf("Server listening on %s", addr)
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	log.Println(fmt.Sprint(<-ch))
	log.Println("Shutting down server...")
}

func Stop(server *http.Server) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.Printf("Server shutdown failed: %v", err)
		os.Exit(1)
	}
}
