package calculator

import "fmt"

type DivisionByZeroError struct {
	operand1, operand2 int64
}

func (e *DivisionByZeroError) Error() string {
	return fmt.Sprintf("выражение '%d / %d': деление на ноль", e.operand1, e.operand2)
}

type WrongNumberOfOperatorsAndOperands struct {
	numberOfOperators, numberOfOperands int
}

func (e *WrongNumberOfOperatorsAndOperands) Error() string {
	return fmt.Sprintf("количество операнд (%d) должно быть на один больше количества операторов (%d)", e.numberOfOperators, e.numberOfOperators)
}

type ConversionToIntegerError struct {
	origin string
}

func (e *ConversionToIntegerError) Error() string {
	return fmt.Sprintf("строка '%v': ошибка при попытке конвертировать в целое число", e.origin)
}

type InvalidExpressionError struct {
	expression string
	err        error
}

func (e *InvalidExpressionError) Error() string {
	return fmt.Sprintf("в выражении '%v' содержится ошибка:\n%v\n", e.expression, e.err)
}
