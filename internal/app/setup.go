package core

import (
	"mrKrabsmr/web-chat/internal/configs"
	"os"
	"sync"

	"github.com/sirupsen/logrus"
)

var (
	config     *configs.Config
	configOnce = new(sync.Once)

	logger     *logrus.Logger
	loggerOnce = new(sync.Once)
)

func GetConfig() *configs.Config {
	configOnce.Do(func() {
		config = &configs.Config{
			Address:  os.Getenv("ADDRESS"),
			LogLevel: os.Getenv("LOG_LEVEL"),
		}
	})

	return config
}

func GetLogger() *logrus.Logger {
	loggerOnce.Do(func ()  {
		logLevel, err := logrus.ParseLevel(GetConfig().LogLevel)
		if err != nil {
			panic(err)
		}

		logger = logrus.New()
		logger.SetLevel(logLevel)
	})

	return logger
}
