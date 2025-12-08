package driven

import "cinema_service/internal/domain"

type NotificationSender interface {
	SendTicketBoughtNotification(ticket domain.Ticket) error
}
