package handler

import (
	"net/http"
	"time"
	"wmh/internal/model"

	"github.com/gin-gonic/gin"
)

// CreateGoal 处理创建目标请求
func (h *Handler) CreateGoal(c *gin.Context) {
	var goal model.Goal
	if err := c.ShouldBindJSON(&goal); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "无效的请求数据",
			"details": err.Error(),
		})
		return
	}

	// 从上下文获取用户ID
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "未授权访问"})
		return
	}
	goal.UserID = userID.(uint)

	// 创建目标
	if err := h.goalService.CreateGoal(goal.UserID, &goal); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "创建目标失败"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "目标创建成功",
		"goal":    goal,
	})
}

// UpdateProgress 处理更新进度请求
func (h *Handler) UpdateProgress(c *gin.Context) {
	var progress model.Progress
	if err := c.ShouldBindJSON(&progress); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "无效的请求数据",
			"details": err.Error(),
		})
		return
	}

	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "未授权访问"})
		return
	}
	progress.UserID = userID.(uint)
	progress.RecordDate = time.Now()

	if err := h.goalService.UpdateProgress(progress.UserID, &progress); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "更新进度失败"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":  "进度更新成功",
		"progress": progress,
	})
}

// GetProgress 获取进度统计
func (h *Handler) GetProgress(c *gin.Context) {
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

	progress, err := h.goalService.GetProgressStats(userID.(uint), goalID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "获取进度统计失败"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"progress_stats": progress,
	})
}
