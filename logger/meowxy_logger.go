package logger

import (
	"encoding/json"

	"go.uber.org/zap"
)

var meowxy_logger *zap.SugaredLogger

func GetMeowxyLogger() *zap.SugaredLogger {
	return meowxy_logger
}

func InitMeowxyLogger() {
	var cfg *zap.Config

	loggerConfig := []byte(`{
		"level": "info",
		"encoding": "console",
		"timekey": "timestamp",
		"outputPaths": ["stdout", "/tmp/logs"],
		"errorOutputPaths": ["stderr"],
		"encoderConfig": {
		  "messageKey": "message",
		  "levelKey": "level",
		  "levelEncoder": "lowercase"
		}
	  }`)

	if err := json.Unmarshal(loggerConfig, &cfg); err != nil {
		panic(err)
	}

	logger, err := zap.NewDevelopment()
	if err != nil {
		panic(err)
	}
	defer logger.Sync()
	meowxy_logger = logger.Sugar()
}
