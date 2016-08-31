package main

import (
	"fmt"
	"net"
	"os"
)

func cr(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func main() {
	socket, err := net.ListenUDP("udp4", &net.UDPAddr{
		IP:   net.IPv4(0, 0, 0, 0),
		Port: 10050,
	})
	cr(err)
	defer socket.Close()
	for {
		data := make([]byte, 1024)
		read, remoteAddr, err := socket.ReadFromUDP(data)
		cr(err)
		fmt.Println(read, remoteAddr)
		fmt.Println(string(data))
	}
}
