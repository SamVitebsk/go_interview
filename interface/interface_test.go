package _interface

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const (
	testRecipient = "user@example.com"
	testText      = "Добро пожаловать!"
)

// ErrorSender — тестовая реализация интерфейса Sender с заданной ошибкой.
type ErrorSender struct {
	Error error
}

// Send возвращает заданную ошибку ErrorSender, соблюдая контракт интерфейса Sender.
func (s *ErrorSender) Send(_ context.Context, _, _ string) error {
	if s == nil {
		return ErrNilReceiver
	}
	return s.Error
}

var _ Sender = (*ErrorSender)(nil)

func TestNewNotificationService(t *testing.T) {
	t.Run("отклонение nil зависимости", func(t *testing.T) {
		service, err := NewNotificationService(nil)

		assert.Nil(t, service)
		assert.ErrorIs(t, err, ErrNilSender)
	})

	t.Run("успешное создание сервиса", func(t *testing.T) {
		sender := &InMemorySender{}
		service, err := NewNotificationService(sender)

		require.NoError(t, err)
		assert.NotNil(t, service)
	})
}

func TestNotificationService_Notify(t *testing.T) {
	t.Run("успешная отправка и сохранение сообщения", func(t *testing.T) {
		sender := &InMemorySender{}
		service, err := NewNotificationService(sender)
		require.NoError(t, err)

		err = service.Notify(context.Background(), testRecipient, testText)

		require.NoError(t, err)
		assert.Equal(t, []Message{
			{Recipient: testRecipient, Text: testText},
		}, sender.Messages)
	})

	t.Run("валидация пустого получателя", func(t *testing.T) {
		sender := &InMemorySender{}
		service, err := NewNotificationService(sender)
		require.NoError(t, err)

		errEmpty := service.Notify(context.Background(), "", testText)
		errSpaces := service.Notify(context.Background(), "   ", testText)

		assert.ErrorIs(t, errEmpty, ErrEmptyRecipient)
		assert.ErrorIs(t, errSpaces, ErrEmptyRecipient)
		assert.Empty(t, sender.Messages)
	})

	t.Run("отмена контекста возвращает ошибку", func(t *testing.T) {
		sender := &InMemorySender{}
		service, err := NewNotificationService(sender)
		require.NoError(t, err)

		ctx, cancel := context.WithCancel(context.Background())
		cancel()

		err = service.Notify(ctx, testRecipient, testText)

		require.Error(t, err)
		assert.ErrorIs(t, err, context.Canceled)
		assert.Empty(t, sender.Messages)
	})

	t.Run("проверка оборачивания ошибки через errors.Is", func(t *testing.T) {
		customErr := errors.New("network failure")
		errorSender := &ErrorSender{Error: customErr}
		service, err := NewNotificationService(errorSender)
		require.NoError(t, err)

		err = service.Notify(context.Background(), testRecipient, testText)

		require.Error(t, err)
		assert.ErrorIs(t, err, customErr)
	})
}
