package services

import (
	"context"
	"fmt"
	"log/slog"

	"invia/internal/app/repository"
)

type MyService struct {
	log     *slog.Logger
	storage *repository.Storage
}

func NewService(log *slog.Logger, connectString string) (IService, error) {
	const op = "service.NewService"

	storage, err := repository.New(connectString)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return &MyService{
		log:     log,
		storage: storage,
	}, nil
}

// Close closes DB connection.
func (s *MyService) Close() error {
	return s.storage.Close()
}

func (s *MyService) LivenessCheck() bool {
	// Implement liveness check logic
	return true
}

func (s *MyService) ReadinessCheck() bool {
	// Implement readiness check logic
	return s.Ping(context.Background())
}

func (s *MyService) Ping(ctx context.Context) bool {
	return s.storage.GetDB().PingContext(ctx) == nil
}

func (s *MyService) Logger() *slog.Logger {
	return s.log
}
