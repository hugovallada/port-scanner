package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/hugovallada/port-scanner/port"
)

func main() {
	fmt.Println("Welcome to the port scanner v0.0.1")

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Do you want to check an specific port, or check all ports ? (spe/all)")
	ans, err := reader.ReadString('\n')

	if err != nil {
		panic("Something went wrong")
	}

	if strings.TrimSpace(ans) == "all" {
		fmt.Println("Scanning ports 1 to 64435")
		port.ScanAllPorts(1, 64435)
	} else {

		fmt.Println("What port do you whant check?")
		portToCheck, err := reader.ReadString('\n')

		portInt, _ := strconv.Atoi(strings.TrimSpace(portToCheck))

		if err != nil {
			panic("Something went wrong")
		}

		result := port.ScanPort("tcp", "localhost", portInt)
		fmt.Println(result)
	}
}
