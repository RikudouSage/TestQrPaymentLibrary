package iban

type Validator interface {
	IsValid(iban Iban) bool
}
