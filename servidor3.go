package main

import (
	"encoding/gob" // golang object, json
	"fmt"
	"net"
)

type Persona struct {
	Nombre string
	Email  []string
}

func server() {
	s, err := net.Listen("tcp", ":9999")
	if err != nil {
		fmt.Println(err)
		return
	}
	for {
		c, err := s.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}
		go handleCliente(c)
	}
}

func handleCliente(c net.Conn) {
	var persona Persona
	err := gob.NewDecoder(c).Decode(&persona)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Mensaje:", persona)
	}
}

func main() {
	go server()

	var input string
	fmt.Scanln(&input)
}
