package log

import "gopkg.in/natefinch/lumberjack.v2"

// 日志切分压缩插件
func logPlugin(filename string) lumberjack.Logger {
	return lumberjack.Logger{
		Filename:   filename,
		MaxSize:    128,  // 每个日志文件保存的最大尺寸 单位：M
		MaxBackups: 30,   // 日志文件最多保存多少个备份
		MaxAge:     7,    // 文件最多保存多少天
		Compress:   true, // 是否压缩
	}
}
