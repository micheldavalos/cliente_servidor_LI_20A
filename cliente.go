package main

import (
	"fmt"
	"net"
)

func cliente() {
	c, err := net.Dial("tcp", ":9999")
	if err != nil {
		fmt.Println(err)
		return
	}
	msg := "Hola mundo"
	fmt.Println("Enviando", msg)
	c.Write([]byte(msg))
	c.Close()
}

func main() {
	go cliente()

	var input string
	fmt.Scanln(&input)
}
