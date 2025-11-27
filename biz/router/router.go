package router

import (
	"github.com/LikkasT/GO-BLOG/biz/middle_ware"
	"github.com/LikkasT/GO-BLOG/biz/views"
	"github.com/gin-gonic/gin"
)

func RouterInit() {
	router := gin.Default()
	router.Use(middle_ware.CORS)
	router.GET("/index", views.Index)
	router.Run()
}
