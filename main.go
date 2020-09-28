package main

import (
	"flag"
	"fmt"
	"github.com/ilievss/lango/web"
	"os"
)

const defaultPort = 8080

func main() {
	portFlag := flag.Int("port", defaultPort, "The port on which to start the server")
	flag.Parse()

	port := *portFlag
	if port < 1000 || port > 65535 {
		fmt.Println("Port not in range [1000, 65535]: " + os.Args[2])
		os.Exit(0)
	}

	addr := fmt.Sprintf(":%d", port)
	server := web.New(addr)
	server.Start()
}
