package main

import (
	"cinema_service/internal/adapter/driven/notification"
	"cinema_service/internal/adapter/driven/postgres"
	"cinema_service/internal/adapter/driving"
	"cinema_service/internal/config"
	"cinema_service/internal/database"
	"cinema_service/internal/port/driven"
	"cinema_service/internal/usecase"
	"context"
	"fmt"
	"log"

	"github.com/joho/godotenv"
	"github.com/sethvargo/go-envconfig"
)

func main() {
	// 1. Загружаем .env в окружение
	if err := godotenv.Load(); err != nil {
		log.Println("⚠️  .env file not found, using system environment")
	}

	var cfg config.Config

	// 2. Читаем переменные окружения
	err := envconfig.ProcessWith(context.TODO(), &envconfig.Config{
		Target:   &cfg,
		Lookuper: envconfig.OsLookuper(),
	})
	if err != nil {
		panic(err)
	}

	db, err := database.NewPostgresDB(&cfg)
	if err != nil {
		log.Fatalf("failed to connect db: %v", err)
	}

	movieRepo := postgres.NewMovieRepo(db)
	sessionRepo := postgres.NewSessionRepo(db)
	ticketRepo := postgres.NewTicketRepo(db)

	var notifier driven.NotificationSender // если хочешь, можешь явно типом порт указать
	amqpSender, err := notification.NewAMQPNotificationSender(cfg.AMQP_URL, cfg.AMQP_QUEUE)
	if err != nil {
		log.Printf("⚠ could not init notifier: %v (notifications disabled)", err)
		notifier = nil
	} else {
		if amqpSender != nil {
			notifier = amqpSender
			defer amqpSender.Close()
		}
	}

	movieUC := usecase.NewMovieService(movieRepo)
	sessionUC := usecase.NewSessionService(sessionRepo)
	ticketUC := usecase.NewTicketService(ticketRepo, sessionRepo, notifier)

	r := driving.SetupRouter(movieUC, sessionUC, ticketUC)

	fmt.Println("cinema_service running on port:", cfg.HTTPPort)
	if err := r.Run(":" + cfg.HTTPPort); err != nil {
		log.Fatalf("server error: %v", err)
	}
}
