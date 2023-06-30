package model

// CardAwakeningReq ...
type CardAwakeningReq struct {
	CardMasterID     int  `json:"card_master_id"`
	IsAwakeningImage bool `json:"is_awakening_image"`
}

// CardInfo ...
type CardInfo struct {
	CardMasterID               int   `json:"card_master_id"`
	Level                      int   `json:"level"`
	Exp                        int   `json:"exp"`
	LovePoint                  int   `json:"love_point"`
	IsFavorite                 bool  `json:"is_favorite"`
	IsAwakening                bool  `json:"is_awakening"`
	IsAwakeningImage           bool  `json:"is_awakening_image"`
	IsAllTrainingActivated     bool  `json:"is_all_training_activated"`
	TrainingActivatedCellCount int   `json:"training_activated_cell_count"`
	MaxFreePassiveSkill        int   `json:"max_free_passive_skill"`
	Grade                      int   `json:"grade"`
	TrainingLife               int   `json:"training_life"`
	TrainingAttack             int   `json:"training_attack"`
	TrainingDexterity          int   `json:"training_dexterity"`
	ActiveSkillLevel           int   `json:"active_skill_level"`
	PassiveSkillALevel         int   `json:"passive_skill_a_level"`
	PassiveSkillBLevel         int   `json:"passive_skill_b_level"`
	PassiveSkillCLevel         int   `json:"passive_skill_c_level"`
	AdditionalPassiveSkill1ID  int   `json:"additional_passive_skill_1_id"`
	AdditionalPassiveSkill2ID  int   `json:"additional_passive_skill_2_id"`
	AdditionalPassiveSkill3ID  int   `json:"additional_passive_skill_3_id"`
	AdditionalPassiveSkill4ID  int   `json:"additional_passive_skill_4_id"`
	AcquiredAt                 int64 `json:"acquired_at"`
	IsNew                      bool  `json:"is_new"`
}
