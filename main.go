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

	err := json.Unmarshal(resp, &c) // Unmarshal function from encoding/json library, used to Decode raw Json and return the value to c
	if err != nil {
		fmt.Println("Error")
	}

	return c

}

// Just check if the key exists in the map, does not need to iterate over it.(actually a map is often used to replace for loops when you know the key)

func FindCurrency(currency Currency, searchKey string) float64 {

	var save float64

	_, exists := currency.ConversionRates[searchKey] // this is a "comma ok idiom" or "map lookup", its a native golang expression to search for values in a map.
	// return the value associated with the key and a bool if the key exist or not.

	if exists {
		save = currency.ConversionRates[searchKey]

	}

	return save

}

func main() {

	body := Getrequest()
	currency := DecodeJson(body)
	find := FindCurrency(currency, "BRL")

	fmt.Println(find)

}
