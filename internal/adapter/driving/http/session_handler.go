package http

import (
	"cinema_service/internal/port/usecase"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type SessionHandler struct {
	uc usecase.SessionUseCase
}

func NewSessionHandler(uc usecase.SessionUseCase) *SessionHandler {
	return &SessionHandler{uc: uc}
}

func (h *SessionHandler) Register(r *gin.RouterGroup) {
	// Сеансы фильма
	r.GET("/movies/:id/sessions", h.ListByMovie)

	// Один сеанс по id
	r.GET("/sessions/:id", h.GetByID)
}

// GET /api/movies/:id/sessions
func (h *SessionHandler) ListByMovie(c *gin.Context) {
	idStr := c.Param("id")

	movieID, err := strconv.Atoi(idStr)
	if err != nil || movieID <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid movie id",
		})
		return
	}

	sessions, err := h.uc.ListSessionsByMovie(movieID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to load sessions",
		})
		return
	}

	c.JSON(http.StatusOK, sessions)
}

// GET /api/sessions/:id
func (h *SessionHandler) GetByID(c *gin.Context) {
	idStr := c.Param("id")

	sessionID, err := strconv.Atoi(idStr)
	if err != nil || sessionID <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid session id",
		})
		return
	}

	session, err := h.uc.GetSession(sessionID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to load session",
		})
		return
	}

	if session == nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "session not found",
		})
		return
	}

	c.JSON(http.StatusOK, session)
}
