package main

import (
	"fmt"
	"log"
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
		log.Printf("allowed connection from %s\n", remoteAddr)
	} else {
		conn.Close()
		log.Printf("Blocked connection from %s\n", remoteAddr)
	}
}

// main func, it defines the port and the rules
// starts the listening on the port and calls the filter connection for each incoming connection

func main() {
	// log_file defines the path for the log file. the file is created if it does not exist
	log_file := "./errlog"
	logFile, err := os.OpenFile(log_file, os.O_APPEND|os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		log.Panic(err)
	}
	defer logFile.Close()
	log.SetOutput(logFile)
	log.SetFlags(log.Lshortfile | log.LstdFlags)
	port := "8080"
	rules := []string{"127.0.0.1"}

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
