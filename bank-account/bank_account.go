package account

import "sync"

// Account has room from some money... int64 ?!
type Account struct {
	balance int64
	open    bool
	mutex   sync.Mutex
}

// Close closes the account and returns the current balance
func (a *Account) Close() (payout int64, ok bool) {
	a.mutex.Lock() // whenever we modify the account, it needs to be locked
	defer a.mutex.Unlock()
	if !a.open { // check if the account is actually open
		return 0, false
	}
	a.open = false
	payout, a.balance = a.balance, 0
	return payout, true
}

// Open creates a new account with initialDeposit as balance
func Open(initialDeposit int64) *Account {
	if initialDeposit < 0 {
		return nil
	}
	a := new(Account)
	a.balance = initialDeposit
	a.open = true
	return a
}

// Balance returns the current balance of the account
func (a *Account) Balance() (balance int64, ok bool) {
	if !a.open { // check if the account is actually open
		return 0, false
	}
	return a.balance, true
}

// Deposit changes the balance of an account by the specified amount
func (a *Account) Deposit(amount int64) (newBalance int64, ok bool) {
	a.mutex.Lock() // whenever we modify the account, it needs to be locked
	defer a.mutex.Unlock()
	if !a.open { // check if the account is actually open
		return 0, false
	}
	if a.balance+amount < 0 {
		return a.balance, false
	}
	a.balance += amount
	return a.balance, true
}
