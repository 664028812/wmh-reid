package model

import "time"

type User struct {
    ID        uint      `json:"id" gorm:"primaryKey"`
    Username  string    `json:"username" gorm:"unique"`
    Password  string    `json:"-" gorm:"not null"`
    Email     string    `json:"email"`
    Phone     string    `json:"phone"`
    Role      string    `json:"role" gorm:"default:'user'"` // 用户角色：user, admin
    GoalType  string    `json:"goal_type"` // study, weight_loss
    CreatedAt time.Time `json:"created_at"`
    UpdatedAt time.Time `json:"updated_at"`
}