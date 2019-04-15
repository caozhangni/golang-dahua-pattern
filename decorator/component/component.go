package main

import "fmt"

type Component interface {
	Operation()
}

type ConcreteComponent struct {
}

func (cc *ConcreteComponent) Operation() {
	fmt.Println("具体对象的操作")
}

type Decorator struct {
	component Component
}

func (d *Decorator) Operation() {
	if d.component != nil {
		d.component.Operation()
	}
}

func (d *Decorator) SetComponent(component Component) {
	d.component = component
}

type ConcreteDecoratorA struct {
	Decorator
	addedState string
}

func (cda *ConcreteDecoratorA) Operation() {
	cda.Decorator.Operation()
	cda.addedState = "New State"
	fmt.Println("具体装饰对象A的操作")
}

type ConcreteDecoratorB struct {
	Decorator
}

func (cdb *ConcreteDecoratorB) AddedBehavior() {
}

func (cdb *ConcreteDecoratorB) Operation() {
	cdb.Decorator.Operation()
	cdb.AddedBehavior()
	fmt.Println("具体装饰对象B的操作")
}

func main() {
	c := new(ConcreteComponent)
	d1 := new(ConcreteDecoratorA)
	d2 := new(ConcreteDecoratorB)

	d1.SetComponent(c)
	d2.SetComponent(d1)
	d2.Operation()
}
