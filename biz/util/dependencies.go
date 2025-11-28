package util

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

// AppDependencies 应用程序共享依赖结构体
type AppDependencies struct {
	Config *viper.Viper
	Logger *logrus.Logger
	DB     *gorm.DB
}

// NewAppDependencies 创建应用程序依赖实例
func NewAppDependencies(config *viper.Viper, logger *logrus.Logger, db *gorm.DB) *AppDependencies {
	return &AppDependencies{
		Config: config,
		Logger: logger,
		DB:     db,
	}
}
