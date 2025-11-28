package router

import (
	"github.com/LikkasT/GO-BLOG/biz/middle_ware"
	"github.com/LikkasT/GO-BLOG/biz/views"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

type Routerdependencies struct {
	config *viper.Viper
}

func RouterDependencies(config *viper.Viper) *Routerdependencies {
	return &Routerdependencies{
		config: config,
	}
}
func (r *Routerdependencies) RouterInit() {
	router := gin.Default()
	router.Use(middle_ware.CORS)
	router.GET("/index", views.Index)
	router.Run(r.config.GetString("router_server"))
}
