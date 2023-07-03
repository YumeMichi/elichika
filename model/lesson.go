package model

type LessonMenuAction struct {
	CardMasterID                  int64 `json:"card_master_id"`
	Position                      int   `json:"position"`
	IsAddedPassiveSkill           bool  `json:"is_added_passive_skill"`
	IsAddedSpecialPassiveSkill    bool  `json:"is_added_special_passive_skill"`
	IsRankupedPassiveSkill        bool  `json:"is_rankuped_passive_skill"`
	IsRankupedSpecialPassiveSkill bool  `json:"is_rankuped_special_passive_skill"`
	IsPromotedSkill               bool  `json:"is_promoted_skill"`
	MaxRarity                     any   `json:"max_rarity"`
	UpCount                       int   `json:"up_count"`
}
