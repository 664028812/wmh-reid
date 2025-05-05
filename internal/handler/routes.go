package handler

import (
	"wmh/internal/middleware"

	"github.com/gin-gonic/gin"
)

// SetupRoutes 设置路由
func (h *Handler) SetupRoutes() *gin.Engine {
	r := gin.Default()

	// API 版本组
	v1 := r.Group("/api/v1")

	// 公开路由组
	public := v1.Group("")
	{
		// 用户认证相关路由
		public.POST("/register", h.Register)
		public.POST("/login", h.Login)
		public.POST("/refresh-token", h.RefreshToken) // 添加刷新token的路由
	}

	// 需要认证的路由组
	authorized := v1.Group("")
	authorized.Use(middleware.AuthMiddleware())
	{
		// 目标管理
		authorized.POST("/goals", h.CreateGoal)
		authorized.GET("/goals", h.GetGoals)
		authorized.GET("/goals/:id", h.GetGoal)
		authorized.PUT("/goals/:id", h.UpdateGoal)
		authorized.DELETE("/goals/:id", h.DeleteGoal)

		// 进度管理
		authorized.POST("/progress", h.UpdateProgress)
		authorized.GET("/progress", h.GetProgress)
		authorized.GET("/progress/stats", h.GetProgressStats)

		// 提醒管理
		authorized.GET("/reminders", h.GetReminders)
		authorized.POST("/reminders", h.CreateReminder)
		authorized.PUT("/reminders/:id", h.UpdateReminder)
		authorized.DELETE("/reminders/:id", h.DeleteReminder)

		// 用户相关
		authorized.GET("/user/profile", h.GetUserProfile)
		authorized.PUT("/user/profile", h.UpdateUserProfile)
	}

	// 管理员路由组
	admin := v1.Group("/admin")
	admin.Use(middleware.AuthMiddleware(), middleware.AdminOnly())
	{
		// 用户管理
		admin.GET("/users", h.ListUsers)
		admin.GET("/users/:id", h.GetUser)
		admin.PUT("/users/:id", h.UpdateUser)
		admin.DELETE("/users/:id", h.DeleteUser)

		// 系统管理
		admin.GET("/stats", h.GetSystemStats)
	}

	return r
}