package main

import "fmt"

type PersonI interface {
	Show()
}

type Person struct {
	name string
}

func (p *Person) Show() {
	fmt.Printf("装扮的%v", p.name)
}

type Finery struct {
	Person
	component PersonI
}

func (f *Finery) Decorate(component PersonI) {
	f.component = component
}

func (f *Finery) Show() {
	if f.component != nil {
		f.component.Show()
	}
}

type TShirts struct {
	Finery
}

func (ts *TShirts) Show() {
	fmt.Print("大T恤 ")
	ts.Finery.Show()
}

type BigTrouser struct {
	Finery
}

func (bt *BigTrouser) Show() {
	fmt.Print("垮裤 ")
	bt.Finery.Show()
}

func main() {
	xc := &Person{"小菜"}
	fmt.Println("第一种装扮:")
	kk := new(BigTrouser)
	dtx := new(TShirts)
	kk.Decorate(xc)
	dtx.Decorate(kk)
	dtx.Show()
}
