package handler

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

// GetProgressStats 获取进度统计信息
func (h *Handler) GetProgressStats(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "未授权访问"})
		return
	}

	goalID := c.Query("goal_id")
	if goalID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "缺少目标ID参数"})
		return
	}

	stats, err := h.goalService.GetProgressStats(userID.(uint), goalID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "获取进度统计失败"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"stats": stats,
	})
}