package utils

import (
	"fmt"

	"github.com/zishang520/socket.io/v2/socket"
)

func GetSocketClient(clients ...any) *socket.Socket {
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
