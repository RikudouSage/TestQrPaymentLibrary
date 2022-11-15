package iban

type Iban interface {
	String() string
	Validator() Validator
}
