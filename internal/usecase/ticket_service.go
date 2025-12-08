package usecase

import (
	"cinema_service/internal/domain"
	"cinema_service/internal/port/driven"
	"cinema_service/internal/port/usecase"
	"errors"
)

type ticketService struct {
	ticketRepo  driven.TicketRepository
	sessionRepo driven.SessionRepository
	notifier    driven.NotificationSender
}

func NewTicketService(
	ticketRepo driven.TicketRepository,
	sessionRepo driven.SessionRepository,
	notifier driven.NotificationSender,
) usecase.TicketUseCase {
	return &ticketService{
		ticketRepo:  ticketRepo,
		sessionRepo: sessionRepo,
		notifier:    notifier,
	}
}

func (s *ticketService) BuyTicket(t domain.Ticket) (int, error) {
	// 1. Проверяем, что сеанс существует
	session, err := s.sessionRepo.GetByID(t.SessionID)
	if err != nil {
		return 0, err
	}
	if session == nil {
		return 0, errors.New("session not found")
	}

	// 2. Проверка корректности ряда и места (валидация)
	if t.Row < 1 || t.Row > session.Rows {
		return 0, errors.New("invalid row number")
	}
	if t.Seat < 1 || t.Seat > session.SeatsPerRow {
		return 0, errors.New("invalid seat number")
	}

	// 3. Проверяем, свободно ли место
	busy, err := s.ticketRepo.IsSeatBusy(t.SessionID, t.Row, t.Seat)
	if err != nil {
		return 0, err
	}
	if busy {
		return 0, errors.New("seat already taken")
	}

	// 4. Создаём билет
	id, err := s.ticketRepo.Create(t)
	if err != nil {
		return 0, err
	}

	t.ID = id

	// 5. Отправляем уведомление (ошибку игнорируем, чтобы не ломать покупку)
	if s.notifier != nil {
		_ = s.notifier.SendTicketBoughtNotification(t)
	}

	return id, nil
}

func (s *ticketService) ListTicketsByUser(userID string) ([]domain.Ticket, error) {
	if userID == "" {
		return nil, errors.New("userID is required")
	}
	return s.ticketRepo.ListByUser(userID)
}
