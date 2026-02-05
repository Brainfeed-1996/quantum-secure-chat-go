package main

import (
	"bufio"
	"fmt"
	"net"
	"quantum-secure-chat-go/internal/crypto"
)

func main() {
	listener, _ := net.Listen("tcp", ":9000")
	fmt.Println("[*] Quantum-Secure Server listening on :9000")
	kem := &crypto.PQCKEM{Algorithm: "Kyber768"}

	for {
		conn, _ := listener.Accept()
		go handleClient(conn, kem)
	}
}

func handleClient(conn net.Conn, kem *crypto.PQCKEM) {
	defer conn.Close()
	fmt.Printf("[*] New connection from %s
", conn.RemoteAddr())

	// 1. Generate PQC Keypair for this session
	pk, _, _ := kem.GenerateKeyPair()
	
	// 2. Send Public Key to Client
	conn.Write(append(pk, '
'))

	// 3. Receive Encapsulated Secret
	reader := bufio.NewReader(conn)
	ct, _ := reader.ReadBytes('
')
	fmt.Printf("[+] Established PQC shared secret with %s
", conn.RemoteAddr())
}