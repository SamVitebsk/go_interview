package pointer

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
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
			wantBalance: 100,
			wantErr:     ErrNonPositiveAmount,
		},
		{
			name:        "ошибка: отрицательная сумма (неположительная)",
			account:     &Account{balance: 100},
			amount:      -50,
			wantBalance: 100,
			wantErr:     ErrNonPositiveAmount,
		},
		{
			name:        "ошибка: переполнение int64 при сложении",
			account:     &Account{balance: math.MaxInt64 - 10},
			amount:      11,
			wantBalance: math.MaxInt64 - 10,
			wantErr:     ErrOverflow,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.account.Deposit(tt.amount)

			if tt.wantErr != nil {
				assert.ErrorIs(t, err, tt.wantErr)
			} else {
				assert.NoError(t, err)
			}

			if tt.account != nil {
				gotBalance, err := tt.account.Balance()
				require.NoError(t, err)
				assert.Equal(t, tt.wantBalance, gotBalance)
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
			wantBalance: 100,
			wantErr:     ErrNonPositiveAmount,
		},
		{
			name:        "ошибка: отрицательная сумма списания (неположительная)",
			account:     &Account{balance: 100},
			amount:      -20,
			wantBalance: 100,
			wantErr:     ErrNonPositiveAmount,
		},
		{
			name:        "ошибка: недостаточный баланс",
			account:     &Account{balance: 50},
			amount:      51,
			wantBalance: 50,
			wantErr:     ErrInsufficientBalance,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.account.Withdraw(tt.amount)

			if tt.wantErr != nil {
				assert.ErrorIs(t, err, tt.wantErr)
			} else {
				assert.NoError(t, err)
			}

			if tt.account != nil {
				gotBalance, err := tt.account.Balance()
				require.NoError(t, err)
				assert.Equal(t, tt.wantBalance, gotBalance)
			}
		})
	}
}
