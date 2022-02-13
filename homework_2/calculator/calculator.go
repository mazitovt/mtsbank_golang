package calculator

import (
	"fmt"
	"strconv"
	"strings"
)

const arithmeticOperatorsByPriority = "/*-+"

type calculatedFunction = func(op1, op2 int64) (int64, error)

type arithmeticOperation struct {
	sign     rune
	function calculatedFunction
}

// арифметические операции в порядке приоритета выполнения
var arithmeticOperations = [...]arithmeticOperation{
	{'/', func(op1, op2 int64) (res int64, err error) {
		if op2 == 0 {
			err = &DivisionByZeroError{op1, op2}
		} else {
			res = op1 / op2
		}
		return
	}},
	{'*', func(op1, op2 int64) (res int64, err error) {
		res = op1 * op2
		return
	}},
	{'-', func(op1, op2 int64) (res int64, err error) {
		res = op1 - op2
		return
	}},
	{'+', func(op1, op2 int64) (res int64, err error) {
		res = op1 + op2
		return
	}},
}

func CalculateExpression(expression string) (result string, err error) {

	defer func() {
		if err != nil {
			err = &InvalidExpressionError{expression, err}
			return
		}
	}()

	operators, strOperands := splitToOperandsAndOperators(expression)

	if l1, l2 := len(strOperands), len(operators); l1 != l2+1 {
		err = &WrongNumberOfOperatorsAndOperands{l1, l1}
		return
	}

	operands, err := getInt64Operands(strOperands)
	if err != nil {
		return
	}

	expressionResult, err := getExpressionResult(operators, operands)
	if err != nil {
		return
	}

	operators = append(operators, '=')
	strOperands = append(strOperands, strconv.FormatInt(expressionResult, 10))

	result = concatExpressionWithWhitespaces(strOperands, operators)

	return
}

func splitToOperandsAndOperators(expression string) (operators []rune, operands []string) {
	splitFunc := func(c rune) bool {
		if strings.ContainsRune(arithmeticOperatorsByPriority, c) {
			operators = append(operators, c)
			return true
		}
		return false
	}
	operands = strings.FieldsFunc(expression, splitFunc)

	return
}

func getInt64Operands(stringOperands []string) (operands []int64, err error) {

	operands = make([]int64, len(stringOperands))

	for i, elem := range stringOperands {
		if number, e := strconv.ParseInt(elem, 10, 64); e != nil {
			err = &ConversionToIntegerError{elem}
			return
		} else {
			operands[i] = number
		}
	}

	return
}

func getExpressionResult(operatorsOrigin []rune, operandsOrigin []int64) (result int64, err error) {

	operators := make([]rune, len(operatorsOrigin))
	operands := make([]int64, len(operandsOrigin))
	copy(operands, operandsOrigin)
	copy(operators, operatorsOrigin)

	for _, operation := range arithmeticOperations {
		if err = calculateOperations(operation, &operators, &operands); err != nil {
			return
		}
	}

	result = operands[0]

	return
}

// Функция последоватльено подсчитывает все операции со знаком равным operation.sign
// Удаляет из operators все знаки, равные operation.sign
// Удаляет из operands каждые два операнда, учавствующие в operation,
//		и добавляет на место первого операнда результат вычисления operation.function
func calculateOperations(operation arithmeticOperation, operators *[]rune, operands *[]int64) (err error) {

	var operationResult int64

	for pos := 0; pos < len(*operators); pos++ {
		if (*operators)[pos] == operation.sign {
			operationResult, err = operation.function((*operands)[pos], (*operands)[pos+1])

			if err != nil {
				return
			}

			(*operands)[pos] = operationResult
			removeOperandAt(operands, pos+1)
			removeOperatorAt(operators, pos)
			pos--
		}
	}

	return
}

func removeOperandAt(slice *[]int64, s int) {
	*slice = append((*slice)[:s], (*slice)[s+1:]...)
}

func removeOperatorAt(slice *[]rune, s int) {
	*slice = append((*slice)[:s], (*slice)[s+1:]...)
}

func concatExpressionWithWhitespaces(operands []string, operators []rune) string {

	sb := strings.Builder{}
	sb.Grow(len(operands))
	sb.WriteString(operands[0])

	for i, operator := range operators {
		sb.WriteString(fmt.Sprintf(" %v %v", string(operator), operands[i+1]))
	}

	return sb.String()
}
