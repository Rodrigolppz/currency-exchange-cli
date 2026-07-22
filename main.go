package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Currency struct {
	Result          string             `json:"result"`
	BaseCode        string             `json:"base_code"`
	ConversionRates map[string]float64 `json:"conversion_rates"`
}

func Getrequest() []byte {

	resp, err := http.Get("https://v6.exchangerate-api.com/v6/<api-key>/latest/GBP")
	if err != nil {
		fmt.Println("Error")
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	fmt.Printf("Body : %s", body)

	return body

}

func DecodeJson(resp []byte) Currency {

	var c Currency

	err := json.Unmarshal(resp, &c)
	if err != nil {
		fmt.Println("Error")
	}

	return c

}

func main() {

	body := Getrequest()
	currency := DecodeJson(body)

	fmt.Println(currency)
	fmt.Printf("%T\n", currency)

}
