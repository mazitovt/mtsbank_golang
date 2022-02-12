package main

import (
	"bufio"
	"fmt"
	"mtsbank_golang/homework_1/bank_deposit"
	"os"
	"strings"
)

func main() {

	for {
		if result, err := bank_deposit.OpenDeposit(); err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(result)
		}

		fmt.Print("Повторить ввод? (y/n): ")
		reader := bufio.NewReader(os.Stdin)
		input, _ := reader.ReadString('\n')
		input = strings.TrimSuffix(input, "\r\n")

		if input == "n" {
			break
		}
	}
}
