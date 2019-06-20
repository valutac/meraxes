package notification

import (
	"context"
	"fmt"

	"github.com/ngurajeka/meraxes/meraxes"

	tgbot "github.com/olebedev/go-tgbot"
	"github.com/olebedev/go-tgbot/client/messages"
	"github.com/olebedev/go-tgbot/models"
	"go.uber.org/zap"
)

const (
	// TypeEmail notification using email
	TypeEmail = "email"
	// TypeTelegram notification using telegram
	TypeTelegram = "telegram"
)

// Service the notification service
type Service struct {
	logger        *zap.Logger
	telegramToken string
}

// NewService return new instance of Notfication Service
func NewService(logger *zap.Logger) *Service {
	return &Service{logger, ""}
}

// SetTelegramToken update telegram token in service
func (s *Service) SetTelegramToken(token string) {
	s.telegramToken = token
}

// NotifyAll send notification to user, or targets
func (s *Service) NotifyAll(statuses []meraxes.Status, targets []string, notificationType string) {
	for _, target := range targets {
		s.Notify(statuses, target, notificationType)
	}
}

// Notify send notification to user, or target
func (s *Service) Notify(statuses []meraxes.Status, target, notificationType string) error {
	switch notificationType {
	case TypeEmail:
		return nil
	case TypeTelegram:
		return s.notifyUsingTelegram(statuses, target)
	}
	return nil
}

func (s *Service) notifyUsingTelegram(statuses []meraxes.Status, target string) error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	api := tgbot.NewClient(ctx, s.telegramToken)

	var message = fmt.Sprintf("%d hosts down\n", len(statuses))
	for i, status := range statuses {
		message += fmt.Sprintf("%d. %s (%s)\n", i+1, status.Host, status.URI)
	}

	if _, err := api.Messages.SendMessage(
		messages.NewSendMessageParams().WithBody(&models.SendMessageBody{
			Text:   &message,
			ChatID: target,
		}),
	); err != nil {
		s.logger.Error("failed to send message", zap.Error(err))
		return err
	}

	return nil
}
