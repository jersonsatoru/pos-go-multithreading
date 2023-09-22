package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"
)

func main() {
	cep := os.Args[1]

	ch1 := make(chan interface{})
	ch2 := make(chan interface{})

	go viaCEP(ch1, cep)

	go apiCEP(ch2, cep)

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

func viaCEP(ch chan interface{}, cep string) {
	res, err := http.Get("http://viacep.com.br/ws/" + cep + "/json/")
	if err != nil {
		panic(err)
	}

	var data interface{}
	json.NewDecoder(res.Body).Decode(&data)
	ch <- data
}

func apiCEP(ch chan interface{}, cep string) {
	res, err := http.Get("https://cdn.apicep.com/file/apicep/" + cep + ".json")
	if err != nil {
		panic(err)
	}

	var data interface{}
	json.NewDecoder(res.Body).Decode(&data)
	ch <- data
}
