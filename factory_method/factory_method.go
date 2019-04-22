package main

import "fmt"

type OperationI interface {
	GetResult() (float64, error)
	SetNumberA(float64)
	SetNumberB(float64)
}

type Operation struct {
	NumberA float64
	NumberB float64
}

func (o *Operation) SetNumberA(a float64) {
	o.NumberA = a
}

func (o *Operation) SetNumberB(b float64) {
	o.NumberB = b
}

type OperationAdd struct {
	Operation
}

func (o *OperationAdd) GetResult() (float64, error) {
	return o.NumberA + o.NumberB, nil
}

type OperationSub struct {
	Operation
}

func (o *OperationSub) GetResult() (float64, error) {
	return o.NumberA - o.NumberB, nil
}

type OperationMul struct {
	Operation
}

func (o *OperationMul) GetResult() (float64, error) {
	return o.NumberA * o.NumberB, nil
}

type OperationDiv struct {
	Operation
}

func (o *OperationDiv) GetResult() (float64, error) {
	return o.NumberA / o.NumberB, nil
}

type IFactory interface {
	CreateOperation() Operation
}

type AddFactory struct {
}

func (af *AddFactory) CreateOperation() OperationI {
	return &OperationAdd{}
}

type SubFactory struct {
}

func (sf *SubFactory) CreateOperation() OperationI {
	return &OperationSub{}
}

type MulFactory struct {
}

func (mf *MulFactory) CreateOperation() OperationI {
	return &OperationMul{}
}

type DivFactory struct {
}

func (df *DivFactory) CreateOperation() OperationI {
	return &OperationDiv{}
}

func main() {
	operFactory := new(AddFactory)
	oper := operFactory.CreateOperation()
	oper.SetNumberA(1)
	oper.SetNumberB(2)
	result, _ := oper.GetResult()
	fmt.Print(result)
}
