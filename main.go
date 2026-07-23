package main

import (
	"fmt"

	"exchange-cli/client"
)

func main() {

	body := client.Getrequest()
	currency := client.DecodeJson(body)
	find := client.FindCurrency(currency, "BRL")

	fmt.Println(find)

}
