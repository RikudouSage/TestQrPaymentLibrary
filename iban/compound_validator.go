package iban

type compoundValidator struct {
	Validators []Validator
}

func (validator *compoundValidator) IsValid(iban Iban) bool {
	if validator.Validators == nil {
		panic("validators must be assigned")
	}
	for _, validator := range validator.Validators {
		if !validator.IsValid(iban) {
			return false
		}
	}
	return true
}

func NewCompoundValidator(validators ...Validator) Validator {
	return &compoundValidator{Validators: validators}
}
