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
	err := r.DB.Where("due_date <= ?", time.Now().Add(time.Duration(interval)*time.Minute)).Find(&tasks).Error
	if err != nil {
		return nil, err
	}
	return tasks, nil
}