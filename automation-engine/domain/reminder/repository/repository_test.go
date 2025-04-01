package repository_test

import (
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

func TestGetTaskByDueDate(t *testing.T) {
	mockDB := new(MockDB)
	repo := &repository.Repository{DB: mockDB}
	mockDB.On("Where", "due_date <= ?", time.Now().Add(time.Duration(1)*time.Minute)).Return(&gorm.DB{})
	tasks, err := repo.GetTaskByDueDate(1)
	assert.Nil(t, err)
	assert.NotNil(t, tasks)
}