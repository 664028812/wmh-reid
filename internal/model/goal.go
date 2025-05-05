package model

import "time"

type Goal struct {
    ID          uint      `json:"id" gorm:"primaryKey"`
    UserID      uint      `json:"user_id"`
    Type        string    `json:"type"` // study, weight_loss
    Target      float64   `json:"target"` // 目标体重或学习时长
    StartValue  float64   `json:"start_value"` // 起始体重或0
    CurrentValue float64  `json:"current_value"`
    Deadline    time.Time `json:"deadline"`
    Status      string    `json:"status"` // ongoing, completed, failed
    CreatedAt   time.Time `json:"created_at"`
    UpdatedAt   time.Time `json:"updated_at"`
}