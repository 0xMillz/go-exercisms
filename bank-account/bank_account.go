package account

import (
	"sync"
)

type Account struct {
	balance int64
	open bool
}

var mutex = &sync.Mutex{}

func Open(initialDeposit int64) (account *Account) {
	if initialDeposit < 0 {
		account = nil
	} else {
		account = &Account{initialDeposit, true}
	}
	return account
}

func (a *Account) Balance() (balance int64, ok bool) {
	mutex.Lock()
	if a.open {
		ok = true
	} else {
		ok = false
	}
	balance = a.balance
	mutex.Unlock()
	return
}

func (a *Account) Close() (payout int64, ok bool) {
	mutex.Lock()
	if a.open {
		a.open = false
		ok = true
		payout = a.balance
		a.balance = 0
	}
	mutex.Unlock()
	return
}

func (a *Account) Deposit(amount int64) (newBalance int64, ok bool) {
	mutex.Lock()
	newBalance = a.balance + amount
	if newBalance >= 0 && a.open {
		ok = true
		a.balance = newBalance
	} else {
		newBalance = a.balance
	}
	mutex.Unlock()
	return
}

