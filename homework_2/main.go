package main

import (
	"bufio"
	"fmt"
	"mtsbank_golang/homework_2/calculator"
	"os"
	"strings"
)

func main() {

	fmt.Println("'n' - выход из программы")
	for {
		fmt.Println("Введите выражение: ")
		reader := bufio.NewReader(os.Stdin)
		inputText, _ := reader.ReadString('\n')
		inputText = strings.TrimSuffix(inputText, "\r\n")

		if inputText == "n" {
			break
		}
		result, err := calculator.CalculateExpression(inputText)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(result)
		}
	}
}
