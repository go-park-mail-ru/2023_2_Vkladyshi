package usecase

import (
	"context"

	"github.com/go-redis/redis/v8"
	"github.com/stretchr/testify/mock"
)

type MockRedisClient struct {
	mock.Mock
}

func (m *MockRedisClient) Process(ctx context.Context, cmd redis.Cmder) error {
	args := m.Called(ctx, cmd)
	return args.Error(0)
}

func (m *MockRedisClient) ProcessPipeline(ctx context.Context, cmds []redis.Cmder) error {
	args := m.Called(ctx, cmds)
	return args.Error(0)
}

func NewMockRedisClient() *MockRedisClient {
	return &MockRedisClient{}
}

// func TestAddSession(t *testing.T) {
// 	mockClient := NewMockRedisClient()

// 	sessionRepo := &session.SessionRepo{
// 		sessionRedisClient: mockClient.Mock(),
// 		Connection:         true,
// 	}

// 	mockClient.On("Process", mock.Anything, mock.AnythingOfType("*redis.SetCmd")).Return(nil)

// 	_, err := sessionRepo.AddSession(session.Session{SID: "123", Login: "user"}, nil)

// 	mockClient.AssertExpectations(t)

// 	if err != nil {
// 		t.Errorf("AddSession returned an error: %v", err)
// 	}
// }
