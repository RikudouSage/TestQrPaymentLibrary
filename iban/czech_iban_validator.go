package iban

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type czechIbanValidator struct {
}

func (validator *czechIbanValidator) IsValid(iban Iban) bool {
	czechIban, ok := iban.(*czechIbanAdapter)

	if !ok {
		panic("only instances of *czechIbanAdapter can be validated using czechIbanValidator")
	}

	regex := regexp.MustCompile("^(?:([0-9]{0,6})-)?([0-9]{1,10})$")
	matches := regex.FindStringSubmatch(czechIban.accountNumber)

	if len(matches) == 0 {
		return false
	}

	var prefixNumber int
	if matches[1] != "" {
		var err error
		prefixNumber, err = strconv.Atoi(matches[1])

		if err != nil {
			panic("Account prefix is not a number")
		}
	}

	prefix := fmt.Sprintf("%06d", prefixNumber)
	account := strings.Repeat("0", 10-len(matches[2])) + matches[2]

	modifiers := [...]int{6, 3, 7, 9, 10, 5, 8, 4, 2, 1}
	prefixModifiers := modifiers[:len(modifiers)-4]

	modPrefix := 0
	for i := 0; i < 6; i++ {
		currentDigitAsInt, _ := strconv.Atoi(string(prefix[i]))
		modPrefix += currentDigitAsInt * prefixModifiers[i]
	}

	modPrefix %= 11
	if modPrefix != 0 {
		return false
	}

	modAccount := 0
	for i := 0; i < 10; i++ {
		currentDigitAsInt, err := strconv.Atoi(string(account[i]))
		if err != nil {
			panic("account number is not a number")
		}
		modAccount += currentDigitAsInt * modifiers[i]
	}
	modAccount %= 11

	if modAccount != 0 {
		return false
	}

	return true
}

func NewCzechIbanValidator() Validator {
	return &czechIbanValidator{}
}
