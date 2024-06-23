package socket_service

import (
	socket_test_route "athena-pos-backend/sockets/test"
	"athena-pos-backend/utils"
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/zishang520/socket.io/v2/socket"
)

func SetupSocket(socket_port string) {

	fmt.Println("Setting Socket Routes.")

	// TODO: Add Options here
	var opts *socket.ServerOptions = socket.DefaultServerOptions()

	// rotues
	io := setupRoutes(opts)

	mux := http.NewServeMux()
	mux.Handle("/athena/", io.ServeHandler(opts))

	server := &http.Server{
		Addr:    ":" + socket_port,
		Handler: mux,
	}

	// Run the server in a goroutine
	go func() {
		fmt.Println("Socket Listening on port:", socket_port)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Error starting Socket.IO server: %v", err)
		}
	}()

	// Handle shutdown signals
	handleShutdown(server, io)

}

func handleShutdown(server *http.Server, io *socket.Server) {

	exit := make(chan os.Signal, 1)
	signal.Notify(exit, os.Interrupt, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	<-exit
	log.Println("Shutting down...")

	context, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Gracefully shut down the HTTP server
	if err := server.Shutdown(context); err != nil {
		log.Fatalf("HTTP Server Shutdown Failed: %v", err)
	}
	// Cleanly close socket server
	io.Close(func(err error) {
		if err != nil {
			log.Fatalf("Error Closing Socket.IO Server: " + err.Error())
		}
	})

	log.Println("Server shut down gracefully.")
	os.Exit(0)
}

func setupSocketDefaultRoute(io *socket.Server) {
	io.Of("/", func(clients ...any) {
		client := utils.GetSocketClient(clients...)
		if client == nil {
			return
		}
	})
}

func setupRoutes(opts *socket.ServerOptions) *socket.Server {

	io := socket.NewServer(nil, opts)

	// Root Path
	setupSocketDefaultRoute(io)

	// Router
	io.Of("/test", socket_test_route.TestSocketRoute)

	return io
}
