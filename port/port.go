package port

import (
	"fmt"
	"net"
	"strconv"
	"time"
)

// ScanPort returns if the port is open or not
func ScanPort(protocol, hostname string, port int) string {
	address := hostname + ":" + strconv.Itoa(port)
	conn, err := net.DialTimeout(protocol, address, 60*time.Second)

	if err != nil {
		return fmt.Sprintf("The port %d is not open", port)
	}

	conn.Close()

	return fmt.Sprintf("The port %d is open", port)

}
