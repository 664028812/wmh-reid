package logger

import (
	"fmt"
	"path/filepath"
	"runtime"
	"runtime/debug"

	"github.com/natefinch/lumberjack"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

/*	初始化zap日志
 */
func init() {
	InitZapLogger()
}

var sugarLogger *zap.SugaredLogger
var logDir string

/*	zap日志初始化实现，日志写入文件夹中
 */
func InitZapLogger() {
	logDir = filepath.Join("./temp", "log")
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder //指定时间格式
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	encoder := zapcore.NewConsoleEncoder(encoderConfig)
	var ws zapcore.WriteSyncer
	var logLevel zapcore.Level
	//文件writeSyncer
	fileWriteSyncer := zapcore.AddSync(&lumberjack.Logger{
		Filename:   filepath.Join(logDir, "iot-sdk.log"), //日志文件存放目录
		MaxSize:    MaxSize,                              //文件大小限制,单位MB
		MaxBackups: MaxBackups,                           //最大保留日志文件数量
		MaxAge:     MaxAge,                               //日志文件保留天数
		LocalTime:  true,
		Compress:   true, //是否压缩处理
	})
	ws = zapcore.NewMultiWriteSyncer(fileWriteSyncer)
	logLevel = zapcore.DebugLevel
	// if cfg.BA {
	// 	ws = zapcore.NewMultiWriteSyncer(fileWriteSyncer)
	// 	logLevel = zapcore.InfoLevel
	// } else {
	// 	ws = zapcore.NewMultiWriteSyncer(fileWriteSyncer, zapcore.AddSync(os.Stdout))
	// 	logLevel = zapcore.DebugLevel
	// }
	fileCore := zapcore.NewCore(encoder, ws, logLevel)
	//logger := zap.New(fileCore, zap.AddCaller(), zap.AddCallerSkip(1), zap.WithClock(zapcore.SystemClock{}))
	logger := zap.New(fileCore,
		zap.AddCaller(),                       // 添加调用者信息
		zap.AddCallerSkip(1),                  // 跳过一层调用栈，以显示实际的调用位置
		zap.AddStacktrace(zapcore.ErrorLevel), // 为错误级别及以上添加堆栈跟踪
		zap.Development(),                     // 开发模式，更详细的日志信息
	)
	sugarLogger = logger.Sugar()
}

/*
写入日志
@param args 参数
@param args 参数
*/
func Info(args ...interface{}) {
	msg := WhetherEncryp("", args)
	sugarLogger.Info(msg)
}

/*
写入带格式日志
@param temp 格式
@param args 参数
*/
func Infof(temp string, args ...interface{}) {
	msg := WhetherEncryp(temp, args)
	sugarLogger.Infof(msg)
}

/*
Debug日志
@param args 参数
@param args 参数
*/
func Debug(args ...interface{}) {
	msg := WhetherEncryp("", args)
	sugarLogger.Debug(msg)
}

/*
写入带格式日志
@param temp 格式
@param args 参数
*/
func Debugf(temp string, args ...interface{}) {
	msg := WhetherEncryp(temp, args)
	sugarLogger.Debugf(msg)
}

/*
错误日志
@param args 参数
*/
func Error(args ...interface{}) {
	msg := WhetherEncryp("", args)
	sugarLogger.Error(msg)
}

/*
带格式错误日志
@param temp 格式
@param args 参数
*/
func Errorf(temp string, args ...interface{}) {
	hasRuntimeErr := false
	for _, arg := range args {
		// 判断当前的error日志传递参数是否有 runtime error 有才需要追加 错误堆栈信息
		if _, ok := arg.(runtime.Error); ok {
			hasRuntimeErr = true
			break
		}
	}
	errStack := debug.Stack()
	if len(errStack) != 0 && hasRuntimeErr && args != nil {
		args = append(args, string(errStack))
		temp += " err_stack: %s"
	}
	msg := WhetherEncryp(temp, args)
	sugarLogger.Errorf(msg)
}

/*
告警日志
@param args 参数
*/
func Warning(args ...interface{}) {
	msg := WhetherEncryp("", args)
	sugarLogger.Warn(msg)
}

/*
带格式告警日志
@param temp 格式
@param args 参数
*/
func Warningf(temp string, args ...interface{}) {
	msg := WhetherEncryp(temp, args)
	sugarLogger.Warnf(msg)
}

// Panicf
func Panicf(temp string, args ...interface{}) {
	msg := WhetherEncryp(temp, args)
	sugarLogger.DPanicf(msg)
}

func getMessage(template string, fmtArgs []interface{}) string {
	if len(fmtArgs) == 0 {
		return template
	}

	if template != "" {
		return fmt.Sprintf(template, fmtArgs...)
	}
	if len(fmtArgs) == 1 {
		if str, ok := fmtArgs[0].(string); ok {
			return str
		}
	}
	return fmt.Sprint(fmtArgs...)
}

func WhetherEncryp(temp string, args []interface{}) string {
	// msg := getMessage(temp, args)
	// if contants.IsProductEnv() {
	// 	msgEncyp := encrypt.GfAesEncrypt([]byte(msg), []byte(contants.LogKey))
	// 	return fmt.Sprintf("start{%s}end", msgEncyp)
	// } else {
	// 	return msg
	// }
	msg := getMessage(temp, args)
	return msg
}
