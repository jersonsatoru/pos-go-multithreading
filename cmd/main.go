package main

import (
	"fmt"
	"os"
	"time"

	"github.com/jersonsatoru/pos-go-multithreading/internal/gateways"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Println("Passe o argumento de CEP: Ex `go run ./cmd/main.go 08710690`")
		return
	}
	cep := os.Args[1]

	ch1 := make(chan interface{})
	ch2 := make(chan interface{})

	// go gateways.ViaCEP(ch1, cep)

	go gateways.ApiCEP(ch2, cep)

	select {
	case data := <-ch1:
		fmt.Println("Via CEP")
		fmt.Printf("%+v", data)
	case data := <-ch2:
		fmt.Println("Api CEP")
		fmt.Printf("%+v", data)
	case <-time.After(time.Second * 1):
		panic("Timeout, passou 1 segundo")
	}
}
