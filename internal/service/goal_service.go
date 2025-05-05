package service

import (
	"fmt"
	"strconv"
	"wmh/internal/model"
	"wmh/pkg/database"
)

type GoalService struct {
	// 可以添加其他依赖
}

// NewGoalService 创建目标服务实例
func NewGoalService() *GoalService {
	return &GoalService{}
}

// GetProgressStats 获取进度统计信息
func (s *GoalService) GetProgressStats(userID uint, goalIDStr string) (map[string]interface{}, error) {
	// 将字符串ID转换为uint
	goalID, err := strconv.ParseUint(goalIDStr, 10, 64)
	if err != nil {
		return nil, fmt.Errorf("无效的目标ID: %v", err)
	}

	// 获取目标信息
	var goal model.Goal
	if err := database.DB.First(&goal, goalID).Error; err != nil {
		return nil, fmt.Errorf("获取目标信息失败: %v", err)
	}

	// 验证目标是否属于该用户
	if goal.UserID != userID {
		return nil, fmt.Errorf("无权访问此目标")
	}

	// 获取该目标的所有进度记录
	var progresses []model.Progress
	if err := database.DB.Where("goal_id = ?", goalID).Order("record_date desc").Find(&progresses).Error; err != nil {
		return nil, fmt.Errorf("获取进度记录失败: %v", err)
	}

	// 计算统计数据
	stats := calculateStats(goal, progresses)

	return stats, nil
}

// calculateStats 计算统计数据
func calculateStats(goal model.Goal, progresses []model.Progress) map[string]interface{} {
	stats := make(map[string]interface{})
	
	// 基本信息
	stats["goal_type"] = goal.Type
	stats["target"] = goal.Target
	stats["start_value"] = goal.StartValue
	stats["current_value"] = goal.CurrentValue
	
	// 计算进度百分比
	if goal.Target != goal.StartValue {
		progress := (goal.CurrentValue - goal.StartValue) / (goal.Target - goal.StartValue) * 100
		stats["progress_percentage"] = progress
	} else {
		stats["progress_percentage"] = 100
	}

	// 最近记录
	if len(progresses) > 0 {
		stats["latest_record"] = progresses[0]
		
		// 计算趋势（最近7条记录）
		trend := calculateTrend(progresses)
		stats["trend"] = trend
	}

	// 计算距离目标还需要的量
	if goal.Type == "weight_loss" {
		stats["remaining"] = goal.CurrentValue - goal.Target
	} else if goal.Type == "study" {
		stats["remaining"] = goal.Target - goal.CurrentValue
	}

	return stats
}

// calculateTrend 计算趋势
func calculateTrend(progresses []model.Progress) string {
	if len(progresses) < 2 {
		return "stable"
	}

	// 获取最近两次记录
	latest := progresses[0].Value
	previous := progresses[1].Value

	if latest < previous {
		return "decreasing"
	} else if latest > previous {
		return "increasing"
	}
	return "stable"
}

// GetUserGoals 获取用户的所有目标
func (s *GoalService) GetUserGoals(userID uint) ([]model.Goal, error) {
	var goals []model.Goal
	if err := database.DB.Where("user_id = ?", userID).Find(&goals).Error; err != nil {
		return nil, fmt.Errorf("获取目标列表失败: %v", err)
	}
	return goals, nil
}

// GetGoal 获取单个目标
func (s *GoalService) GetGoal(userID uint, goalIDStr string) (*model.Goal, error) {
	goalID, err := strconv.ParseUint(goalIDStr, 10, 64)
	if err != nil {
		return nil, fmt.Errorf("无效的目标ID: %v", err)
	}

	var goal model.Goal
	if err := database.DB.First(&goal, goalID).Error; err != nil {
		return nil, fmt.Errorf("目标不存在: %v", err)
	}

	// 验证目标是否属于该用户
	if goal.UserID != userID {
		return nil, fmt.Errorf("无权访问此目标")
	}

	return &goal, nil
}

// UpdateGoal 更新目标
func (s *GoalService) UpdateGoal(goalIDStr string, goal *model.Goal) error {
	goalID, err := strconv.ParseUint(goalIDStr, 10, 64)
	if err != nil {
		return fmt.Errorf("无效的目标ID: %v", err)
	}

	// 检查目标是否存在
	var existingGoal model.Goal
	if err := database.DB.First(&existingGoal, goalID).Error; err != nil {
		return fmt.Errorf("目标不存在: %v", err)
	}

	// 验证目标是否属于该用户
	if existingGoal.UserID != goal.UserID {
		return fmt.Errorf("无权访问此目标")
	}

	// 更新目标
	goal.ID = uint(goalID)
	if err := database.DB.Save(goal).Error; err != nil {
		return fmt.Errorf("更新目标失败: %v", err)
	}

	return nil
}

// DeleteGoal 删除目标
func (s *GoalService) DeleteGoal(userID uint, goalIDStr string) error {
	goalID, err := strconv.ParseUint(goalIDStr, 10, 64)
	if err != nil {
		return fmt.Errorf("无效的目标ID: %v", err)
	}

	// 检查目标是否存在
	var goal model.Goal
	if err := database.DB.First(&goal, goalID).Error; err != nil {
		return fmt.Errorf("目标不存在: %v", err)
	}

	// 验证目标是否属于该用户
	if goal.UserID != userID {
		return fmt.Errorf("无权访问此目标")
	}

	// 删除目标
	if err := database.DB.Delete(&goal).Error; err != nil {
		return fmt.Errorf("删除目标失败: %v", err)
	}

	return nil
}

// CreateGoal 创建目标
func (s *GoalService) CreateGoal(userID uint, goal *model.Goal) error {
	goal.UserID = userID
	if err := database.DB.Create(goal).Error; err != nil {
		return fmt.Errorf("创建目标失败: %v", err)
	}
	return nil
}

// UpdateProgress 更新进度
func (s *GoalService) UpdateProgress(userID uint, progress *model.Progress) error {
	// 检查目标是否存在且属于该用户
	var goal model.Goal
	if err := database.DB.First(&goal, progress.GoalID).Error; err != nil {
		return fmt.Errorf("目标不存在: %v", err)
	}

	if goal.UserID != userID {
		return fmt.Errorf("无权访问此目标")
	}

	// 创建进度记录
	progress.UserID = userID
	if err := database.DB.Create(progress).Error; err != nil {
		return fmt.Errorf("创建进度记录失败: %v", err)
	}

	// 更新目标的当前值
	goal.CurrentValue = progress.Value
	if err := database.DB.Save(&goal).Error; err != nil {
		return fmt.Errorf("更新目标当前值失败: %v", err)
	}

	return nil
}

func (s *GoalService) AnalyzeProgress(goalID uint) (interface{}, error) {
	// 分析进度，返回统计数据
	return nil, nil
}
