package handler

import (
	"net/http"
	"wmh/internal/model"

	"github.com/gin-gonic/gin"
)

// GetReminders 获取用户的提醒列表
func (h *Handler) GetReminders(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "未授权访问"})
		return
	}

	reminders, err := h.reminderService.GetUserReminders(userID.(uint))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "获取提醒列表失败"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"reminders": reminders,
	})
}

// CreateReminder 创建新的提醒
func (h *Handler) CreateReminder(c *gin.Context) {
	var reminder model.Reminder
	if err := c.ShouldBindJSON(&reminder); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "无效的请求数据",
			"details": err.Error(),
		})
		return
	}

	userID, _ := c.Get("userID")
	reminder.UserID = userID.(uint)

	if err := h.reminderService.CreateReminder(&reminder); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "创建提醒失败"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message":  "提醒创建成功",
		"reminder": reminder,
	})
}

// UpdateReminder 更新提醒
func (h *Handler) UpdateReminder(c *gin.Context) {
	reminderID := c.Param("id")
	var reminder model.Reminder
	if err := c.ShouldBindJSON(&reminder); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "无效的请求数据",
			"details": err.Error(),
		})
		return
	}

	userID, _ := c.Get("userID")
	reminder.UserID = userID.(uint)

	if err := h.reminderService.UpdateReminder(reminderID, &reminder); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "更新提醒失败"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":  "提醒更新成功",
		"reminder": reminder,
	})
}

// DeleteReminder 删除提醒
func (h *Handler) DeleteReminder(c *gin.Context) {
	reminderID := c.Param("id")
	userID, _ := c.Get("userID")

	if err := h.reminderService.DeleteReminder(userID.(uint), reminderID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "删除提醒失败"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "提醒删除成功"})
}