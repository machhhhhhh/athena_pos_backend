package socket_test_route

import (
	"athena-pos-backend/utils"
	"fmt"
)

func TestSocketRoute(clients ...any) {
	client := utils.GetSocketClient(clients...)

	client.On("athena", func(messages ...any) {
		if len(messages) == 0 {
			return
		}
		fmt.Println(client.Id(), "sent", messages)

		// boardcast to reply room
		client.Broadcast().Emit("reply", messages...)
	})

}
