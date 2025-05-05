package handler

import "wmh/internal/service"

type Handler struct {
	userService     *service.UserService
	goalService     *service.GoalService
	reminderService *service.ReminderService
}

// NewHandler 创建handler实例
func NewHandler(userService *service.UserService, goalService *service.GoalService, reminderService *service.ReminderService) *Handler {
	return &Handler{
		userService:     userService,
		goalService:     goalService,
		reminderService: reminderService,
	}
}
