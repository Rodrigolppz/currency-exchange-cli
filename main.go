package main

import (
	"exchange-cli/client"
	"fmt"
	"time"
)

func main() {

	now := time.Now()

	body := client.Getrequest()
	currency := client.DecodeJson(body)
	find := client.FindCurrency(currency, "BRL")
	businessdays := client.BusinessDays()
	val := client.CalculateValue(find, businessdays)

	currentmonth := now.Month()

	fmt.Printf("\n Your salary in %v is: R$%v \n", currentmonth, val)
	fmt.Printf("And we have %v business days in %v \n", businessdays, currentmonth)

}
