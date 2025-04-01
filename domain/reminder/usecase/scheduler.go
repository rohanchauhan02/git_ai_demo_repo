package usecase

import (
	"fmt"
	"log"
	"time"
)

func (u *usecase) Reminder() {
	// Run reminder in every minute
	const interval = 1
	ticker := time.NewTicker(interval * time.Minute)

	for {
		select {
		case <-ticker.C:
			tasks, err := u.repository.GetTaskByDueDate(interval)
			if err != nil {
				log.Println("Error fetching data from DB:", err)
			}
			if len(tasks) == 0 {
				fmt.Println("No task in bucket")
			}
			for _, task := range tasks {
				userID := task.UserID
				user, err := u....