package gateways

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type ApiCepDTO struct {
	Code       string `json:"code"`
	State      string `json:"state"`
	City       string `json:"city"`
	District   string `json:"district"`
	Address    string `json:"address"`
	Status     int    `json:"status"`
	OK         bool   `json:"ok"`
	StatusText string `json:"statusText"`
}

func ApiCEP(ch chan interface{}, cep string) {
	if len(cep) == 8 {
		cep = fmt.Sprintf("%s-%s", cep[:5], cep[5:])
	}
	res, err := http.Get("https://cdn.apicep.com/file/apicep/" + cep + ".json")
	if err != nil {
		panic(err)
	}

	var data ApiCepDTO
	err = json.NewDecoder(res.Body).Decode(&data)
	if err != nil {
		panic(err)
	}

	ch <- data
}
