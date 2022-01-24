package datadragon

type Spell struct {
	MScriptName string `json:"mScriptName,omitempty"`
	MSpell      struct {
		Flags             int    `json:"flags,omitempty"`
		MAlternateName    string `json:"mAlternateName"`
		MAffectsTypeFlags int    `json:"mAffectsTypeFlags,omitempty"`
		MRequiredUnitTags struct {
			MTagList string `json:"mTagList,omitempty"`
			Type     string `json:"__type,omitempty"`
		} `json:"mRequiredUnitTags,omitempty"`
		MSpellTags    []string `json:"mSpellTags,omitempty"`
		MEffectAmount []struct {
			Value []float64 `json:"value,omitempty"`
			Type  string    `json:"__type"`
		} `json:"mEffectAmount"`
		MDataValues []struct {
			MName    string    `json:"mName,omitempty"`
			MValues  []float64 `json:"mValues,omitempty"`
			MBaseP   float64   `json:"mBaseP,omitempty"`
			MFormula string    `json:"mFormula,omitempty"`
			Type     string    `json:"__type,omitempty"`
		} `json:"mDataValues,omitempty"`
		MSpellCalculations              map[string]interface{} `json:"mSpellCalculations,omitempty"`
		MCoefficient                    float64                `json:"mCoefficient,omitempty"`
		MAnimationName                  string                 `json:"mAnimationName,omitempty"`
		MImgIconName                    []string               `json:"mImgIconName,omitempty"`
		MCastTime                       float64                `json:"mCastTime,omitempty"`
		CooldownTime                    []float64              `json:"cooldownTime,omitempty"`
		DelayCastOffsetPercent          float64                `json:"delayCastOffsetPercent,omitempty"`
		DelayTotalTimePercent           float64                `json:"delayTotalTimePercent,omitempty"`
		MCantCancelWhileWindingUp       bool                   `json:"mCantCancelWhileWindingUp,omitempty"`
		MChannelIsInterruptedByDisables bool                   `json:"mChannelIsInterruptedByDisables,omitempty"`
		MCanMoveWhileChanneling         bool                   `json:"mCanMoveWhileChanneling,omitempty"`
		UseAnimatorFramerate            bool                   `json:"useAnimatorFramerate,omitempty"`
		CastRange                       []float64              `json:"castRange,omitempty"`
		CastRadius                      []float64              `json:"castRadius,omitempty"`
		CastConeDistance                float64                `json:"castConeDistance,omitempty"`
		CastFrame                       float64                `json:"castFrame,omitempty"`
		MissileSpeed                    float64                `json:"missileSpeed,omitempty"`
		BHaveHitBone                    bool                   `json:"bHaveHitBone,omitempty"`
		MHitBoneName                    string                 `json:"mHitBoneName,omitempty"`
		MTargetingTypeData              struct {
			Type string `json:"__type,omitempty"`
		} `json:"mTargetingTypeData,omitempty"`
		MClientData ClientData `json:"mClientData,omitempty"`
		Type        string     `json:"__type,omitempty"`
	} `json:"mSpell,omitempty"`
	MBuff struct {
		MDescription string `json:"mDescription,omitempty"`
		Type         string `json:"__type,omitempty"`
	} `json:"mBuff,omitempty"`
	Type string `json:"__type,omitempty"`
}

type ClientData struct {
	MTooltipData struct {
		MObjectName string `json:"mObjectName,omitempty"`
		MFormat     string `json:"mFormat,omitempty"`
		MLocKeys    struct {
			KeyName                     string `json:"keyName,omitempty"`
			KeyCost                     string `json:"keyCost,omitempty"`
			KeyTooltip                  string `json:"keyTooltip,omitempty"`
			KeySummary                  string `json:"keySummary,omitempty"`
			KeyTooltipExtendedBelowLine string `json:"keyTooltipExtendedBelowLine,omitempty"`
		} `json:"mLocKeys,omitempty"`
		MLists struct {
			LevelUp struct {
				LevelCount int `json:"levelCount,omitempty"`
				Elements   []struct {
					Type         string  `json:"type,omitempty"`
					Multiplier   float64 `json:"multiplier,omitempty"`
					NameOverride string  `json:"nameOverride,omitempty"`
					Style        int     `json:"Style,omitempty"`
					Type1        string  `json:"__type,omitempty"`
					TypeIndex    int     `json:"typeIndex,omitempty"`
				} `json:"elements,omitempty"`
				Type string `json:"__type,omitempty"`
			} `json:"LevelUp,omitempty"`
		} `json:"mLists,omitempty"`
		Type string `json:"__type,omitempty"`
	} `json:"mTooltipData,omitempty"`
	MTargeterDefinitions []struct {
		UseCasterBoundingBox bool `json:"useCasterBoundingBox,omitempty"`
		OverrideBaseRange    struct {
			MPerLevelValues []float64 `json:"mPerLevelValues,omitempty"`
			MValueType      int       `json:"mValueType,omitempty"`
			Type            string    `json:"__type,omitempty"`
		} `json:"overrideBaseRange,omitempty"`
		EndLocator struct {
			BasePosition    int    `json:"basePosition,omitempty"`
			OrientationType int    `json:"orientationType,omitempty"`
			Type            string `json:"__type"`
		} `json:"endLocator"`
		AlwaysDraw            bool    `json:"alwaysDraw,omitempty"`
		MinimumDisplayedRange float64 `json:"minimumDisplayedRange,omitempty"`
		LineWidth             struct {
			MPerLevelValues []float64 `json:"mPerLevelValues,omitempty"`
			MValueType      int       `json:"mValueType,omitempty"`
			Type            string    `json:"__type"`
		} `json:"lineWidth"`
		TextureBaseOverrideName   string `json:"textureBaseOverrideName,omitempty"`
		TextureTargetOverrideName string `json:"textureTargetOverrideName,omitempty"`
		Type                      string `json:"__type,omitempty"`
	} `json:"mTargeterDefinitions,omitempty"`
	Type string `json:"__type,omitempty"`
}

type Ability struct {
	MRootSpell   string   `json:"mRootSpell"`
	MChildSpells []string `json:"mChildSpells"`
	MName        string   `json:"mName"`
	MType        int      `json:"mType"`
	Type         string   `json:"__type"`
}

type Champion struct {
	Character CharacterRecord
	Spells    []Spell
	Abilities []Ability
}

type ApiAbilityResource struct {
	ArType                        int     `json:"arType,omitempty"`
	ArBase                        float64 `json:"arBase,omitempty"`
	ArPerLevel                    float64 `json:"arPerLevel,omitempty"`
	ArBaseStaticRegen             float64 `json:"arBaseStaticRegen,omitempty"`
	ArIncrements                  float64 `json:"arIncrements,omitempty"`
	ArMaxSegments                 int     `json:"arMaxSegments,omitempty"`
	ArAllowMaxValueToBeOverridden bool    `json:"arAllowMaxValueToBeOverridden,omitempty"`
	ArRegenPerLevel               float64 `json:"arRegenPerLevel,omitempty"`
	Type                          string  `json:"__type"`
}

type ApiAttack struct {
	MAttackTotalTime                              float64 `json:"mAttackTotalTime,omitempty"`
	MAttackCastTime                               float64 `json:"mAttackCastTime,omitempty"`
	MAttackDelayCastOffsetPercentAttackSpeedRatio float64 `json:"mAttackDelayCastOffsetPercentAttackSpeedRatio,omitempty"`
	MAttackName                                   string  `json:"mAttackName,omitempty"`
	Type                                          string  `json:"__type"`
}

type CharacterRecord struct {
	MCharacterName             string             `json:"mCharacterName"`
	BaseHP                     float64            `json:"baseHP"`
	HpPerLevel                 float64            `json:"hpPerLevel"`
	BaseStaticHPRegen          float64            `json:"baseStaticHPRegen"`
	HpRegenPerLevel            float64            `json:"hpRegenPerLevel"`
	HealthBarHeight            float64            `json:"healthBarHeight"`
	PrimaryAbilityResource     ApiAbilityResource `json:"primaryAbilityResource"`
	SecondaryAbilityResource   ApiAbilityResource `json:"secondaryAbilityResource"`
	BaseDamage                 float64            `json:"baseDamage"`
	DamagePerLevel             float64            `json:"damagePerLevel"`
	BaseArmor                  float64            `json:"baseArmor"`
	ArmorPerLevel              float64            `json:"armorPerLevel"`
	BaseSpellBlock             float64            `json:"baseSpellBlock"`
	SpellBlockPerLevel         float64            `json:"spellBlockPerLevel"`
	BaseMoveSpeed              float64            `json:"baseMoveSpeed"`
	AttackRange                float64            `json:"attackRange"`
	AttackSpeed                float64            `json:"attackSpeed"`
	AttackSpeedRatio           float64            `json:"attackSpeedRatio"`
	AttackSpeedPerLevel        float64            `json:"attackSpeedPerLevel"`
	BasicAttack                ApiAttack          `json:"basicAttack"`
	ExtraAttacks               []ApiAttack        `json:"extraAttacks"`
	CritAttacks                []ApiAttack        `json:"critAttacks"`
	SpellNames                 []string           `json:"spellNames"`
	MAbilities                 []string           `json:"mAbilities"`
	PassiveLuaName             string             `json:"passiveLuaName"`
	PassiveSpell               string             `json:"passiveSpell"`
	Passive1IconName           string             `json:"passive1IconName"`
	Name                       string             `json:"name"`
	SelectionHeight            float64            `json:"selectionHeight"`
	SelectionRadius            float64            `json:"selectionRadius"`
	PathfindingCollisionRadius float64            `json:"pathfindingCollisionRadius"`
	UnitTagsString             string             `json:"unitTagsString"`
	CharacterToolData          struct {
		SearchTags          string `json:"searchTags"`
		SearchTagsSecondary string `json:"searchTagsSecondary"`
		ChampionId          int    `json:"championId"`
		Roles               string `json:"roles"`
		Description         string `json:"description"`
		Classification      string `json:"classification"`
		Type                string `json:"__type"`
	} `json:"characterToolData"`
	PlatformEnabled      bool     `json:"platformEnabled"`
	UseRiotRelationships bool     `json:"useRiotRelationships"`
	PurchaseIdentities   []string `json:"purchaseIdentities"`
	MPreferredPerkStyle  string   `json:"mPreferredPerkStyle"`
	MPerkReplacements    struct {
		MReplacements []struct {
			MReplaceTarget string `json:"mReplaceTarget"`
			MReplaceWith   string `json:"mReplaceWith"`
			Type           string `json:"__type"`
		} `json:"mReplacements"`
		Type string `json:"__type"`
	} `json:"mPerkReplacements"`
	MCharacterPassiveSpell string `json:"mCharacterPassiveSpell"`
	MCharacterPassiveBuffs []struct {
		Bd3C31E4 string `json:"{bd3c31e4}"`
		Type     string `json:"__type"`
	} `json:"mCharacterPassiveBuffs"`
	Type string `json:"__type"`
}
