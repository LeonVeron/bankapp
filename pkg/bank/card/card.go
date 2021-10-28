package card

import (
	"bank/pkg/bank/types"
)

//IssueCard issues a new card
func IssueCard(currency types.Currency, color string, name string) types.Card  {
	Card:=types.Card {
		ID: 1000,
		PAN: "5058 xxxx xxxx 0001",
		Balance: 0,
		Currency: currency,
		Color: color,
		Name: name,
		Active: true,
}
	return Card
}

const withdrawLimit=20_000_00

//Withdraw withdraws cash of card
func Withdraw(card *types.Card, amount types.Money) {
	
	if amount<0{
		return 
	}
	if amount>withdrawLimit{
		return 
	}
	if !card.Active{
		return 
	}
	if card.Balance<amount{
		return 
	}
	
	card.Balance-=amount

} 

const depositLimit=50_000_00

//func deposits money to account
func Deposit(card *types.Card, amount types.Money)  {
	if !card.Active{
		return 
	}
	if amount>depositLimit{
		return 
	}
	if amount<0{
		return 
	}
	card.Balance+=amount
}

const maxBonus=5000_00

//AddBonus adds persent to account
func AddBonus(card *types.Card,persent int,dayslnMonth int,dayslnYear int)  {
	bonus:=0

	if !card.Active{
		return 
	}
	if card.Balance<=0{
		return
	}
	if card.MinBalance<10_000{
		return
	}
	bonus=int(card.MinBalance)*persent/100*dayslnMonth/dayslnYear
	
	if bonus<=maxBonus{
		Deposit(card,types.Money(bonus))
		return
	}
	bonus=0
	Deposit(card,types.Money(bonus))
}

//Total вычисляет общую сумму средств на всех картах
func Total(cards []types.Card) types.Money {
	sum:=types.Money(0)
	for _, card:=range cards {
		if !card.Active{
			continue
		}
		if card.Balance<=0 {
			continue
		}
		sum +=card.Balance
	}
	return sum
}