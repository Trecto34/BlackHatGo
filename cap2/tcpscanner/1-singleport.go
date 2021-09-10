package main

import (
	"fmt"
	"net"
)

func main() {
	_, err := net.Dial("tcp", "scanme.nmap.org:700")
	if err == nil {
		fmt.Println("Conection successful")
	} else {
		fmt.Println("Port Closed")
	}
}
