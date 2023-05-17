package logger

import (
	syslog_org "log/syslog"
	"shop/pkg/logger/internal/file"
	"shop/pkg/logger/internal/syslog"
)

// Logger defines custom logger
type Logger interface {
	Config()
	Log(message string) bool
}

// Factory creates a new logger
func Factory(lgType string) Logger {
	switch lgType {
	case "syslog":
		logManager := syslog.Create(syslog_org.LOG_ERR)
		logManager.Config()
		return logManager
	default:
		logFile := file.Create()
		logFile.Config()
		return logFile
	}
}
