package iban

import (
	"fmt"
	"math/big"
	"strconv"
)

type genericIbanValidator struct {
}

func getNumericRepresentation(str string) string {
	result := ""
	length := len(str)

	for i := 0; i < length; i++ {
		char := str[i]
		_, errorInt := strconv.Atoi(string(char))

		if errorInt != nil {
			result += strconv.Itoa(int(char) - int('A') + 10)
		} else {
			result += string(char)
		}
	}

	return result
}

func (validator *genericIbanValidator) IsValid(iban Iban) bool {
	if iban == nil {
		panic("nil IBAN is unsupported")
	}
	stringIban := fmt.Sprint(iban)

	country := stringIban[0:2]
	checksum := stringIban[2:4]
	account := stringIban[4:]

	numericCountry := getNumericRepresentation(country)
	numericAccount := getNumericRepresentation(account)

	inverted := new(big.Int)
	inverted.SetString(numericAccount+numericCountry+checksum, 10)

	return inverted.Mod(inverted, big.NewInt(97)).Cmp(big.NewInt(1)) == 0
}

func NewGenericIbanValidator() Validator {
	return &genericIbanValidator{}
}
