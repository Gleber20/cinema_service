package usecase

import "cinema_service/internal/domain"

type TicketUseCase interface {
	BuyTicket(t domain.Ticket) (int, error)
	ListTicketsByUser(userID string) ([]domain.Ticket, error) // по JWT из auth_service
}
