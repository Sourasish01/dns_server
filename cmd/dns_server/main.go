package main

import (
	"fmt"

	"github.com/bhotto/dns_server/internal/server"
)

func main() {
	fmt.Println("Starting DNS server...")

	err := server.StartUDPServer()

	if err != nil {
		fmt.Println("Server error:", err)
	}
}
