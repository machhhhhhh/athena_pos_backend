package socket_helper

import "fmt"

func TestSocketRoute(clients ...any) {

	client := getSocketClient(clients...)
	if client == nil {
		return
	}

	client.On("athena", func(messages ...any) {
		if len(messages) == 0 {
			return
		}

		fmt.Println(client.Id(), "sent", messages)

		// Join the client to the "athena" room
		client.Emit("todo", messages)
	})

}
