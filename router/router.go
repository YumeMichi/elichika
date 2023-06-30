package router

import (
	"elichika/handler"
	"elichika/middleware"

	"github.com/gin-gonic/gin"
)

func Router(r *gin.Engine) {
	r.Static("/ep3110/static", "static")
	api := r.Group("ep3110").Use(middleware.Common)
	{
		api.POST("/login/startup", handler.StartUp)
		api.POST("/login/login", handler.Login)
		api.POST("/asset/getPackUrl", handler.GetPackUrl)
		api.POST("/bootstrap/fetchBootstrap", handler.FetchBootstrap)
		api.POST("/notice/fetchNotice", handler.FetchNotice)
		api.POST("/navi/saveUserNaviVoice", handler.SaveUserNaviVoice)
		api.POST("/card/updateCardNewFlag", handler.UpdateCardNewFlag)
		api.POST("/bootstrap/getClearedPlatformAchievement", handler.GetClearedPlatformAchievement)
		api.POST("/gameSettings/updatePushNotificationSettings", handler.UpdatePushNotificationSettings)
		api.POST("/userProfile/fetchProfile", handler.FetchProfile)
	}
}
