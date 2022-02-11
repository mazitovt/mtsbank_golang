package calculator

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

const arithmetics = "/*-+"

type arithmeticOperation = func(op1, op2 int64) (int64, error)

var arithmeticOperations = map[rune]arithmeticOperation{
	'/': func(op1, op2 int64) (res int64, err error) {
		if op2 == 0 {
			err = errors.New("деление на ноль")
		} else {
			res = op1 / op2
		}
		return
	},
	'*': func(op1, op2 int64) (res int64, err error) {
		res = op1 * op2
		return
	},
	'-': func(op1, op2 int64) (res int64, err error) {
		res = op1 - op2
		return
	},
	'+': func(op1, op2 int64) (res int64, err error) {
		res = op1 + op2
		return
	},
}

func handleExpressionError(expression string, e error) error {
	return fmt.Errorf("выражение '%v' некорректно:\n%v\n", expression, e)
}

func Calculate(expression string) (result string, err error) {

	var operations []rune

	splitFunc := func(c rune) bool {
		if strings.ContainsRune(arithmetics, c) {
			operations = append(operations, c)
			return true
		}
		return false
	}
	stringOperands := strings.FieldsFunc(expression, splitFunc)

	if len(stringOperands) != len(operations)+1 {
		err = handleExpressionError(expression, errors.New("количество операнд должно быть на один больше количества операторов"))
		return
	}

	operands, err := getInt64Operands(stringOperands)
	if err != nil {
		err = handleExpressionError(expression, err)
		return
	}

	expressionResult, err := calculateResult(operations[:], operands)
	if err != nil {
		err = handleExpressionError(expression, err)
		return
	}

	operations = append(operations, '=')

	stringOperands = append(stringOperands, strconv.FormatInt(expressionResult, 10))

	result = concatExpressionWithWhitespaces(stringOperands, operations)

	return
}

func getInt64Operands(stringOperands []string) (operands []int64, err error) {

	operands = make([]int64, len(stringOperands))

	for i, elem := range stringOperands {
		number, e := strconv.ParseInt(elem, 10, 64)
		if e != nil {
			err = fmt.Errorf("ошибка при попытке конвертировать строку '%v' в число: %v", elem, e)
			return
		}
		operands[i] = number
	}

	return
}

func concatExpressionWithWhitespaces(stringOperands []string, operations []rune) string {

	sb := strings.Builder{}

	sb.Grow(len(stringOperands))

	sb.WriteString(stringOperands[0])

	for i, operation := range operations {
		sb.WriteString(fmt.Sprintf(" %v %v", string(operation), stringOperands[i+1]))
	}

	return sb.String()
}

func calculateResult(operationsOrigin []rune, operandsOrigin []int64) (result int64, err error) {

	operations := make([]rune, len(operationsOrigin))
	operands := make([]int64, len(operandsOrigin))
	copy(operands, operandsOrigin)
	copy(operations, operationsOrigin)

	for _, sign := range arithmetics {
		for i := 0; i < len(operations); {
			operation := operations[i]
			if operation == sign {
				op1 := operands[i]
				op2 := operands[i+1]
				r, e := arithmeticOperations[operation](op1, op2)

				if e != nil {
					err = fmt.Errorf("ошибка при подсчете значения выражения (%v %v %v): %v", op1, string(operation), op2, e)
					return
				}

				operands[i] = r
				operands = removeInt64(operands, i+1)
				operations = removeRunes(operations, i)
			} else {
				i++
			}
		}
	}

	result = operands[0]

	return
}

func removeInt64(slice []int64, s int) []int64 {
	return append(slice[:s], slice[s+1:]...)
}

func removeRunes(slice []rune, s int) []rune {
	return append(slice[:s], slice[s+1:]...)
}
