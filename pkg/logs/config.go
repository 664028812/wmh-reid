//go:build !embed
// +build !embed

package logger

const (
	MaxSize    = 7  // 文件最大大小，单位M
	MaxBackups = 30 //最大保留日志文件数量
	MaxAge     = 5  //文件最大保存天数
)
