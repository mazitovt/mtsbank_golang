package bank_deposit

import (
	"bufio"
	"errors"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type year int
type percent float64
type kopeck int

func OpenDeposit() (result string, err error) {

	reader := bufio.NewReader(os.Stdin)

	readInput := func(invitationLine string) string {
		fmt.Print(invitationLine)
		input, _ := reader.ReadString('\n')
		return strings.TrimSuffix(input, "\r\n")
	}

	deposit, e := parseDeposit(readInput("Введите размер депозита (в копейках): "))
	if e != nil {
		err = handleOpenDepositError(e)
		return
	}

	annualInterest, e := parseAnnualInterest(readInput("Введите размер годового процента (в формате #.#): "))
	if e != nil {
		err = handleOpenDepositError(e)
		return
	}

	timePeriod, e := parseTimePeriod(readInput("Введите срок вклада (в годах): "))
	if e != nil {
		err = handleOpenDepositError(e)
		return
	}

	finalDeposit := calculateDeposit(deposit, annualInterest, timePeriod)

	result = fmt.Sprintf("Размер вашего депозита через %v %v составит %v %v\n", timePeriod, yearRussianSpelling(timePeriod), finalDeposit, kopeckRussianSpelling(finalDeposit))

	return
}

func calculateDeposit(startDeposit kopeck, rate percent, timePeriod year) kopeck {

	curDeposit := startDeposit

	for curYear := year(0); curYear < timePeriod; curYear++ {
		curDeposit = kopeck(math.Round(float64(percent(curDeposit) * (1 + rate/100))))
	}

	return curDeposit
}

func parseDeposit(strDeposit string) (deposit kopeck, err error) {

	d, err := strconv.ParseInt(strDeposit, 10, 64)
	deposit = kopeck(d)
	if err != nil {
		err = fmt.Errorf("невозможно сконвертировать %v в целое число", strDeposit)
		return
	}
	if deposit < 0 {
		err = errors.New("депозит не может быть меньше нуля")
		return
	}

	return
}

func parseAnnualInterest(strAnnualInterest string) (annualInterest percent, err error) {

	ai, err := strconv.ParseFloat(strAnnualInterest, 64)
	annualInterest = percent(ai)
	if err != nil {
		err = fmt.Errorf("невозможно сконвертировать %v в дробное число", strAnnualInterest)
		return
	}
	if annualInterest < 0 {
		err = errors.New("годовой процент не может быть меньше нуля")
		return
	}

	return
}

func parseTimePeriod(strTimePeriod string) (timePeriod year, err error) {

	tp, err := strconv.ParseInt(strTimePeriod, 10, 64)
	timePeriod = year(tp)
	if err != nil {
		err = fmt.Errorf("невозможно сконвертировать %v в целое число", strTimePeriod)
		return
	}
	if timePeriod < 0 {
		err = errors.New("срок вклада не может быть меньше нуля")
		return
	}

	return
}

func yearRussianSpelling(year year) (spelling string) {
	switch year % 10 {
	case 1:
		spelling = "год"
	case 2, 3, 4:
		spelling = "года"
	default:
		spelling = "лет"
	}

	return
}

func kopeckRussianSpelling(kopeck kopeck) (spelling string) {
	switch kopeck % 10 {
	case 1:
		spelling = "копейка"
	case 2, 3, 4:
		spelling = "копейки"
	default:
		spelling = "копеек"
	}

	return
}

func handleOpenDepositError(e error) error {
	return fmt.Errorf("Открытие вклада невозможно: %v\n", e)
}
