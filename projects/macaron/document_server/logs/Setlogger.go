package logs

import "github.com/wonderivan/logger"

func SetLogger(filepath string) {
	// 配置logger
	logger.SetLogger(filepath)
}
