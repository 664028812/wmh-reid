package database

import (
	"gorm.io/gorm"
)

// Transaction 执行事务
func Transaction(fn func(tx *gorm.DB) error) error {
	return DB.Transaction(fn)
}

// Create 创建记录
func Create(value interface{}) error {
	return DB.Create(value).Error
}

// Update 更新记录
func Update(value interface{}) error {
	return DB.Save(value).Error
}

// Delete 删除记录
func Delete(value interface{}) error {
	return DB.Delete(value).Error
}

// First 查找第一条记录
func First(dest interface{}, conds ...interface{}) error {
	return DB.First(dest, conds...).Error
}

// Find 查找多条记录
func Find(dest interface{}, conds ...interface{}) error {
	return DB.Find(dest, conds...).Error
}