package model

import "time"

type Reminder struct {
    ID          uint      `json:"id" gorm:"primaryKey"`
    UserID      uint      `json:"user_id"`
    Type        string    `json:"type"` // daily_check, exercise, study, water
    Title       string    `json:"title"`
    Content     string    `json:"content"`
    Schedule    string    `json:"schedule"` // cron表达式
    IsActive    bool      `json:"is_active"`
    LastTrigger time.Time `json:"last_trigger"`
    CreatedAt   time.Time `json:"created_at"`
    UpdatedAt   time.Time `json:"updated_at"`
}