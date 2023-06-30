package model

// UserCommunicationMemberDetailBadgeByID ...
type UserCommunicationMemberDetailBadgeByID struct {
	MemberMasterID     int  `json:"member_master_id"`
	IsStoryMemberBadge bool `json:"is_story_member_badge"`
	IsStorySideBadge   bool `json:"is_story_side_badge"`
	IsVoiceBadge       bool `json:"is_voice_badge"`
	IsThemeBadge       bool `json:"is_theme_badge"`
	IsCardBadge        bool `json:"is_card_badge"`
	IsMusicBadge       bool `json:"is_music_badge"`
}

// UserMemberInfo ...
type UserMemberInfo struct {
	MemberMasterID           int  `json:"member_master_id"`
	CustomBackgroundMasterID int  `json:"custom_background_master_id"`
	SuitMasterID             int  `json:"suit_master_id"`
	LovePoint                int  `json:"love_point"`
	LovePointLimit           int  `json:"love_point_limit"`
	LoveLevel                int  `json:"love_level"`
	ViewStatus               int  `json:"view_status"`
	IsNew                    bool `json:"is_new"`
}

// SuitInfo ...
type SuitInfo struct {
	SuitMasterID int  `json:"suit_master_id"`
	IsNew        bool `json:"is_new"`
}
