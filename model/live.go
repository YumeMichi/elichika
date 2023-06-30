package model

// SaveDeckReq ...
type SaveDeckReq struct {
	DeckID       int   `json:"deck_id"`
	CardWithSuit []int `json:"card_with_suit"`
	SquadDict    []any `json:"squad_dict"`
}

// DeckInfo ...
type DeckInfo struct {
	UserLiveDeckID int      `json:"user_live_deck_id"`
	Name           DeckName `json:"name"`
	CardMasterID1  int      `json:"card_master_id_1"`
	CardMasterID2  int      `json:"card_master_id_2"`
	CardMasterID3  int      `json:"card_master_id_3"`
	CardMasterID4  int      `json:"card_master_id_4"`
	CardMasterID5  int      `json:"card_master_id_5"`
	CardMasterID6  int      `json:"card_master_id_6"`
	CardMasterID7  int      `json:"card_master_id_7"`
	CardMasterID8  int      `json:"card_master_id_8"`
	CardMasterID9  int      `json:"card_master_id_9"`
	SuitMasterID1  int      `json:"suit_master_id_1"`
	SuitMasterID2  int      `json:"suit_master_id_2"`
	SuitMasterID3  int      `json:"suit_master_id_3"`
	SuitMasterID4  int      `json:"suit_master_id_4"`
	SuitMasterID5  int      `json:"suit_master_id_5"`
	SuitMasterID6  int      `json:"suit_master_id_6"`
	SuitMasterID7  int      `json:"suit_master_id_7"`
	SuitMasterID8  int      `json:"suit_master_id_8"`
	SuitMasterID9  int      `json:"suit_master_id_9"`
}

// DeckName ...
type DeckName struct {
	DotUnderText string `json:"dot_under_text"`
}

// PartyInfo ...
type PartyInfo struct {
	PartyID          int       `json:"party_id"`
	UserLiveDeckID   int       `json:"user_live_deck_id"`
	Name             PartyName `json:"name"`
	IconMasterID     int       `json:"icon_master_id"`
	CardMasterID1    int       `json:"card_master_id_1"`
	CardMasterID2    int       `json:"card_master_id_2"`
	CardMasterID3    int       `json:"card_master_id_3"`
	UserAccessoryID1 int64     `json:"user_accessory_id_1"`
	UserAccessoryID2 int64     `json:"user_accessory_id_2"`
	UserAccessoryID3 int64     `json:"user_accessory_id_3"`
}

// PartyName ...
type PartyName struct {
	DotUnderText string `json:"dot_under_text"`
}

// DeckSquadDict ...
type DeckSquadDict struct {
	CardMasterIds    []int   `json:"card_master_ids"`
	UserAccessoryIds []int64 `json:"user_accessory_ids"`
}

// LiveDaily ...
type LiveDaily struct {
	LiveDailyMasterID      int `json:"live_daily_master_id" xorm:"id"`
	LiveMasterID           int `json:"live_master_id" xorm:"live_id"`
	EndAt                  int `json:"end_at"`
	RemainingPlayCount     int `json:"remaining_play_count"`
	RemainingRecoveryCount int `json:"remaining_recovery_count"`
}

// LiveStartReq ...
type LiveStartReq struct {
	LiveDifficultyID    int  `json:"live_difficulty_id"`
	DeckID              int  `json:"deck_id"`
	PartnerUserID       int  `json:"partner_user_id"`
	PartnerCardMasterID int  `json:"partner_card_master_id"`
	LpMagnification     int  `json:"lp_magnification"`
	IsAutoPlay          bool `json:"is_auto_play"`
	IsReferenceBook     bool `json:"is_reference_book"`
}

// LivePartnerInfo ...
type LivePartnerInfo struct {
	UserID                              int                 `json:"user_id"`
	Name                                PartnerName         `json:"name"`
	Rank                                int                 `json:"rank"`
	LastPlayedAt                        int64               `json:"last_played_at"`
	RecommendCardMasterID               int                 `json:"recommend_card_master_id"`
	RecommendCardLevel                  int                 `json:"recommend_card_level"`
	IsRecommendCardImageAwaken          bool                `json:"is_recommend_card_image_awaken"`
	IsRecommendCardAllTrainingActivated bool                `json:"is_recommend_card_all_training_activated"`
	EmblemID                            int                 `json:"emblem_id"`
	IsNew                               bool                `json:"is_new"`
	IntroductionMessage                 IntroductionMessage `json:"introduction_message"`
	FriendApprovedAt                    any                 `json:"friend_approved_at"`
	RequestStatus                       int                 `json:"request_status"`
	IsRequestPending                    bool                `json:"is_request_pending"`
}

// PartnerName ...
type PartnerName struct {
	DotUnderText string `json:"dot_under_text"`
}

// IntroductionMessage ...
type IntroductionMessage struct {
	DotUnderText string `json:"dot_under_text"`
}

// LiveResultAchievementStatus ...
type LiveResultAchievementStatus struct {
	ClearCount       int64 `json:"clear_count"`
	GotVoltage       int64 `json:"got_voltage"`
	RemainingStamina int64 `json:"remaining_stamina"`
}

// MvpInfo ...
type MvpInfo struct {
	CardMasterID        int64 `json:"card_master_id"`
	GetVoltage          int64 `json:"get_voltage"`
	SkillTriggeredCount int64 `json:"skill_triggered_count"`
	AppealCount         int64 `json:"appeal_count"`
}

// LiveSaveDeckReq ...
type LiveSaveDeckReq struct {
	LiveMasterID        int   `json:"live_master_id"`
	LiveMvDeckType      int   `json:"live_mv_deck_type"`
	MemberMasterIDByPos []int `json:"member_master_id_by_pos"`
	SuitMasterIDByPos   []int `json:"suit_master_id_by_pos"`
	ViewStatusByPos     []int `json:"view_status_by_pos"`
}

// UserLiveMvDeckInfo ...
type UserLiveMvDeckInfo struct {
	LiveMasterID     any `json:"live_master_id"`
	MemberMasterID1  any `json:"member_master_id_1"`
	MemberMasterID2  any `json:"member_master_id_2"`
	MemberMasterID3  any `json:"member_master_id_3"`
	MemberMasterID4  any `json:"member_master_id_4"`
	MemberMasterID5  any `json:"member_master_id_5"`
	MemberMasterID6  any `json:"member_master_id_6"`
	MemberMasterID7  any `json:"member_master_id_7"`
	MemberMasterID8  any `json:"member_master_id_8"`
	MemberMasterID9  any `json:"member_master_id_9"`
	MemberMasterID10 any `json:"member_master_id_10"`
	MemberMasterID11 any `json:"member_master_id_11"`
	MemberMasterID12 any `json:"member_master_id_12"`
	SuitMasterID1    any `json:"suit_master_id_1"`
	SuitMasterID2    any `json:"suit_master_id_2"`
	SuitMasterID3    any `json:"suit_master_id_3"`
	SuitMasterID4    any `json:"suit_master_id_4"`
	SuitMasterID5    any `json:"suit_master_id_5"`
	SuitMasterID6    any `json:"suit_master_id_6"`
	SuitMasterID7    any `json:"suit_master_id_7"`
	SuitMasterID8    any `json:"suit_master_id_8"`
	SuitMasterID9    any `json:"suit_master_id_9"`
	SuitMasterID10   any `json:"suit_master_id_10"`
	SuitMasterID11   any `json:"suit_master_id_11"`
	SuitMasterID12   any `json:"suit_master_id_12"`
}
