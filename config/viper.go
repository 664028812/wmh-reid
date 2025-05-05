package config

import (
	"fmt"
	"github.com/spf13/viper"
)

var globalConfig *Config

// Load 加载配置文件
func Load() (*Config, error) {
	if globalConfig != nil {
		return globalConfig, nil
	}

	viper.SetConfigName("config")  // 配置文件名称
	viper.SetConfigType("yaml")    // 配置文件类型
	viper.AddConfigPath(".")       // 在当前目录查找
	viper.AddConfigPath("./config") // 在config目录查找

	// 设置默认值
	setDefaults()

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// 配置文件不存在时创建默认配置文件
			if err := createDefaultConfig(); err != nil {
				return nil, fmt.Errorf("创建默认配置文件失败: %v", err)
			}
		} else {
			return nil, fmt.Errorf("读取配置文件失败: %v", err)
		}
	}

	config := &Config{}
	if err := viper.Unmarshal(config); err != nil {
		return nil, fmt.Errorf("解析配置文件失败: %v", err)
	}

	globalConfig = config
	return config, nil
}

// setDefaults 设置默认配置
func setDefaults() {
	viper.SetDefault("server.address", ":8080")
	viper.SetDefault("server.mode", "debug")

	viper.SetDefault("database.driver", "sqlite3")
	viper.SetDefault("database.dbfile", "wmh.db")
	viper.SetDefault("database.host", "localhost")
	viper.SetDefault("database.port", 3306)
	viper.SetDefault("database.username", "root")
	viper.SetDefault("database.password", "")
	viper.SetDefault("database.dbname", "wmh")

	viper.SetDefault("redis.host", "localhost")
	viper.SetDefault("redis.port", 6379)
	viper.SetDefault("redis.password", "")
	viper.SetDefault("redis.db", 0)

	viper.SetDefault("jwt.secret", "your-secret-key")
	viper.SetDefault("jwt.expire_time", 24)
}

// createDefaultConfig 创建默认配置文件
func createDefaultConfig() error {
	return viper.SafeWriteConfig()
}