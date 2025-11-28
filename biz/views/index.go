package views

import (
	"github.com/LikkasT/GO-BLOG/biz/util"
	"github.com/gin-gonic/gin"
)

type ViewsHandler struct {
	*util.AppDependencies
}

func NewViewsHandler(deps *util.AppDependencies) *ViewsHandler {
	return &ViewsHandler{
		AppDependencies: deps,
	}
}

func (v *ViewsHandler) Index(c *gin.Context) {
	v.Logger.Info("Index view called")
	c.JSON(200, gin.H{
		"message": "pong",
	})
}
