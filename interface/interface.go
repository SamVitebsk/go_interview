package _interface

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"sync"
)

var (
	// ErrNilSender возвращается, если в конструктор передана nil-зависимость Sender.
	ErrNilSender = errors.New("sender cannot be nil")

	// ErrEmptyRecipient возвращается, если получатель уведомления не указан.
	ErrEmptyRecipient = errors.New("recipient cannot be empty")

	// ErrNilReceiver возвращается при вызове метода на nil-ресивере.
	ErrNilReceiver = errors.New("receiver cannot be nil")
)

// Sender определяет контракт для отправки сообщений.
type Sender interface {
	Send(ctx context.Context, recipient, text string) error
}

// Message описывает отправляемое сообщение.
type Message struct {
	Recipient string
	Text      string
}

// InMemorySender хранит отправленные сообщения в оперативной памяти.
type InMemorySender struct {
	mu       sync.Mutex
	Messages []Message
}

// Send сохраняет сообщение в памяти с учётом отмены контекста и синхронизации.
func (s *InMemorySender) Send(ctx context.Context, recipient, text string) error {
	if s == nil {
		return ErrNilReceiver
	}

	if err := ctx.Err(); err != nil {
		return err
	}

	s.mu.Lock()
	defer s.mu.Unlock()

	s.Messages = append(s.Messages, Message{Recipient: recipient, Text: text})
	return nil
}

var _ Sender = (*InMemorySender)(nil)

// NotificationService отвечает за отправку уведомлений через внешнюю зависимость Sender.
type NotificationService struct {
	sender Sender
}

// NewNotificationService создает сервис уведомлений, проверяя зависимость на nil.
func NewNotificationService(sender Sender) (*NotificationService, error) {
	if sender == nil {
		return nil, ErrNilSender
	}
	return &NotificationService{sender: sender}, nil
}

// Notify валидирует входные данные, проверяет контекст и передает отправку в Sender.
func (n *NotificationService) Notify(ctx context.Context, recipient, text string) error {
	if n == nil {
		return ErrNilReceiver
	}

	if n.sender == nil {
		return ErrNilSender
	}

	if strings.TrimSpace(recipient) == "" {
		return ErrEmptyRecipient
	}

	if err := ctx.Err(); err != nil {
		return fmt.Errorf("notification context error: %w", err)
	}

	if err := n.sender.Send(ctx, recipient, text); err != nil {
		return fmt.Errorf("failed to send notification: %w", err)
	}

	return nil
}
