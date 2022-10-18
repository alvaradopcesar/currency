package repository

import "time"

type Config struct {
	DbTimeOUT int
}

type Currency struct {
	ID         string
	CustomerID int
	Code       string
	Value      float64
	CreatedAt  time.Time
}

//type GetAllCurrency struct {
//	Code      string
//	Value     float64
//	CreatedAt time.Time
//}
//
//type GetByIDCurrency struct {
//	CustomerID int
//	Code       string
//	Value      float64
//	CreatedAt  time.Time
//}

type CurrencyFilter struct {
	Code  string
	FInit time.Time
	FEnd  time.Time
}

type InsertCurrency struct {
	Code  string
	Value float64
}

type InsertQuery struct {
	Method  string
	Address string
	Code    int
	Time    float64
}
