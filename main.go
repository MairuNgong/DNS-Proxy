package main

import (
	"fmt"
	"log"
	"net"

	"github.com/miekg/dns"
)

func main() {
	addr := net.UDPAddr{
		Port: 53,
		IP:   net.ParseIP("0.0.0.0"),
	}

	conn, err := net.ListenUDP("udp", &addr)
	if err != nil {
		log.Fatalf("Failed to listen on port 53: %v", err)
	}
	defer conn.Close()

	fmt.Println("DNS Proxy listening on port 53...")

	buffer := make([]byte, 1024)

	for {
		n, clientAddr, err := conn.ReadFromUDP(buffer)
		if err != nil {
			log.Printf("Error reading UDP: %v", err)
			continue
		}

		var msg dns.Msg
		msg.Unpack(buffer[:n])

		fmt.Printf("From %s:\n", clientAddr)
		for _, q := range msg.Question {
			fmt.Printf(" - Question: %s, Type: %s\n", q.Name, dns.TypeToString[q.Qtype])
		}
		println("--------------------------------------")
	}
}
