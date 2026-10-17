package pointer

import (
	"errors"
	"math"
)

var (
	// ErrNilReceiver возвращается при вызове метода у nil указателя.
	ErrNilReceiver = errors.New("receiver is nil")

	// ErrNonPositiveAmount возвращается, если сумма операции неположительна (<= 0).
	ErrNonPositiveAmount = errors.New("amount must be positive")

	// ErrInsufficientBalance возвращается при попытке списать больше, чем есть на балансе.
	ErrInsufficientBalance = errors.New("insufficient balance")

	// ErrOverflow возвращается при превышении максимального значения int64 при пополнении.
	ErrOverflow = errors.New("balance overflow")
)

// Account представляет банковский счет с поддержкой операций пополнения и списания.
type Account struct {
	balance int64
}

// Balance возвращает текущий баланс счета.
func (a *Account) Balance() (int64, error) {
	if a == nil {
		return 0, ErrNilReceiver
	}
	return a.balance, nil
}

// Deposit пополняет счет на указанную сумму.
func (a *Account) Deposit(amount int64) error {
	if a == nil {
		return ErrNilReceiver
	}
	if amount <= 0 {
		return ErrNonPositiveAmount
	}
	if a.balance > math.MaxInt64-amount {
		return ErrOverflow
	}

	a.balance += amount
	return nil
}

// Withdraw списывает указанную сумму со счета.
func (a *Account) Withdraw(amount int64) error {
	if a == nil {
		return ErrNilReceiver
	}
	if amount <= 0 {
		return ErrNonPositiveAmount
	}
	if a.balance < amount {
		return ErrInsufficientBalance
	}

	a.balance -= amount
	return nil
}
