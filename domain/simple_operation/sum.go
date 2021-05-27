package simple_operation

type SumOperation struct{}

func (SumOperation) Execute(first int64, second int64) int64 {
	return first + second;
}
