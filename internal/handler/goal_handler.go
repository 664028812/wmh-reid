package handler

import (
	"net/http"
	"wmh/internal/model"

	"github.com/gin-gonic/gin"
)

// GetGoals 获取用户的所有目标
func (h *Handler) GetGoals(c *gin.Context) {
	userID, _ := c.Get("userID")
	goals, err := h.goalService.GetUserGoals(userID.(uint))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "获取目标列表失败"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"goals": goals})
}

// GetGoal 获取单个目标详情
func (h *Handler) GetGoal(c *gin.Context) {
	goalID := c.Param("id")
	userID, _ := c.Get("userID")

	goal, err := h.goalService.GetGoal(userID.(uint), goalID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "目标不存在"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"goal": goal})
}

// UpdateGoal 更新目标
func (h *Handler) UpdateGoal(c *gin.Context) {
	var goal model.Goal
	if err := c.ShouldBindJSON(&goal); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的请求数据"})
		return
	}

	goalID := c.Param("id")
	userID, _ := c.Get("userID")
	goal.UserID = userID.(uint)

	if err := h.goalService.UpdateGoal(goalID, &goal); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "更新目标失败"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "目标更新成功"})
}

// DeleteGoal 删除目标
func (h *Handler) DeleteGoal(c *gin.Context) {
	goalID := c.Param("id")
	userID, _ := c.Get("userID")

	if err := h.goalService.DeleteGoal(userID.(uint), goalID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "删除目标失败"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "目标删除成功"})
}