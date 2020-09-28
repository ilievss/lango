package main

import (
	"fmt"
	"github.com/ilievss/lango/web"
	"os"
	"strconv"
	"strings"
)

func main() {
	if len(os.Args) > 3 && !strings.Contains(os.Args[1], "-port") {
		fmt.Println("Program accepts only one argument: -port")
		os.Exit(0)
	}

	port, err := strconv.Atoi(os.Args[2])
	if err != nil {
		fmt.Println("Not a valid port: " + os.Args[2])
		os.Exit(0)
	}

	if port < 1000 || port > 65535 {
		fmt.Println("Port not in range [1000, 65535]: " + os.Args[2])
		os.Exit(0)
	}

	addr := fmt.Sprintf(":%d", port)
	server := web.New(addr)
	server.Start()
}
