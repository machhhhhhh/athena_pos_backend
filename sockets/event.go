package socket_helper

import "fmt"

func TestSocketRoute(clients ...any) {
	client := getSocketClient(clients...)

	client.On("athena", func(messages ...any) {
		if len(messages) == 0 {
			return
		}
		fmt.Println(client.Id(), "sent", messages)

		// boardcast to reply room
		client.Broadcast().Emit("reply", messages...)
	})

}
