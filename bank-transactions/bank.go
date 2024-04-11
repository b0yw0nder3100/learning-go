package main

import "fmt"

//See Daily transaction of users
type user struct {
	name        string
	inputtedPin int
	action
	accountDetails
}

type action struct {
	name   string
	amount int
}

type accountDetails struct {
	balance int
	pin     int
}

type userActions interface {
	getUserAction() string
	getUserName() string
	withdraw() int
	check() int
	transfer() int
}

func (u user) getUserName() string {
	return u.name
}

func (u user) getUserAction() string {
	return u.action.name
}

func (u user) withdraw() int {
	if u.balance > u.amount {
		return u.balance - u.amount
	}
	return 0
}

func (u user) check() int {
	return u.balance
}

func (u user) transfer() int {
	return u.balance - u.amount
}

func test(ua userActions) {
	fmt.Println("------------------------------------------------------------")
	fmt.Println("", ua.getUserName(), ua.getUserAction(), ua.withdraw())
}

func main() {
	fmt.Println("Daily Transactions")
	u1 := user{
		name:        "Marty",
		inputtedPin: 2123,
		action: action{
			name:   "withdraw",
			amount: 300,
		},
		accountDetails: accountDetails{
			balance: 40000,
			pin:     2000,
		},
	}
	test(u1)
	u2 := user{
		name:        "John",
		inputtedPin: 1350,
		action: action{
			name:   "Check",
			amount: 0,
		},
		accountDetails: accountDetails{
			balance: 3000,
			pin:     1350,
		},
	}
	test(u2)
}
