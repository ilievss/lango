package main

import (
	"fmt"
	"github.com/ilievss/lango/web"
)

func main() {
	port := 8080
	addr := fmt.Sprintf(":%d", port)
	server := web.New(addr)
	server.Start()
}
