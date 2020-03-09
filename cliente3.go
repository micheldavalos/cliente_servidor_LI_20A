package main

import (
	"encoding/gob"
	"fmt"
	"net"
)

type Persona struct {
	Nombre string
	Email  []string
}

func cliente(p Persona) {
	c, err := net.Dial("tcp", "10.214.7.239:9999")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Enviando", p)
	err = gob.NewEncoder(c).Encode(p)
	if err != nil {
		fmt.Println(err)
	}
	c.Close()
}

func main() {
	p := Persona{
		Nombre: "Michel Dávalos",
		Email: []string{
			"michel.davalos@academicos.udg.mx",
			"michel.davalos@red.cucei.udg.mx",
		},
	}
	go cliente(p)

	var input string
	fmt.Scanln(&input)
}
