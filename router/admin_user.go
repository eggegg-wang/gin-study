package router

import (
	"FuckingVersion1/api/v1"
	"FuckingVersion1/middleware"
	"github.com/gin-gonic/gin"
)

func InitUserRouter(Router *gin.RouterGroup) {
	UserRouter := Router.Group("user").Use(middleware.JWTAuth()).Use(middleware.CasbinHandler())
	//.Use(middleware.JWTAuth()).Use(middleware.CasbinHandler())
	{
		UserRouter.POST("changePassword", v1.ChangePassword)     // 修改密码
		UserRouter.POST("uploadHeaderImg", v1.UploadHeaderImg)	//上传头像-七牛云
		UserRouter.POST("UploadImage", v1.UploadImage) 			//上传头像-本地存储
		UserRouter.POST("getUserList", v1.GetUserList)           // 分页获取用户列表
		UserRouter.POST("setUserAuthority", v1.SetUserAuthority) //设置用户权限
	}
}