package repository

import (
	"github.com/jinzhu/gorm"
	"github.com/rohanchauhan02/automation-engine/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
	"time"
)

type MockDB struct {
	mock.Mock
}

func (m *MockDB) Where(query interface{}, args ...interface{}) *gorm.DB {
	return m.Called(query, args).Get(0).(*gorm.DB)
}

func (m *MockDB) Find(out interface{}) *gorm.DB {
	return m.Called(out).Get(0).(*gorm.DB)
}

func TestRepository_GetTaskByDueDate(t *testing.T) {
	// Mock DB
	mockDB := new(MockDB)
	// Mock current time
	mockTime := time.Now()
	// Mock interval
	mockInterval := int64(60)
	// Mock tasks
	mockTasks := []models.Task{{ID: 1, Name: 'Task 1', DueDate: mockTime.Add(time.Duration(mockInterval) * time.Minute)}}
	// Mock repository
	mockRepo := &Repository{DB: mockDB}
	// Mock DB calls
	mockDB.On("Where", "due_date <= ?", mock.Anything).Return(mockDB)
	mockDB.On("Find", &[]models.Task{}).Return(mockDB)
	// Call function
	tasks, err := mockRepo.GetTaskByDueDate(mockInterval)
	// Assert function call
	assert.NoError(t, err)
	assert.Equal(t, mockTasks, tasks)
}