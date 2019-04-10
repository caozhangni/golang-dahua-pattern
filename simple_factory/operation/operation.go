package operation

type OperType int

const (
	Add OperType = 1 << iota
	Sub
	Mul
	Div
)

type OperationI interface {
	GetResult() (float32, error)
	SetNumberA(float32)
	SetNumberB(float32)
}

type Operation struct {
	NumberA float32
	NumberB float32
}

func (o *Operation) SetNumberA(a float32) {
	o.NumberA = a
}

func (o *Operation) SetNumberB(b float32) {
	o.NumberB = b
}

type OperationAdd struct {
	Operation
}

func (o *OperationAdd) GetResult() (float32, error) {
	return o.NumberA + o.NumberB, nil
}

type OperationSub struct {
	Operation
}

func (o *OperationSub) GetResult() (float32, error) {
	return o.NumberA - o.NumberB, nil
}

type OperationMul struct {
	Operation
}

func (o *OperationMul) GetResult() (float32, error) {
	return o.NumberA * o.NumberB, nil
}

type OperationDiv struct {
	Operation
}

func (o *OperationDiv) GetResult() (float32, error) {
	return o.NumberA / o.NumberB, nil
}

func NewOperFactory(t OperType) OperationI {
	switch t {
	case Add:
		return new(OperationAdd)
	case Sub:
		return new(OperationSub)
	case Mul:
		return new(OperationMul)
	case Div:
		return new(OperationDiv)
	default:
		return nil
	}
}
