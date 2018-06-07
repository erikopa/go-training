package util

import (
	"github.com/sirupsen/logrus"
	"os"
)

var logger *logrus.Logger

func initLogger(){
	logger = logrus.New()
	logger.Formatter = &logrus.JSONFormatter{}
}

// GetLogger ...
func GetLogger() *logrus.Entry{
	if logger == nil{
		initLogger()
	}

	hostname, _ := os.Hostname()

	return logger.WithFields(map[string]interface{}{
		"hostname":hostname,
	})
}