package main

import (
	"fmt"
	"patterns/simple_factory/operation"
)

func main() {
	oper := operation.NewOperFactory(operation.Div)
	oper.SetNumberA(1)
	oper.SetNumberB(2)
	if r, err := oper.GetResult(); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(r)
	}
}
