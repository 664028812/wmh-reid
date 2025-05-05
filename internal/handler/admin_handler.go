package handler

import (
	"net/http"
	"wmh/internal/model"

	"github.com/gin-gonic/gin"
)

// ListUsers 获取所有用户列表
func (h *Handler) ListUsers(c *gin.Context) {
	users, err := h.userService.GetAllUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "获取用户列表失败"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"users": users})
}

// GetUser 获取单个用户信息
func (h *Handler) GetUser(c *gin.Context) {
	userID := c.Param("id")
	user, err := h.userService.GetUserByIDString(userID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "用户不存在"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"user": user})
}

// UpdateUser 更新用户信息
func (h *Handler) UpdateUser(c *gin.Context) {
	var user model.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的请求数据"})
		return
	}

	userID := c.Param("id")
	if err := h.userService.UpdateUserByID(userID, &user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "更新用户失败"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "用户更新成功"})
}

// DeleteUser 删除用户
func (h *Handler) DeleteUser(c *gin.Context) {
	userID := c.Param("id")
	if err := h.userService.DeleteUser(userID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "删除用户失败"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "用户删除成功"})
}

// GetSystemStats 获取系统统计信息
func (h *Handler) GetSystemStats(c *gin.Context) {
	stats, err := h.userService.GetSystemStats()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "获取系统统计失败"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"stats": stats})
}