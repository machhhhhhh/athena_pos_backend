package socket_helper

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/zishang520/socket.io/v2/socket"
)

func SetupSocket(socket_port string) {

	fmt.Println("Setting Socket Routes.")

	// TODO: Add Options here
	var opts *socket.ServerOptions = socket.DefaultServerOptions()

	// rotues
	io := setupRoutes(opts)

	// Set up the HTTP server for Socket.IO
	// TODO: can create req,res and wrap to context here (need to read more doc if want to parse service_controller here !!!)
	go func() {
		mux := http.NewServeMux()
		mux.Handle("/athena/", io.ServeHandler(opts))

		fmt.Println("Socket Listening on port:", socket_port)
		if err := http.ListenAndServe(":"+socket_port, mux); err != nil {
			log.Fatalf("Error starting Socket.IO server: %v", err)
		}
	}()

	exit := make(chan struct{})
	change_signal := make(chan os.Signal)

	signal.Notify(change_signal, os.Interrupt, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	go func() {
		for socket := range change_signal {
			switch socket {
			case os.Interrupt, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT:
				close(exit)
				return
			}
		}
	}()

	fmt.Println("Successfully Connected Socket")

	// grateful shutdown
	handleShutdown(io)

}

func handleShutdown(io *socket.Server) {

	exit := make(chan os.Signal, 1)
	signal.Notify(exit, os.Interrupt, syscall.SIGTERM)

	<-exit // Wait for a SIGINT or SIGTERM signal
	log.Println("Shutting down...")

	// Cleanly close socket server
	io.Close(func(err error) {
		if err != nil {
			log.Fatalf("Error Closing Socket.IO Server: " + err.Error())
		}
	})
	os.Exit(0)
}

func getSocketClient(clients ...any) *socket.Socket {
	if len(clients) == 0 || clients == nil {
		return nil
	}
	client := clients[0].(*socket.Socket)

	fmt.Println(client.Id(), "had connected.")

	client.On("disconnect", func(...any) {
		fmt.Println(client.Id(), "had disconnected.")
	})

	return client
}

func setupSocketDefaultRoute(io *socket.Server) {
	io.Of("/", func(clients ...any) {
		client := getSocketClient(clients...)
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
	io.Of("/test", TestSocketRoute)

	return io
}
