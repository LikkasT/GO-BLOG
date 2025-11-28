package router

import (
	"github.com/LikkasT/GO-BLOG/biz/middle_ware"
	"github.com/LikkasT/GO-BLOG/biz/views"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type Routerdependencies struct {
	config *viper.Viper
	logger *logrus.Logger
}

func RouterDependencies(config *viper.Viper, logger *logrus.Logger) *Routerdependencies {
	return &Routerdependencies{
		config: config,
		logger: logger,
	}
}
func (r *Routerdependencies) RouterInit() {
	r.logger.Info("开始初始化router")
	router := gin.Default()
	router.Use(middle_ware.CORS)
	router.GET("/index", views.Index)
	url := r.config.GetString("router_server")
	r.logger.Infof("router将运行在%s", url)
	router.Run(url)
}
