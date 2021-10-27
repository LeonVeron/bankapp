package types

//Money represents 
type Money int64

//Currency представляет код валюты
type Currency string

//коды валют
const (
	TJS Currency="TJS"
	RUB Currency="RUB"
	USD Currency="USD"
)

//PAN представляет номер карты
type PAN string

//Card представляет информацию о карте
type Card struct {
	ID int
	PAN PAN
	Balance Money
	Currency Currency
	Color string
	Name string
	Active bool
	MinBalance Money
}

//Payment shows info about a payment
type Payment struct {
	ID int
	Amount Money
}