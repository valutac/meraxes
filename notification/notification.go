package notification

import (
	"context"
	"fmt"

	"github.com/Valutac/meraxes/meraxes"

	"github.com/olebedev/go-tgbot"
	"github.com/olebedev/go-tgbot/client/messages"
	"github.com/olebedev/go-tgbot/models"
	"go.uber.org/zap"
	"gopkg.in/gomail.v2"
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
	smtpSender    string
	smtpHost      string
	smtpPort      int
	smtpUsername  string
	smtpPassword  string
}

// NewService return new instance of Notfication Service
func NewService(logger *zap.Logger) *Service {
	return &Service{logger: logger}
}

// SetTelegramToken update telegram token in service
func (s *Service) SetTelegramToken(token string) {
	s.telegramToken = token
}

// SetSMTP update smtp config in service
func (s *Service) SetSMTP(sender, host, username, password string, port int) {
	s.smtpSender = sender
	s.smtpHost = host
	s.smtpUsername = username
	s.smtpPassword = password
	s.smtpPort = port
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
		return s.notifyUsingEmail(statuses, target)
	case TypeTelegram:
		return s.notifyUsingTelegram(statuses, target)
	}
	return nil
}

func (s *Service) notifyUsingEmail(statuses []meraxes.Status, target string) error {
	m := gomail.NewMessage()
	m.SetHeader("From", s.smtpSender)
	m.SetHeader("To", target)
	m.SetHeader("Subject", "[Meraxes] Alert! Hosts Down")

	var message = fmt.Sprintf("<strong>%d hosts down :(</strong>", len(statuses))
	message += "<ul>"
	for _, status := range statuses {
		message += fmt.Sprintf("<li>%s (%s)</li>", status.Host, status.URI)
	}
	message += "</ul>"
	m.SetBody("text/html", message)

	d := gomail.NewDialer(s.smtpHost, s.smtpPort, s.smtpUsername, s.smtpPassword)

	if err := d.DialAndSend(m); err != nil {
		s.logger.Error("failed to send message", zap.Error(err))
		return err
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
