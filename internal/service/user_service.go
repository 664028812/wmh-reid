package service

import (
	"fmt"
	"strconv"
	"wmh/internal/model"
	"wmh/pkg/database"
)

type UserService struct {
	// 可以添加其他依赖
}

// NewUserService 创建用户服务实例
func NewUserService() *UserService {
	return &UserService{}
}

// CreateUser 创建新用户
func (s *UserService) CreateUser(user *model.User) error {
	if err := database.DB.Create(user).Error; err != nil {
		return fmt.Errorf("创建用户失败: %v", err)
	}
	return nil
}

// GetUserByUsername 通过用户名获取用户
func (s *UserService) GetUserByUsername(username string) (*model.User, error) {
	var user model.User
	if err := database.DB.Where("username = ?", username).First(&user).Error; err != nil {
		return nil, fmt.Errorf("用户不存在: %v", err)
	}
	return &user, nil
}

// GetUserByID 通过ID获取用户
func (s *UserService) GetUserByID(userID uint) (*model.User, error) {
	var user model.User
	if err := database.DB.First(&user, userID).Error; err != nil {
		return nil, fmt.Errorf("用户不存在: %v", err)
	}
	return &user, nil
}

// GetUserByIDString 通过字符串ID获取用户
func (s *UserService) GetUserByIDString(userID string) (*model.User, error) {
	id, err := strconv.ParseUint(userID, 10, 64)
	if err != nil {
		return nil, fmt.Errorf("无效的用户ID: %v", err)
	}
	return s.GetUserByID(uint(id))
}

// UpdateUser 更新用户信息
func (s *UserService) UpdateUser(user *model.User) error {
	if err := database.DB.Save(user).Error; err != nil {
		return fmt.Errorf("更新用户失败: %v", err)
	}
	return nil
}

// UpdateUserByID 通过ID更新用户信息
func (s *UserService) UpdateUserByID(userID string, user *model.User) error {
	id, err := strconv.ParseUint(userID, 10, 64)
	if err != nil {
		return fmt.Errorf("无效的用户ID: %v", err)
	}
	
	// 检查用户是否存在
	var existingUser model.User
	if err := database.DB.First(&existingUser, id).Error; err != nil {
		return fmt.Errorf("用户不存在: %v", err)
	}
	
	user.ID = uint(id)
	return s.UpdateUser(user)
}

// DeleteUser 删除用户
func (s *UserService) DeleteUser(userID string) error {
	id, err := strconv.ParseUint(userID, 10, 64)
	if err != nil {
		return fmt.Errorf("无效的用户ID: %v", err)
	}
	
	if err := database.DB.Delete(&model.User{}, id).Error; err != nil {
		return fmt.Errorf("删除用户失败: %v", err)
	}
	return nil
}

// GetAllUsers 获取所有用户
func (s *UserService) GetAllUsers() ([]model.User, error) {
	var users []model.User
	if err := database.DB.Find(&users).Error; err != nil {
		return nil, fmt.Errorf("获取用户列表失败: %v", err)
	}
	return users, nil
}

// GetSystemStats 获取系统统计信息
func (s *UserService) GetSystemStats() (map[string]interface{}, error) {
	var userCount int64
	if err := database.DB.Model(&model.User{}).Count(&userCount).Error; err != nil {
		return nil, fmt.Errorf("获取用户数量失败: %v", err)
	}

	var goalCount int64
	if err := database.DB.Model(&model.Goal{}).Count(&goalCount).Error; err != nil {
		return nil, fmt.Errorf("获取目标数量失败: %v", err)
	}

	stats := map[string]interface{}{
		"total_users": userCount,
		"total_goals": goalCount,
	}

	return stats, nil
}