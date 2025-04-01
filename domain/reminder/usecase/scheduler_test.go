package usecase_test

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

type MockRepository struct {
	mock.Mock
}

func (m *MockRepository) GetTaskByDueDate(interval int64) ([]models.Task, error) {
	args := m.Called(interval)
	return args.Get(0).([]models.Task), args.Error(1)
}

func TestReminder(t *testing.T) {
	mockRepo := new(MockRepository)
	usecase := &usecase.Usecase{repository: mockRepo}
	mockRepo.On("GetTaskByDueDate", 1).Return([]models.Task{}, nil)
	usecase.Reminder()
	mockRepo.AssertExpectations(t)
}