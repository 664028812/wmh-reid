package handler

import (
	"net/http"
	"wmh/internal/model"

	"github.com/gin-gonic/gin"
)

// GetUserProfile 获取用户个人资料
func (h *Handler) GetUserProfile(c *gin.Context) {
	userID, _ := c.Get("userID")
	user, err := h.userService.GetUserByID(userID.(uint))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "用户不存在"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"profile": user})
}

// UpdateUserProfile 更新用户个人资料
func (h *Handler) UpdateUserProfile(c *gin.Context) {
	var profile model.User
	if err := c.ShouldBindJSON(&profile); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的请求数据"})
		return
	}

	userID, _ := c.Get("userID")
	profile.ID = userID.(uint)

	if err := h.userService.UpdateUser(&profile); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "更新个人资料失败"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "个人资料更新成功"})
}