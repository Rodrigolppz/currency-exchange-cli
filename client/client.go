package client

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

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
