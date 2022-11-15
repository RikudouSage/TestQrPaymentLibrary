package QrPaymentCZ

import (
	"errors"
	"fmt"
	"github.com/RikudouSage/TestQrPaymentLibrary/iban"
	"github.com/shopspring/decimal"
	"time"
)

type QrPayment struct {
	VariableSymbol string
	SpecificSymbol string
	ConstantSymbol string
	Currency       string
	Comment        string
	Repeat         uint
	InternalId     string
	DueDate        time.Time
	Amount         decimal.Decimal
	PayeeName      string
	InstantPayment bool
	Iban           iban.Iban
}

func (payment *QrPayment) setDefaultValues() {
	if payment.Currency == "" {
		payment.Currency = "CZK"
	}
	if payment.Repeat == 0 {
		payment.Repeat = 7
	}
}

func (payment *QrPayment) GetQrString() (string, error) {
	payment.setDefaultValues()

	ibanStruct := payment.Iban
	if ibanStruct == nil {
		return "", errors.New("IBAN cannot be nil")
	}

	validator := ibanStruct.Validator()
	if validator != nil && !validator.IsValid(ibanStruct) {
		return "", errors.New("the IBAN is not valid")
	}

	qrString := "SPD*1.0*"
	qrString += fmt.Sprintf("ACC:%s*", ibanStruct.String())
	qrString += fmt.Sprintf("AM:%s*", payment.Amount.StringFixedBank(2))
	qrString += fmt.Sprintf("CC:%s*", payment.Currency)
	qrString += fmt.Sprintf("X-PER:%d*", payment.Repeat)

	if payment.Comment != "" {
		qrString += fmt.Sprintf("MSG:%.60s*", payment.Comment)
	}
	if payment.InternalId != "" {
		qrString += fmt.Sprintf("X-ID:%s*", payment.InternalId)
	}
	if payment.VariableSymbol != "" {
		qrString += fmt.Sprintf("X-VS:%s*", payment.VariableSymbol)
	}
	if payment.SpecificSymbol != "" {
		qrString += fmt.Sprintf("X-SS:%s*", payment.SpecificSymbol)
	}
	if payment.ConstantSymbol != "" {
		qrString += fmt.Sprintf("X-KS:%s*", payment.ConstantSymbol)
	}
	if payment.PayeeName != "" {
		qrString += fmt.Sprintf("RN:%s*", payment.PayeeName)
	}
	if !payment.DueDate.IsZero() {
		qrString += fmt.Sprintf("DT:%04d%02d%02d*", payment.DueDate.Year(), payment.DueDate.Month(), payment.DueDate.Day())
	}
	if payment.InstantPayment {
		qrString += "PT:IP*"
	}

	return qrString[:len(qrString)-1], nil
}

func (payment *QrPayment) String() string {
	result, err := payment.GetQrString()
	if err != nil {
		panic(err)
	}
	return result
}
