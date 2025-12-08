package driven

import "cinema_service/internal/domain"

type TicketRepository interface {
	Create(t domain.Ticket) (int, error)
	IsSeatBusy(sessionID, row, seat int) (bool, error)
	// Получить билеты пользователя (по userID из auth_service)
	ListByUser(userID string) ([]domain.Ticket, error)
}
