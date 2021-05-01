package model

type Customer struct {
	ID        int
	Name      string
	Documents Documents
}

type Documents struct {
	RG  string
	CPF string
}
