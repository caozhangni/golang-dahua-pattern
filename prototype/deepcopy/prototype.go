package main

import "fmt"

type Cloneable interface {
	Clone() Cloneable
}

type WorkExperience struct {
	WorkDate string
	Company  string
}

func (we *WorkExperience) SetWorkDate(WorkExperience string) {
	we.WorkDate = WorkExperience
}

func (we *WorkExperience) SetCompany(Company string) {
	we.Company = Company
}

func (we *WorkExperience) Clone() Cloneable {
	workExperience := *we
	return &workExperience
}

type Resume struct {
	name string
	sex  string
	age  string
	//timeArea string
	//company  string
	work *WorkExperience
}

func (r *Resume) SetPersonalInfo(sex string, age string) {
	r.sex = sex
	r.age = age
}

func (r *Resume) SetWorkExperience(workDate string, company string) {
	r.work.WorkDate = workDate
	r.work.Company = company
}

func (r *Resume) Display() {
	fmt.Printf("%v %v %v\n", r.name, r.sex, r.age)
	fmt.Printf("工作经历: %v %v\n", r.work.WorkDate, r.work.Company)
}

func (r *Resume) Clone() Cloneable {
	resume := *r
	resume.work = resume.work.Clone().(*WorkExperience)
	return &resume
}

func NewResume(name string) *Resume {
	return &Resume{name: name, work: &WorkExperience{}}
}

func main() {
	a := NewResume("大鸟")
	a.SetPersonalInfo("男", "29")
	a.SetWorkExperience("1998-2000", "XX公司")

	b := a.Clone().(*Resume)
	b.SetWorkExperience("1998-2006", "YY企业")

	c := b.Clone().(*Resume)
	c.SetPersonalInfo("男", "24")
	b.SetWorkExperience("1998-2006", "ZZ企业")

	a.Display()
	b.Display()
	c.Display()
}
