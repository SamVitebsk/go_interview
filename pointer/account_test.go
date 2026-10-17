package pointer

import (
	"errors"
	"math"
	"testing"
)

func TestAccount_Deposit(t *testing.T) {
	tests := []struct {
		name        string
		account     *Account
		amount      int64
		wantBalance int64
		wantErr     error
	}{
		{
			name:        "успешное пополнение счета",
			account:     &Account{balance: 100},
			amount:      50,
			wantBalance: 150,
			wantErr:     nil,
		},
		{
			name:        "успешное пополнение ровно до MaxInt64",
			account:     &Account{balance: math.MaxInt64 - 100},
			amount:      100,
			wantBalance: math.MaxInt64,
			wantErr:     nil,
		},
		{
			name:        "ошибка: nil receiver возвращает ErrNilReceiver без panic",
			account:     nil,
			amount:      100,
			wantBalance: 0,
			wantErr:     ErrNilReceiver,
		},
		{
			name:        "ошибка: нулевая сумма (неположительная)",
			account:     &Account{balance: 100},
			amount:      0,
			wantBalance: 100, // баланс не должен измениться
			wantErr:     ErrNonPositiveAmount,
		},
		{
			name:        "ошибка: отрицательная сумма (неположительная)",
			account:     &Account{balance: 100},
			amount:      -50,
			wantBalance: 100, // баланс не должен измениться
			wantErr:     ErrNonPositiveAmount,
		},
		{
			name:        "ошибка: переполнение int64 при сложении",
			account:     &Account{balance: math.MaxInt64 - 10},
			amount:      11,
			wantBalance: math.MaxInt64 - 10, // баланс не должен измениться
			wantErr:     ErrOverflow,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.account.Deposit(tt.amount)

			// 1. Проверяем ошибку операции
			if tt.wantErr != nil {
				if !errors.Is(err, tt.wantErr) {
					t.Fatalf("Deposit() error = %v, wantErr %v", err, tt.wantErr)
				}
			} else if err != nil {
				t.Fatalf("Deposit() неожиданная ошибка: %v", err)
			}

			// 2. Если аккаунт не nil — проверяем баланс через метод Balance()
			if tt.account != nil {
				gotBalance, err := tt.account.Balance()
				if err != nil {
					t.Fatalf("Balance() ошибка чтения: %v", err)
				}
				if gotBalance != tt.wantBalance {
					t.Errorf("Balance() = %v, want %v", gotBalance, tt.wantBalance)
				}
			}
		})
	}
}

func TestAccount_Withdraw(t *testing.T) {
	tests := []struct {
		name        string
		account     *Account
		amount      int64
		wantBalance int64
		wantErr     error
	}{
		{
			name:        "успешное списание части средств",
			account:     &Account{balance: 100},
			amount:      40,
			wantBalance: 60,
			wantErr:     nil,
		},
		{
			name:        "успешное списание всей суммы под ноль",
			account:     &Account{balance: 100},
			amount:      100,
			wantBalance: 0,
			wantErr:     nil,
		},
		{
			name:        "ошибка: nil receiver возвращает ErrNilReceiver",
			account:     nil,
			amount:      50,
			wantBalance: 0,
			wantErr:     ErrNilReceiver,
		},
		{
			name:        "ошибка: нулевая сумма списания (неположительная)",
			account:     &Account{balance: 100},
			amount:      0,
			wantBalance: 100, // баланс не изменился
			wantErr:     ErrNonPositiveAmount,
		},
		{
			name:        "ошибка: отрицательная сумма списания (неположительная)",
			account:     &Account{balance: 100},
			amount:      -20,
			wantBalance: 100, // баланс не изменился
			wantErr:     ErrNonPositiveAmount,
		},
		{
			name:        "ошибка: недостаточный баланс",
			account:     &Account{balance: 50},
			amount:      51,
			wantBalance: 50, // баланс не изменился
			wantErr:     ErrInsufficientBalance,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.account.Withdraw(tt.amount)

			// 1. Проверяем ошибку операции
			if tt.wantErr != nil {
				if !errors.Is(err, tt.wantErr) {
					t.Fatalf("Withdraw() error = %v, wantErr %v", err, tt.wantErr)
				}
			} else if err != nil {
				t.Fatalf("Withdraw() неожиданная ошибка: %v", err)
			}

			// 2. Если аккаунт не nil — проверяем баланс через метод Balance()
			if tt.account != nil {
				gotBalance, err := tt.account.Balance()
				if err != nil {
					t.Fatalf("Balance() ошибка чтения: %v", err)
				}
				if gotBalance != tt.wantBalance {
					t.Errorf("Balance() = %v, want %v", gotBalance, tt.wantBalance)
				}
			}
		})
	}
}
