package main

import (
	"fmt"
	"time"

	client "github.com/silastgoes/multithreading-fullcycle/internal/httpclients"
)

func main() {
	cep := "45077000"
	viacep := make(chan string)
	brasilapi := make(chan string)

	go func() {
		brasilapi <- client.Getbrasilapi(cep)
	}()

	go func() {
		viacep <- client.Getviacep(cep)
	}()

	select {
	case address := <-viacep:
		fmt.Println("Endereço recuperado da https://brasilapi.com.br")
		fmt.Println(address)
	case address := <-brasilapi:
		fmt.Println("Endereço recuperado da https://viacep.com.br")
		fmt.Println(address)
	case <-time.After(time.Second):
		fmt.Println("TimeOut")
	}
}
