package server

import (
	"fmt"
	"net"
)

func StartUDPServer() error {

	//----------------------------------------------------
	// The net.ResolveUDPAddr function is used to resolve the string to a UDP address structure.
	// It takes two arguments: the network type ("udp") and the address string
	// The function returns a *net.UDPAddr struct or an error if the string is not a valid UDP address.
	// thee struct contains the IP address and port number that the server will listen on.
	// So address is a struct that holds the resolved UDP address, and err is used to capture any error that occurs during the resolution process.
	address, err := net.ResolveUDPAddr(
		"udp",
		"127.0.0.1:8053",
	)

	if err != nil {
		return err
	}
	//----------------------------------------------------

	// It tells the operating system to create a UDP socket and bind it to the specified address and port.
	// The returned conn object represents the UDP connection, which can be used to read from and write to the socket.
	// A UDP socket is a software endpoint that lets a program send and receive independent messages called datagrams over a network using the User Datagram Protocol (UDP).
	conn, err := net.ListenUDP("udp", address)

	if err != nil {
		return err
	}

	//----------------------------------------------------
	// The defer statement is used to ensure that the conn.Close() method is called when the StartUDPServer function returns,
	//  regardless of whether it returns normally or due to an error.
	//  This is important for resource management,
	//  as it ensures that the UDP connection is properly closed and any associated resources are released when the function exits.
	defer conn.Close()

	//----------------------------------------------------

	fmt.Println("UDP server listening on", address)

	// Create a buffer to hold the incoming data.
	// it is a byte slice with a length of 512 bytes. This buffer will be used to store the data received from clients.
	buffer := make([]byte, 512)

	//----------------------------------------------------
	// The for loop is an infinite loop that continuously listens for incoming UDP packets.
	for {
		// The conn.ReadFromUDP method waits for a UDP packet to arrive on the socket.
		// When a packet is received, it copies the data into the buffer
		// and returns the number of bytes read (n), the address of the client that sent the packet (clientAddress), and any error that occurred during the read operation.

		n, clientAddress, err := conn.ReadFromUDP(buffer)

		if err != nil {
			return err
		}

		// Print the number of bytes received and the client's address.
		fmt.Println(
			"Received",
			n,
			"bytes from",
			clientAddress,
		)
	}
}
