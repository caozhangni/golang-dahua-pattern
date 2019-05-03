package main

import "fmt"

type TestPaperI interface {
	Answer1() string
	Answer2() string
	Answer3() string
}

type TestPaper struct {
	Specific TestPaperI // 由于go没有继承的概念，所以通过Specific来绑定具体的试卷
}

func (tp *TestPaper) TestQuestion1() {
	if tp.Specific == nil {
		return
	}
	fmt.Println("题目1")
	fmt.Printf("答案:%v\n", tp.Specific.Answer1())
}

func (tp *TestPaper) TestQuestion2() {
	if tp.Specific == nil {
		return
	}
	fmt.Println("题目2")
	fmt.Printf("答案:%v\n", tp.Specific.Answer2())
}

func (tp *TestPaper) TestQuestion3() {
	if tp.Specific == nil {
		return
	}
	fmt.Println("题目3")
	fmt.Printf("答案:%v\n", tp.Specific.Answer3())
}

func (tp *TestPaper) Answer1() string {
	return "P 答案1"
}

func (tp *TestPaper) Answer2() string {
	return "P 答案2"
}

func (tp *TestPaper) Answer3() string {
	return "P 答案3"
}

type TestPaperA struct {
}

func (tp *TestPaperA) Answer1() string {
	return "A 答案1"
}

func (tp *TestPaperA) Answer2() string {
	return "A 答案2"
}

func (tp *TestPaperA) Answer3() string {
	return "A 答案3"
}

type TestPaperB struct {
}

func (tp *TestPaperB) Answer1() string {
	return "B 答案1"
}

func (tp *TestPaperB) Answer2() string {
	return "B 答案2"
}

func (tp *TestPaperB) Answer3() string {
	return "B 答案3"
}

func main() {
	fmt.Println("学生甲的试卷")
	pa := TestPaper{Specific: &TestPaperA{}}
	pa.TestQuestion1()
	pa.TestQuestion2()
	pa.TestQuestion3()

	fmt.Println("学生乙的试卷")
	pb := TestPaper{Specific: &TestPaper{}}
	pb.TestQuestion1()
	pb.TestQuestion2()
	pb.TestQuestion3()
}
