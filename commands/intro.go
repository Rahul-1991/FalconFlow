package commands

import (
	"fmt"
	"net"
	"turtlequeue/utility"
)

func Intro(conn net.Conn) string {
	// Get the remote address of the client
	clientAddr := conn.RemoteAddr().String()
	data := map[string]interface{}{
		"host":        "0.0.0.0",
		"port":        4222,
		"client_ip":   clientAddr,
		"max_payload": 64,
	}
	return fmt.Sprintf("INFO %s", utility.JsonDump(data))
}
