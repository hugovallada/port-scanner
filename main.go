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
	fmt.Println("Scanning ports...")

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("What port do you whant check?")
	portToCheck, err := reader.ReadString('\n')

	portInt, _ := strconv.Atoi(strings.TrimSpace(portToCheck))

	if err != nil {
		panic("Something went wrong")
	}

	result := port.ScanPort("tcp", "localhost", portInt)
	fmt.Println(result)
}
