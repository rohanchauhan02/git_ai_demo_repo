package repository

import (
	"github.com/jinzhu/gorm"
	"github.com/rohanchauhan02/automation-engine/models"
)

type Repository struct {
	DB *gorm.DB
}

func (r *Repository) GetTaskByDueDate(interval int64) ([]models.Task, error) {
	var tasks []models.Task
	// Convert interval to time.Duration
	intervalDuration := time.Duration(interval) * time.Minute
	// Get the current time
	currentTime := time.Now()
	// Calculate the due date
	dueDate := currentTime.Add(intervalDuration)
	// Fetch tasks due within the interval
	err := r.DB.Where("due_date <= ?", dueDate).Find(&tasks).Error
	if err != nil {
		return nil, err
	}
	return tasks, nil
}