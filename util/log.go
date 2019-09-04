package util

import (
	"github.com/sirupsen/logrus"
	"os"
)

// 日志配置
var Log *logrus.Logger

func init() {
	Log = logrus.New()
	Log.Formatter = new(logrus.JSONFormatter)
	Log.Level = logrus.TraceLevel
	Log.Out = os.Stdout
}
