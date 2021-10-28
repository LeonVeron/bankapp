package card

import (
	"bank/pkg/bank/types"
	"fmt"
)

/*func ExampleWithdraw_positive()  {
	card:=types.Card{Balance: 20_000_00,Active: true}
	Withdraw(&card,10_000_00)
	fmt.Println(card.Balance)

	//Output: 1000000
}

func ExampleWithdraw_noMoney()  {
	card:=types.Card{Balance: 0,Active: true}
	Withdraw(&card,10_000_00)
	fmt.Println(card.Balance)

	//Output: 0
}

func ExampleWithdraw_limit()  {
	card:=types.Card{Balance: 20_000_00,Active: true}
	Withdraw(&card,30_000_00)
	fmt.Println(card.Balance)

	//Output: 2000000
}

func ExampleWithdraw_inactive()  {
	card:=types.Card{Balance: 20_000_00,Active: false}
	Withdraw(&card,30_000_00)
	fmt.Println(card.Balance)

	//Output: 2000000
}

func ExampleDeposit_positive()  {
	card:=types.Card{Balance: 20_000_00,Active: true}
	Deposit(&card,10_000_00)
	fmt.Println(card.Balance)

	//Output: 3000000
}

func ExampleDeposit_negative()  {
	card:=types.Card{Balance: 20_000_00,Active: true}
	Deposit(&card,-10_000_00)
	fmt.Println(card.Balance)

	//Output: 2000000
}

func ExampleDeposit_inactive()  {
	card:=types.Card{Balance: 20_000_00,Active: false}
	Deposit(&card,30_000_00)
	fmt.Println(card.Balance)

	//Output: 2000000
}

func ExampleDeposit_limit()  {
	card:=types.Card{Balance: 20_000_00,Active: true}
	Withdraw(&card,60_000_00)
	fmt.Println(card.Balance)

	//Output: 2000000
}*/

/*func ExampleAddBonus_positive()  {
	card:=types.Card{Balance: 10_000_00,Active: true,MinBalance: 10_000_00}
	AddBonus(&card,3,30,365)
	fmt.Println(card.Balance)

	//Output: 1002465
}

func ExampleAddBonus_inactive()  {
	card:=types.Card{Balance: 20_000_00,Active: false,MinBalance: 10_000_00}
	AddBonus(&card,3,30,365)
	fmt.Println(card.Balance)

	//Output: 2000000
}

func ExampleAddBonus_negativeBalance()  {
	card:=types.Card{Balance: -20_000_00,Active: true,MinBalance: 10_000_00}
	AddBonus(&card,3,30,365)
	fmt.Println(card.Balance)

	//Output: -2000000
}

func ExampleAddBonus_overLimit()  {
	card:=types.Card{Balance: 1,Active: true,MinBalance: 2_500_000_00}
	AddBonus(&card,3,30,365)
	fmt.Println(card.Balance)

	//Output: 1
}*/

func ExampleTotal_positive()  {
	cards:= []types.Card{
		{
			Balance: 10_000_00,
			Active: true,
		},
		{
			Balance: 20_000_00,
			Active: true,	
		},
	}
	fmt.Println(Total(cards))

	//Output: 3000000
}

/*func ExampleTotal_negative()  {
	cards:= []types.Card{
		{
			Balance: 10_000_00,
			Active: false,
		},
		{
			Balance: 20_000_00,
			Active: true,	
		},
		{
			Balance: 30_000_00,
			Active: true,	
		},
		{
			Balance: -10_000_00,
			Active: true,	
		},
		{
			Balance: -40_000_00,
			Active: false,	
		},
	}
	fmt.Println(Total(cards))

	//Output: 5000000
}*/	