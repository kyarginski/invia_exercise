package services

import (
	"context"
	"log/slog"

	"invia/internal/app/health"
)

type IService interface {
	Logger() *slog.Logger
	Ping(ctx context.Context) bool
	Close() error

	health.LivenessChecker
	health.ReadinessChecker
}
