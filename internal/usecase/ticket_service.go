package usecase

import (
	"cinema_service/internal/domain"
	"cinema_service/internal/port/driven"
	"cinema_service/internal/port/usecase"
	"errors"
	"fmt"
)

type ticketService struct {
	ticketRepo  driven.TicketRepository
	sessionRepo driven.SessionRepository
	movieRepo   driven.MovieRepository
	notifier    driven.NotificationSender
}

func NewTicketService(
	ticketRepo driven.TicketRepository,
	sessionRepo driven.SessionRepository,
	movieRepo driven.MovieRepository,
	notifier driven.NotificationSender,
) usecase.TicketUseCase {
	return &ticketService{
		ticketRepo:  ticketRepo,
		sessionRepo: sessionRepo,
		movieRepo:   movieRepo,
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

	// 4.1. Достаём фильм по MovieID из сеанса
	movie, err := s.movieRepo.GetByID(session.MovieID)
	if err != nil {
		return 0, err
	}
	if movie == nil {
		// чтобы не ломать покупку из-за уведомления, можно просто залогировать
		fmt.Println("movie not found for session:", session.ID)
		// и вернуть id без письма
		return id, nil
	}

	// 5. Отправляем уведомление
	if s.notifier != nil {
		if err := s.notifier.SendTicketBoughtNotification(t, movie.Title); err != nil {
			fmt.Println("failed to send notification:", err)
		} else {
			fmt.Printf("notification sent for ticket %d (movie %q)\n", t.ID, movie.Title)
		}
	} else {
		fmt.Println("notifier is nil, skipping notification")
	}

	return id, nil
}

func (s *ticketService) ListTicketsByUser(userID string) ([]domain.Ticket, error) {
	if userID == "" {
		return nil, errors.New("userID is required")
	}
	return s.ticketRepo.ListByUser(userID)
}
