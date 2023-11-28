package commands

import (
	"fmt"
	"net"
	"turtlequeue/store"
)

func PublishClient(commandArgs map[string]string, conn net.Conn) string {
	value, exists := store.SubscriptionMap.Load(commandArgs["topic"])
	if exists {
		clientInfo := value.([]net.Conn)
		for i := 0; i < len(clientInfo); i++ {
			clientInfo[i].Write([]byte(fmt.Sprintf("%s\r\n", commandArgs["payload"])))
		}
	}
	return "+OK"
}
