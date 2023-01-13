package server

import (
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/gorilla/mux"
	"github.com/projects/loans/websocket/handlers"
	"github.com/valyala/fasthttp"
)

func WebSockerServer() {

	router := mux.NewRouter()
	router.
		HandleFunc("/", handlers.HomeWebSocketHandler).
		Methods(fasthttp.MethodGet)

	address := os.Getenv("localhost")
	port := os.Getenv("1000")
	log.Printf("Starting server on %s:%s ...", address, port)
	log.Fatal(http.ListenAndServe(":1000", router))

	// SIGINT is the signal sent when we press Ctrl+C
	// SIGTERM gracefully kills the process
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	<-quit

	log.Println("Shutting down server.....")

}
