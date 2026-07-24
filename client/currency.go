package client

import (
	"encoding/json"
	"fmt"
)

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

	var found float64

	_, exists := currency.ConversionRates[searchKey] // this is a "comma ok idiom" or "map lookup", its a native golang expression to search for values in a map.
	// return the value associated with the key and a bool if the key exist or not.

	if exists {
		found = currency.ConversionRates[searchKey]

	}

	return found

}
