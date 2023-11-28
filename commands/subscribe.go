package commands

import (
	"net"
	"turtlequeue/store"
)

func SubscribeClient(commandArgs map[string]string, conn net.Conn) string {
	clientList, exists := store.SubscriptionMap.Load(commandArgs["topic"])
	if !exists {
		clientList = []net.Conn{conn}
	} else {
		clientList = append(clientList.([]net.Conn), conn)
	}
	store.SubscriptionMap.Store(commandArgs["topic"], clientList)
	return "+OK"
}
