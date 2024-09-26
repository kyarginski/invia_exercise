package mocks

import (
	"context"
	"log/slog"

	"github.com/stretchr/testify/mock"

	models "invia/api/restapi"
)

// MockIService - мок для интерфейса IService
type MockIService struct {
	mock.Mock
}

// AddUser - мок метода AddUser
func (m *MockIService) AddUser(ctx context.Context, user *models.User) error {
	args := m.Called(ctx, user)
	return args.Error(0)
}

// GetUserById - мок метода GetUserById
func (m *MockIService) GetUserById(ctx context.Context, id int) (*models.User, error) {
	args := m.Called(ctx, id)
	if args.Get(0) != nil {
		return args.Get(0).(*models.User), args.Error(1)
	}
	return nil, args.Error(1)
}

// UpdateUser - мок метода UpdateUser
func (m *MockIService) UpdateUser(ctx context.Context, id int, user *models.User) error {
	args := m.Called(ctx, id, user)
	return args.Error(0)
}

// DeleteUser - мок метода DeleteUser
func (m *MockIService) DeleteUser(ctx context.Context, id int) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

// ListUsers - мок метода ListUsers
func (m *MockIService) ListUsers(ctx context.Context, page, limit int) ([]*models.User, error) {
	args := m.Called(ctx, page, limit)
	return args.Get(0).([]*models.User), args.Error(1)
}

// Logger - мок метода Logger
func (m *MockIService) Logger() *slog.Logger {
	args := m.Called()
	return args.Get(0).(*slog.Logger)
}

// Ping - мок метода Ping
func (m *MockIService) Ping(ctx context.Context) bool {
	args := m.Called(ctx)
	return args.Bool(0)
}

// Close - мок метода Close
func (m *MockIService) Close() error {
	args := m.Called()
	return args.Error(0)
}

// LivenessCheck - мок метода LivenessCheck (для проверки жизнеспособности сервиса)
func (m *MockIService) LivenessCheck() bool {
	args := m.Called()
	return args.Bool(0)
}

// ReadinessCheck - мок метода ReadinessCheck (для проверки готовности сервиса)
func (m *MockIService) ReadinessCheck() bool {
	args := m.Called()
	return args.Bool(0)
}
