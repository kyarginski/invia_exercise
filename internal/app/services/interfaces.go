package services

import (
	"context"
	"log/slog"

	"invia/internal/app/health"
	"invia/internal/app/repository"
)

type IService interface {
	Logger() *slog.Logger
	Ping(ctx context.Context) bool
	Close() error

	repository.IUserRepository

	health.LivenessChecker
	health.ReadinessChecker
}
