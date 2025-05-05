package service

import (
	"fmt"
	"strconv"
	"wmh/internal/model"
	"wmh/pkg/database"
)

type ReminderService struct {
	// 可以添加其他依赖
}

// NewReminderService 创建提醒服务实例
func NewReminderService() *ReminderService {
	return &ReminderService{}
}

// GetUserReminders 获取用户的所有提醒
func (s *ReminderService) GetUserReminders(userID uint) ([]model.Reminder, error) {
	var reminders []model.Reminder
	if err := database.DB.Where("user_id = ?", userID).Find(&reminders).Error; err != nil {
		return nil, fmt.Errorf("获取提醒列表失败: %v", err)
	}
	return reminders, nil
}

// CreateReminder 创建新的提醒
func (s *ReminderService) CreateReminder(reminder *model.Reminder) error {
	if err := database.DB.Create(reminder).Error; err != nil {
		return fmt.Errorf("创建提醒失败: %v", err)
	}
	return nil
}

// UpdateReminder 更新提醒
func (s *ReminderService) UpdateReminder(reminderID string, reminder *model.Reminder) error {
	id, err := strconv.ParseUint(reminderID, 10, 64)
	if err != nil {
		return fmt.Errorf("无效的提醒ID: %v", err)
	}

	// 检查提醒是否存在
	var existingReminder model.Reminder
	if err := database.DB.First(&existingReminder, id).Error; err != nil {
		return fmt.Errorf("提醒不存在: %v", err)
	}

	// 验证提醒是否属于该用户
	if existingReminder.UserID != reminder.UserID {
		return fmt.Errorf("无权访问此提醒")
	}

	// 更新提醒
	reminder.ID = uint(id)
	if err := database.DB.Save(reminder).Error; err != nil {
		return fmt.Errorf("更新提醒失败: %v", err)
	}

	return nil
}

// DeleteReminder 删除提醒
func (s *ReminderService) DeleteReminder(userID uint, reminderID string) error {
	id, err := strconv.ParseUint(reminderID, 10, 64)
	if err != nil {
		return fmt.Errorf("无效的提醒ID: %v", err)
	}

	// 检查提醒是否存在
	var reminder model.Reminder
	if err := database.DB.First(&reminder, id).Error; err != nil {
		return fmt.Errorf("提醒不存在: %v", err)
	}

	// 验证提醒是否属于该用户
	if reminder.UserID != userID {
		return fmt.Errorf("无权访问此提醒")
	}

	// 删除提醒
	if err := database.DB.Delete(&reminder).Error; err != nil {
		return fmt.Errorf("删除提醒失败: %v", err)
	}

	return nil
}

func (s *ReminderService) SendReminder(reminder *model.Reminder) error {
	// 发送提醒（可以通过推送通知、短信或邮件）
	return nil
}

func (s *ReminderService) CheckAndTriggerReminders() error {
	// 定时检查并触发提醒
	return nil
}
