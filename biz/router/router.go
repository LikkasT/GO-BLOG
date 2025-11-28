package router

import (
	"github.com/LikkasT/GO-BLOG/biz/middle_ware"
	"github.com/LikkasT/GO-BLOG/biz/util"
	"github.com/LikkasT/GO-BLOG/biz/views"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

type Routerdependencies struct {
	*util.AppDependencies
	viewsHandler *views.ViewsHandler
}

func RouterDependencies(config *viper.Viper, logger *logrus.Logger, db *gorm.DB) *Routerdependencies {
	appDeps := util.NewAppDependencies(config, logger, db)
	viewsHandler := views.NewViewsHandler(appDeps)
	return &Routerdependencies{
		AppDependencies: appDeps,
		viewsHandler:    viewsHandler,
	}
}

func (r *Routerdependencies) RouterInit() {
	r.Logger.Info("开始初始化router")
	router := gin.Default()
	router.Use(middle_ware.CORS)
	router.GET("/index", r.viewsHandler.Index)
	apiGroup := router.Group("/api")
	authGroup := apiGroup.Group("/auth")
	authGroup.POST("/register", r.viewsHandler.User_register)
	url := r.Config.GetString("router_server")
	r.Logger.Infof("router将运行在%s", url)
	router.Run(url)
}
