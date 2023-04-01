package main

import (
	"fmt"
	"net"
	"os"
)

// the func accepts 2 variables : rules and ip addresses(incoming network connections)
// the func below checks if a connection is from an ip address already in the rule book.
// if yes, the connection is allowed and vice versa

func filterconection(conn net.Conn, rules []string) {
	remoteAddr := conn.RemoteAddr().String()
	allowed := false
	for _, rule := range rules {
		if rule == remoteAddr {
			allowed = true
			break
		}
	}

	if allowed {
		fmt.Printf("allowed connection from %s\n", remoteAddr)
	} else {
		conn.Close()
		fmt.Printf("Blocked connection from %s\n", remoteAddr)
	}
}

// main func, it defines the port and the rules
// starts the listening on the port and calls the filter connection for each incoming connection

func main() {
	port := "8080"
	rules := []string{"192.168.0.1", "192.168.136.10", "192.168.136.10"}

	listener, err := net.Listen("tcp", ":"+port)
	if err != nil {
		fmt.Println("error starting server", err)
		os.Exit(1)
	}
	fmt.Printf("server listening on port %s\n", port)

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("error accepting connection", err)
			continue
		}
		go filterconection(conn, rules)
	}
}
