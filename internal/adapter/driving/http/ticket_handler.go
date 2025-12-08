package http

import (
	"fmt"
	"net/http"

	"cinema_service/internal/domain"
	"cinema_service/internal/port/usecase"

	"github.com/gin-gonic/gin"
)

type TicketHandler struct {
	uc usecase.TicketUseCase
}

func NewTicketHandler(uc usecase.TicketUseCase) *TicketHandler {
	return &TicketHandler{uc: uc}
}

func (h *TicketHandler) Register(r *gin.RouterGroup) {
	r.POST("/tickets/buy", h.BuyTicket)
	r.GET("/me/tickets", h.ListMyTickets)
}

// POST /api/tickets/buy
func (h *TicketHandler) BuyTicket(c *gin.Context) {
	var req struct {
		SessionID int    `json:"session_id"`
		Row       int    `json:"row"`
		Seat      int    `json:"seat"`
		Email     string `json:"email"`
	}

	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		return
	}

	userIDVal, ok := c.Get(UserIDCtx)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "user not found in context"})
		return
	}

	userIDInt, ok := userIDVal.(int)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "invalid user id type"})
		return
	}

	ticket := domain.Ticket{
		SessionID: req.SessionID,
		Row:       req.Row,
		Seat:      req.Seat,
		UserID:    fmt.Sprintf("%d", userIDInt),
		Email:     req.Email,
		IsPaid:    true,
	}

	id, err := h.uc.BuyTicket(ticket)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"ticket_id": id,
	})
}

// GET /api/me/tickets?user_id=xxx
func (h *TicketHandler) ListMyTickets(c *gin.Context) {
	userIDVal, ok := c.Get(UserIDCtx)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "user not found in context"})
		return
	}
	userIDInt, ok := userIDVal.(int)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "invalid user id type"})
		return
	}

	userID := fmt.Sprintf("%d", userIDInt)

	tickets, err := h.uc.ListTicketsByUser(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to load tickets"})
		return
	}

	c.JSON(http.StatusOK, tickets)
}
