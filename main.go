package main

import (
	"fmt"
	"net"
	"os"
)

type Socket struct {
	Host string
	Port string
	Proto string
}

func (s *Socket) Connect() {
	fmt.Println("Connecting to " + s.Proto + "://" + s.Host + ":" + s.Port)
	conn, err := net.Dial(s.Proto, s.Host + ":" + s.Port)
	if err != nil {
		fmt.Println("Error connecting to " + s.Proto + "://" + s.Host + ":" + s.Port)
		fmt.Println(err)
		os.Exit(1)
	}
	defer conn.Close()
	fmt.Println("Connected to " + s.Proto + "://" + s.Host + ":" + s.Port)
}

func main() {
	proto := os.Args[1]
	host := os.Args[2]
	port := os.Args[3]
	s := Socket{Host: host, Port: port, Proto: proto}
	s.Connect()
}