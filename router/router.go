package router

import (
	"elichika/handler"
	"elichika/middleware"

	"github.com/gin-gonic/gin"
)

func Router(r *gin.Engine) {
	r.Static("/ep3120/static", "static")
	r.Static("/static", "static")
	api := r.Group("ep3120").Use(middleware.Common)
	{
		api.POST("/asset/getPackUrl", handler.GetPackUrl)
		api.POST("/bootstrap/fetchBootstrap", handler.FetchBootstrap)
		api.POST("/bootstrap/getClearedPlatformAchievement", handler.GetClearedPlatformAchievement)
		api.POST("/card/changeIsAwakeningImage", handler.ChangeIsAwakeningImage)
		api.POST("/card/updateCardNewFlag", handler.UpdateCardNewFlag)
		api.POST("/communicationMember/fetchCommunicationMemberDetail", handler.FetchCommunicationMemberDetail)
		api.POST("/communicationMember/finishUserStoryMember", handler.FinishUserStoryMember)
		api.POST("/communicationMember/finishUserStorySide", handler.FinishUserStorySide)
		api.POST("/communicationMember/setTheme", handler.SetTheme)
		api.POST("/communicationMember/setFavoriteMember", handler.SetFavoriteMember)
		api.POST("/communicationMember/updateUserCommunicationMemberDetailBadge", handler.UpdateUserCommunicationMemberDetailBadge)
		api.POST("/communicationMember/updateUserLiveDifficultyNewFlag", handler.UpdateUserLiveDifficultyNewFlag)
		api.POST("/emblem/activateEmblem", handler.ActivateEmblem)
		api.POST("/emblem/fetchEmblem", handler.FetchEmblem)
		api.POST("/gameSettings/updatePushNotificationSettings", handler.UpdatePushNotificationSettings)
		api.POST("/lesson/executeLesson", handler.ExecuteLesson)
		api.POST("/lesson/resultLesson", handler.ResultLesson)
		api.POST("/lesson/saveDeck", handler.SaveDeckLesson)
		api.POST("/lesson/skillEditResult", handler.SkillEditResult)
		api.POST("/liveDeck/fetchLiveDeckSelect", handler.FetchLiveDeckSelect)
		api.POST("/liveDeck/saveDeckAll", handler.SaveDeckAll)
		api.POST("/liveDeck/saveDeck", handler.SaveDeck)
		api.POST("/liveDeck/saveSuit", handler.SaveSuit)
		api.POST("/live/fetchLiveMusicSelect", handler.FetchLiveMusicSelect)
		api.POST("/live/finish", handler.LiveFinish)
		api.POST("/liveMv/saveDeck", handler.LiveMvSaveDeck)
		api.POST("/liveMv/start", handler.LiveMvStart)
		api.POST("/livePartners/fetch", handler.FetchLivePartners)
		api.POST("/live/start", handler.LiveStart)
		api.POST("/login/login", handler.Login)
		api.POST("/login/startup", handler.StartUp)
		api.POST("/mission/clearMissionBadge", handler.ClearMissionBadge)
		api.POST("/mission/fetchMission", handler.FetchMission)
		api.POST("/navi/saveUserNaviVoice", handler.SaveUserNaviVoice)
		api.POST("/notice/fetchNotice", handler.FetchNotice)
		api.POST("/present/fetch", handler.FetchPresent)
		api.POST("/storyEventHistory/finishStory", handler.FinishStory)
		api.POST("/story/finishStoryLinkage", handler.FinishStoryLinkage)
		api.POST("/story/finishUserStoryMain", handler.FinishStoryMain)
		api.POST("/terms/agreement", handler.Agreement)
		api.POST("/userProfile/fetchProfile", handler.FetchProfile)
		api.POST("/userProfile/setProfile", handler.SetProfile)
	}
}
