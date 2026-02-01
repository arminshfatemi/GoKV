package executor

type ResultType int

const (
	ResultError ResultType = iota
	ResultOK
	ResultInt
	ResultString
	ResultArray
)

type ExecutionResult struct {
	Type  ResultType
	Value any
}
