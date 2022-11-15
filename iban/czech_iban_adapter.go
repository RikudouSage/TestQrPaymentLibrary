package iban

import (
	"fmt"
	"math/big"
	"strings"
)

type czechIbanAdapter struct {
	accountNumber string
	bankCode      string
	iban          *string
}

func (iban *czechIbanAdapter) String() string {
	if iban.iban == nil {
		part1 := getNumericRepresentation("C")
		part2 := getNumericRepresentation("Z")

		accountPrefix := "0"
		accountNumber := iban.accountNumber

		if strings.Contains(accountNumber, "-") {
			accountParts := strings.Split(accountNumber, "-")
			accountPrefix = accountParts[0]
			accountNumber = accountParts[1]
		}

		numeric := new(big.Int)
		numeric.SetString(fmt.Sprintf("%04s%06s%010s%s%s00", iban.bankCode, accountPrefix, accountNumber, part1, part2), 10)

		mod := numeric.Mod(numeric, big.NewInt(97)).Uint64()

		iban.iban = new(string)
		*iban.iban = fmt.Sprintf("%.2s%02d%04s%06s%010s", "CZ", 98-mod, iban.bankCode, accountPrefix, accountNumber)
	}

	return *iban.iban
}

func (iban *czechIbanAdapter) Validator() Validator {
	return NewCompoundValidator(NewCzechIbanValidator(), NewGenericIbanValidator())
}

func NewCzechIbanAdapter(accountNumber string, bankCode string) Iban {
	return &czechIbanAdapter{
		accountNumber: accountNumber,
		bankCode:      bankCode,
	}
}
