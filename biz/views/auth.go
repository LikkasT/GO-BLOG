package views

import (
	"github.com/LikkasT/GO-BLOG/biz/models"
	"github.com/gin-gonic/gin"
)

type CreateForm struct {
	Username string `form:"username" binding:"max=20"`
	Password string `form:"password" binding:"max=20"`
}

func (v *ViewsHandler) User_register(c *gin.Context) {
	v.Logger.Info("User_register view called")
	var form CreateForm
	if err := c.ShouldBind(&form); err != nil {
		c.JSON(422, gin.H{
			"message": "Form binding error",
		})
	}
	info_username := form.Username
	info_password := form.Password
	authorization_create := models.Users{Username: info_username, Password: info_password, Level: 0}
	result := v.DB.Create(&authorization_create)
	if result.Error != nil {
		c.JSON(422, gin.H{
			"message": "数据库数据创建失败（数据库出错或者是传参与数据库定义不匹配）",
		})
	}
	c.JSON(200, gin.H{
		"message": "用户创建成功",
	})
}
