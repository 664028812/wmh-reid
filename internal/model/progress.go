package model

import "time"

type Progress struct {
    ID          uint      `json:"id" gorm:"primaryKey"`
    UserID      uint      `json:"user_id"`
    GoalID      uint      `json:"goal_id"`
    Type        string    `json:"type"` // weight, study_time
    Value       float64   `json:"value"` // 体重或学习时长
    Note        string    `json:"note"`
    RecordDate  time.Time `json:"record_date"`
    CreatedAt   time.Time `json:"created_at"`
}