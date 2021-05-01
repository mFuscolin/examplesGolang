package model

import "time"

type Debt struct {
	Value      float32
	IDCustomer int
	DueDate    time.Time
	Customer   Customer
	Products   []Product
}
