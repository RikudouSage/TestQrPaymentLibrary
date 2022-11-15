package iban

type genericIban struct {
	Iban string
}

func (iban *genericIban) String() string {
	return iban.Iban
}

func (iban *genericIban) Validator() Validator {
	return NewGenericIbanValidator()
}

func NewGenericIban(iban string) Iban {
	return &genericIban{Iban: iban}
}
