package driving

import (
	httpHandlers "cinema_service/internal/adapter/driving/http"
	"cinema_service/internal/port/usecase"

	"github.com/gin-gonic/gin"
)

func SetupRouter(
	movieUC usecase.MovieUseCase,
	sessionUC usecase.SessionUseCase,
	ticketUC usecase.TicketUseCase,
) *gin.Engine {

	r := gin.Default()

	api := r.Group("/api")

	movieHandler := httpHandlers.NewMovieHandler(movieUC)
	movieHandler.Register(api)

	sessionHandler := httpHandlers.NewSessionHandler(sessionUC)
	sessionHandler.Register(api)

	authGroup := api.Group("")
	authGroup.Use(httpHandlers.CheckUserAuthentication)

	ticketHandler := httpHandlers.NewTicketHandler(ticketUC)
	ticketHandler.Register(authGroup)

	return r
}
