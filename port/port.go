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

// ScanAllPorts resturns
func ScanAllPorts(initialPort, finalPort int) {
	for i := initialPort; i <= finalPort; i++ {
		if ret := ScanPort("tcp", "localhost", i); ret == fmt.Sprintf("The port %d is open", i) {
			fmt.Println(ret)
		}
	}
}

func AllPorts() {
	go ScanAllPorts(1, 10000)
	go ScanAllPorts(10001, 20000)
	go ScanAllPorts(20001, 30000)
	go ScanAllPorts(30001, 40000)
	go ScanAllPorts(40001, 50000)
	go ScanAllPorts(50001, 60000)
	ScanAllPorts(60001, 64435)

}
