package main

import (
	"log"
	"wmh/config"
	"wmh/internal/handler"
	"wmh/internal/service"
	"wmh/pkg/database"
)

func main() {
	// 加载配置
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("加载配置失败: %v", err)
	}

	// 初始化数据库连接
	if err := database.InitDB(cfg.Database); err != nil {
		log.Fatalf("初始化数据库失败: %v", err)
	}
	defer database.CloseDB()

	// 初始化服务
	userService := service.NewUserService()
	goalService := service.NewGoalService()
	reminderService := service.NewReminderService()

	// 初始化处理器
	h := handler.NewHandler(userService, goalService, reminderService)

	// 设置路由
	r := h.SetupRoutes()

	// 启动服务器
	log.Printf("服务器正在启动，监听地址: %s", cfg.Server.Address)
	if err := r.Run(cfg.Server.Address); err != nil {
		log.Fatalf("服务器启动失败: %v", err)
	}
}
