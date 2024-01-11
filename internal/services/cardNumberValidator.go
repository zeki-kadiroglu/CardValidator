package main

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

func LuhnValidate(creditCardNumber string) bool {

	creditCardNumberWithoutDashes := func() string {
		if strings.Contains(creditCardNumber, "-") {
			return strings.ReplaceAll(creditCardNumber, "-", "")
		}
		return creditCardNumber

	}()

	var cardNumber string

	for _, r := range creditCardNumberWithoutDashes {
		if !unicode.IsSpace(r) {
			cardNumber += string(r)
		}
	}
	var sum int64 = 0
	parity := len(cardNumber) % 2

	cardNumWithoutChecksum := cardNumber[:len(cardNumber)-1]

	for i, v := range cardNumWithoutChecksum {

		item, err := strconv.Atoi(string(v))

		if err != nil {
			fmt.Println(err)
			return false
		}
		if int64(i)%2 != int64(parity) {
			sum += int64(item)
		} else if item > 4 {
			sum += int64(2*item - 9)
		} else {
			sum += int64(2 * item)
		}

	}

	checkDigit, err := strconv.Atoi(cardNumber[len(cardNumber)-1:])

	if err != nil {
		fmt.Println(err)
		return false
	}

	SumMod := sum % 10

	if SumMod == int64(0) {
		return SumMod == int64(checkDigit)
	}
	return int64(10)-SumMod == int64(checkDigit)
}

// func main() {
// 	isValid := LuhnValidate("4000-0566-5566-5556")
// 	fmt.Println(isValid)
// } // result true
