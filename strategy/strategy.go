package main

import (
	"fmt"
	"golang-dahua-pattern/strategy/cash"
)

func main() {
	csuper, err := cash.NewCashContext("打8折")
	if err != nil {
		fmt.Print(err)
	} else {
		result := csuper.GetResult(700)
		fmt.Println(result)
	}
}
