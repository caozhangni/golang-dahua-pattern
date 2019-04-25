package main

import "fmt"

type Cloneable interface {
	Clone() Cloneable
}

type Resume struct {
	name     string
	sex      string
	age      string
	timeArea string
	company  string
}

func (r *Resume) SetPersonalInfo(sex string, age string) {
	r.sex = sex
	r.age = age
}

func (r *Resume) SetWorkExperience(timeArea string, company string) {
	r.timeArea = timeArea
	r.company = company
}

func (r *Resume) Display() {
	fmt.Printf("%v %v %v\n", r.name, r.sex, r.age)
	fmt.Printf("工作经历: %v %v\n", r.timeArea, r.company)
}

func (r *Resume) Clone() Cloneable {
	resume := *r
	return &resume
}

func main() {
	a := &Resume{name: "大鸟"}
	a.SetPersonalInfo("男", "29")
	a.SetWorkExperience("1998-2000", "XX公司")

	b := a.Clone().(*Resume)
	b.SetWorkExperience("1998-2006", "YY企业")

	c := b.Clone().(*Resume)
	c.SetPersonalInfo("男", "24")

	a.Display()
	b.Display()
	c.Display()
}
