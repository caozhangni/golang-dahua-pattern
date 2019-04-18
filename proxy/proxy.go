package main

import "fmt"

type SchoolGirl struct {
	Name string
}

type IGiveGift interface {
	GiveDolls()
	GiveFlowers()
	GiveChocolate()
}

type Pursuit struct {
	mm *SchoolGirl
}

func (p *Pursuit) GiveDolls() {
	fmt.Println(p.mm.Name + " 送你洋娃娃")
}

func (p *Pursuit) GiveFlowers() {
	fmt.Println(p.mm.Name + " 送你鲜花")
}

func (p *Pursuit) GiveChocolate() {
	fmt.Println(p.mm.Name + " 送你巧克力")
}

type Proxy struct {
	gg *Pursuit
}

func (p *Proxy) GiveDolls() {
	p.gg.GiveDolls()
}

func (p *Proxy) GiveFlowers() {
	p.gg.GiveFlowers()
}

func (p *Proxy) GiveChocolate() {
	p.gg.GiveChocolate()
}

func NewProxy(mm *SchoolGirl) *Proxy {
	pursuit := &Pursuit{mm}
	return &Proxy{pursuit}
}

func main() {
	jiaojiao := &SchoolGirl{"李娇娇"}
	daili := NewProxy(jiaojiao)

	daili.GiveDolls()
	daili.GiveFlowers()
	daili.GiveChocolate()
}
