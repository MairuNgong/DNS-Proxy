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
		go ResolveDNSRequest(conn, clientAddr, buffer[:n])
	}
}

func ResolveDNSRequest(conn *net.UDPConn, clientAddr *net.UDPAddr, query []byte) {
	upstreamAddr := net.UDPAddr{
		Port: 53,
		IP:   net.ParseIP("8.8.8.8"),
	}

	upConn, err := net.Dial("udp", upstreamAddr.String())
	if err != nil {
		log.Fatalf("Failed to connect to google dns: %v", err)
	}
	defer upConn.Close()

	_, err = upConn.Write(query)
	if err != nil {
		log.Println("Failed to forward DNS query:", err)
		return
	}

	buffer := make([]byte, 1024)

	n, err := upConn.Read(buffer)
	if err != nil {
		log.Println("Failed to read DNS answer:", err)
		return
	}

	var msg dns.Msg
	msg.Unpack(buffer[:n])

	for _, ans := range msg.Answer {
		switch rr := ans.(type) {
		case *dns.A:
			fmt.Printf(" - Resolved A Record: %s -> %s\n", rr.Hdr.Name, rr.A.String())
		case *dns.AAAA:
			fmt.Printf(" - Resolved AAAA Record: %s -> %s\n", rr.Hdr.Name, rr.AAAA.String())
		}
	}
}
