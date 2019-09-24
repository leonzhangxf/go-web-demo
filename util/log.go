package util

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"io"
	"os"
	"time"
)

// 日志配置
var Log *logrus.Logger

var Out io.Writer

func init() {
	Out = os.Stdout

	Log = logrus.New()
	Log.Formatter = new(logrus.JSONFormatter)
	Log.Level = logrus.TraceLevel
	Log.Out = Out
}

func GetGinLogrusConfig() *gin.LoggerConfig {
	return &gin.LoggerConfig{
		Formatter: GinLogrusFormatterAdapter,
		Output:    Out,
		SkipPaths: nil,
	}
}

func GinLogrusFormatterAdapter(param gin.LogFormatterParams) string {
	if !logrus.IsLevelEnabled(logrus.DebugLevel) {
		return ""
	}

	if param.Latency > time.Minute {
		// Truncate in a golang < 1.8 safe way
		param.Latency = param.Latency - param.Latency%time.Second
	}

	entry := logrus.NewEntry(Log)
	entry.Level = logrus.DebugLevel
	entry.Time = param.TimeStamp
	entry.Message = fmt.Sprintf(
		"GIN status: %3d | latency: %13v | client ip: %15s | method: %s | path: %-7s "+
			" errorMessage: %s \n",
		param.StatusCode,
		param.Latency,
		param.ClientIP,
		param.Method,
		param.Path,
		param.ErrorMessage,
	)
	bytes, err := Log.Formatter.Format(entry)
	if nil != err {
		return ""
	}
	return string(bytes)
}
