package http

import (
	"cinema_service/internal/port/usecase"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type MovieHandler struct {
	uc usecase.MovieUseCase
}

func NewMovieHandler(uc usecase.MovieUseCase) *MovieHandler {
	return &MovieHandler{uc: uc}
}

func (h *MovieHandler) Register(r *gin.RouterGroup) {
	r.GET("/movies", h.ListMovies)
	r.GET("/movies/:id", h.GetMovieByID)
}

func (h *MovieHandler) ListMovies(c *gin.Context) {
	movies, err := h.uc.ListMovies()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to load movies",
		})
		return
	}

	c.JSON(http.StatusOK, movies)
}

func (h *MovieHandler) GetMovieByID(c *gin.Context) {
	idStr := c.Param("id")

	id, err := strconv.Atoi(idStr)
	if err != nil || id <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid movie id",
		})
		return
	}

	movie, err := h.uc.GetMovie(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to load movie",
		})
		return
	}

	if movie == nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "movie not found",
		})
		return
	}

	c.JSON(http.StatusOK, movie)
}
