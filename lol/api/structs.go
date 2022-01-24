package api

type MatchMetadata struct {
	DataVersion  string   `json:"dataVersion"`
	MatchId      string   `json:"matchId"`
	Participants []string `json:"participants"`
}

type MatchTeam struct {
	Bans []struct {
		ChampionId int `json:"championId"`
		PickTurn   int `json:"pickTurn"`
	} `json:"bans"`
	Objectives struct {
		Baron struct {
			First bool `json:"first"`
			Kills int  `json:"kills"`
		} `json:"baron"`
		Champion struct {
			First bool `json:"first"`
			Kills int  `json:"kills"`
		} `json:"champion"`
		Dragon struct {
			First bool `json:"first"`
			Kills int  `json:"kills"`
		} `json:"dragon"`
		Inhibitor struct {
			First bool `json:"first"`
			Kills int  `json:"kills"`
		} `json:"inhibitor"`
		RiftHerald struct {
			First bool `json:"first"`
			Kills int  `json:"kills"`
		} `json:"riftHerald"`
		Tower struct {
			First bool `json:"first"`
			Kills int  `json:"kills"`
		} `json:"tower"`
	} `json:"objectives"`
	TeamId int  `json:"teamId"`
	Win    bool `json:"win"`
}

type MatchParticipant struct {
	Assists                     int    `json:"assists"`
	BaronKills                  int    `json:"baronKills"`
	BountyLevel                 int    `json:"bountyLevel"`
	ChampExperience             int    `json:"champExperience"`
	ChampLevel                  int    `json:"champLevel"`
	ChampionId                  int    `json:"championId"`
	ChampionName                string `json:"championName"`
	ChampionTransform           int    `json:"championTransform"`
	ConsumablesPurchased        int    `json:"consumablesPurchased"`
	DamageDealtToBuildings      int    `json:"damageDealtToBuildings"`
	DamageDealtToObjectives     int    `json:"damageDealtToObjectives"`
	DamageDealtToTurrets        int    `json:"damageDealtToTurrets"`
	DamageSelfMitigated         int    `json:"damageSelfMitigated"`
	Deaths                      int    `json:"deaths"`
	DetectorWardsPlaced         int    `json:"detectorWardsPlaced"`
	DoubleKills                 int    `json:"doubleKills"`
	DragonKills                 int    `json:"dragonKills"`
	FirstBloodAssist            bool   `json:"firstBloodAssist"`
	FirstBloodKill              bool   `json:"firstBloodKill"`
	FirstTowerAssist            bool   `json:"firstTowerAssist"`
	FirstTowerKill              bool   `json:"firstTowerKill"`
	GameEndedInEarlySurrender   bool   `json:"gameEndedInEarlySurrender"`
	GameEndedInSurrender        bool   `json:"gameEndedInSurrender"`
	GoldEarned                  int    `json:"goldEarned"`
	GoldSpent                   int    `json:"goldSpent"`
	IndividualPosition          string `json:"individualPosition"`
	InhibitorKills              int    `json:"inhibitorKills"`
	InhibitorTakedowns          int    `json:"inhibitorTakedowns"`
	InhibitorsLost              int    `json:"inhibitorsLost"`
	Item0                       int    `json:"item0"`
	Item1                       int    `json:"item1"`
	Item2                       int    `json:"item2"`
	Item3                       int    `json:"item3"`
	Item4                       int    `json:"item4"`
	Item5                       int    `json:"item5"`
	Item6                       int    `json:"item6"`
	ItemsPurchased              int    `json:"itemsPurchased"`
	KillingSprees               int    `json:"killingSprees"`
	Kills                       int    `json:"kills"`
	Lane                        string `json:"lane"`
	LargestCriticalStrike       int    `json:"largestCriticalStrike"`
	LargestKillingSpree         int    `json:"largestKillingSpree"`
	LargestMultiKill            int    `json:"largestMultiKill"`
	LongestTimeSpentLiving      int    `json:"longestTimeSpentLiving"`
	MagicDamageDealt            int    `json:"magicDamageDealt"`
	MagicDamageDealtToChampions int    `json:"magicDamageDealtToChampions"`
	MagicDamageTaken            int    `json:"magicDamageTaken"`
	NeutralMinionsKilled        int    `json:"neutralMinionsKilled"`
	NexusKills                  int    `json:"nexusKills"`
	NexusLost                   int    `json:"nexusLost"`
	NexusTakedowns              int    `json:"nexusTakedowns"`
	ObjectivesStolen            int    `json:"objectivesStolen"`
	ObjectivesStolenAssists     int    `json:"objectivesStolenAssists"`
	ParticipantId               int    `json:"participantId"`
	PentaKills                  int    `json:"pentaKills"`
	Perks                       struct {
		StatPerks struct {
			Defense int `json:"defense"`
			Flex    int `json:"flex"`
			Offense int `json:"offense"`
		} `json:"statPerks"`
		Styles []struct {
			Description string `json:"description"`
			Selections  []struct {
				Perk int `json:"perk"`
				Var1 int `json:"var1"`
				Var2 int `json:"var2"`
				Var3 int `json:"var3"`
			} `json:"selections"`
			Style int `json:"style"`
		} `json:"styles"`
	} `json:"perks"`
	PhysicalDamageDealt            int    `json:"physicalDamageDealt"`
	PhysicalDamageDealtToChampions int    `json:"physicalDamageDealtToChampions"`
	PhysicalDamageTaken            int    `json:"physicalDamageTaken"`
	ProfileIcon                    int    `json:"profileIcon"`
	Puuid                          string `json:"puuid"`
	QuadraKills                    int    `json:"quadraKills"`
	RiotIdName                     string `json:"riotIdName"`
	RiotIdTagline                  string `json:"riotIdTagline"`
	Role                           string `json:"role"`
	SightWardsBoughtInGame         int    `json:"sightWardsBoughtInGame"`
	Spell1Casts                    int    `json:"spell1Casts"`
	Spell2Casts                    int    `json:"spell2Casts"`
	Spell3Casts                    int    `json:"spell3Casts"`
	Spell4Casts                    int    `json:"spell4Casts"`
	Summoner1Casts                 int    `json:"summoner1Casts"`
	Summoner1Id                    int    `json:"summoner1Id"`
	Summoner2Casts                 int    `json:"summoner2Casts"`
	Summoner2Id                    int    `json:"summoner2Id"`
	SummonerId                     string `json:"summonerId"`
	SummonerLevel                  int    `json:"summonerLevel"`
	SummonerName                   string `json:"summonerName"`
	TeamEarlySurrendered           bool   `json:"teamEarlySurrendered"`
	TeamId                         int    `json:"teamId"`
	TeamPosition                   string `json:"teamPosition"`
	TimeCCingOthers                int    `json:"timeCCingOthers"`
	TimePlayed                     int    `json:"timePlayed"`
	TotalDamageDealt               int    `json:"totalDamageDealt"`
	TotalDamageDealtToChampions    int    `json:"totalDamageDealtToChampions"`
	TotalDamageShieldedOnTeammates int    `json:"totalDamageShieldedOnTeammates"`
	TotalDamageTaken               int    `json:"totalDamageTaken"`
	TotalHeal                      int    `json:"totalHeal"`
	TotalHealsOnTeammates          int    `json:"totalHealsOnTeammates"`
	TotalMinionsKilled             int    `json:"totalMinionsKilled"`
	TotalTimeCCDealt               int    `json:"totalTimeCCDealt"`
	TotalTimeSpentDead             int    `json:"totalTimeSpentDead"`
	TotalUnitsHealed               int    `json:"totalUnitsHealed"`
	TripleKills                    int    `json:"tripleKills"`
	TrueDamageDealt                int    `json:"trueDamageDealt"`
	TrueDamageDealtToChampions     int    `json:"trueDamageDealtToChampions"`
	TrueDamageTaken                int    `json:"trueDamageTaken"`
	TurretKills                    int    `json:"turretKills"`
	TurretTakedowns                int    `json:"turretTakedowns"`
	TurretsLost                    int    `json:"turretsLost"`
	UnrealKills                    int    `json:"unrealKills"`
	VisionScore                    int    `json:"visionScore"`
	VisionWardsBoughtInGame        int    `json:"visionWardsBoughtInGame"`
	WardsKilled                    int    `json:"wardsKilled"`
	WardsPlaced                    int    `json:"wardsPlaced"`
	Win                            bool   `json:"win"`
}

type MatchInfo struct {
	GameCreation       int64              `json:"gameCreation"`
	GameDuration       int                `json:"gameDuration"`
	GameEndTimestamp   int64              `json:"gameEndTimestamp"`
	GameId             int64              `json:"gameId"`
	GameMode           string             `json:"gameMode"`
	GameName           string             `json:"gameName"`
	GameStartTimestamp int64              `json:"gameStartTimestamp"`
	GameType           string             `json:"gameType"`
	GameVersion        string             `json:"gameVersion"`
	MapId              int                `json:"mapId"`
	Participants       []MatchParticipant `json:"participants"`
	PlatformId         string             `json:"platformId"`
	QueueId            int                `json:"queueId"`
	Teams              []MatchTeam        `json:"teams"`
	TournamentCode     string             `json:"tournamentCode"`
}

type MatchTimelineFrameParticipant struct {
	ChampionStats struct {
		AbilityHaste         int `json:"abilityHaste"`
		AbilityPower         int `json:"abilityPower"`
		Armor                int `json:"armor"`
		ArmorPen             int `json:"armorPen"`
		ArmorPenPercent      int `json:"armorPenPercent"`
		AttackDamage         int `json:"attackDamage"`
		AttackSpeed          int `json:"attackSpeed"`
		BonusArmorPenPercent int `json:"bonusArmorPenPercent"`
		BonusMagicPenPercent int `json:"bonusMagicPenPercent"`
		CcReduction          int `json:"ccReduction"`
		CooldownReduction    int `json:"cooldownReduction"`
		Health               int `json:"health"`
		HealthMax            int `json:"healthMax"`
		HealthRegen          int `json:"healthRegen"`
		Lifesteal            int `json:"lifesteal"`
		MagicPen             int `json:"magicPen"`
		MagicPenPercent      int `json:"magicPenPercent"`
		MagicResist          int `json:"magicResist"`
		MovementSpeed        int `json:"movementSpeed"`
		Omnivamp             int `json:"omnivamp"`
		PhysicalVamp         int `json:"physicalVamp"`
		Power                int `json:"power"`
		PowerMax             int `json:"powerMax"`
		PowerRegen           int `json:"powerRegen"`
		SpellVamp            int `json:"spellVamp"`
	} `json:"championStats"`
	CurrentGold int `json:"currentGold"`
	DamageStats struct {
		MagicDamageDone               int `json:"magicDamageDone"`
		MagicDamageDoneToChampions    int `json:"magicDamageDoneToChampions"`
		MagicDamageTaken              int `json:"magicDamageTaken"`
		PhysicalDamageDone            int `json:"physicalDamageDone"`
		PhysicalDamageDoneToChampions int `json:"physicalDamageDoneToChampions"`
		PhysicalDamageTaken           int `json:"physicalDamageTaken"`
		TotalDamageDone               int `json:"totalDamageDone"`
		TotalDamageDoneToChampions    int `json:"totalDamageDoneToChampions"`
		TotalDamageTaken              int `json:"totalDamageTaken"`
		TrueDamageDone                int `json:"trueDamageDone"`
		TrueDamageDoneToChampions     int `json:"trueDamageDoneToChampions"`
		TrueDamageTaken               int `json:"trueDamageTaken"`
	} `json:"damageStats"`
	GoldPerSecond       int `json:"goldPerSecond"`
	JungleMinionsKilled int `json:"jungleMinionsKilled"`
	Level               int `json:"level"`
	MinionsKilled       int `json:"minionsKilled"`
	ParticipantId       int `json:"participantId"`
	Position            struct {
		X int `json:"x"`
		Y int `json:"y"`
	} `json:"position"`
	TimeEnemySpentControlled int `json:"timeEnemySpentControlled"`
	TotalGold                int `json:"totalGold"`
	Xp                       int `json:"xp"`
}

type MatchTimelineInfo struct {
	FrameInterval int `json:"frameInterval"`
	Frames        []struct {
		Events []struct {
			RealTimestamp           int64  `json:"realTimestamp,omitempty"`
			Timestamp               int    `json:"timestamp"`
			Type                    string `json:"type"`
			LevelUpType             string `json:"levelUpType,omitempty"`
			ParticipantId           int    `json:"participantId,omitempty"`
			SkillSlot               int    `json:"skillSlot,omitempty"`
			ItemId                  int    `json:"itemId,omitempty"`
			CreatorId               int    `json:"creatorId,omitempty"`
			WardType                string `json:"wardType,omitempty"`
			Level                   int    `json:"level,omitempty"`
			AssistingParticipantIds []int  `json:"assistingParticipantIds,omitempty"`
			Bounty                  int    `json:"bounty,omitempty"`
			KillStreakLength        int    `json:"killStreakLength,omitempty"`
			KillerId                int    `json:"killerId,omitempty"`
			Position                struct {
				X int `json:"x"`
				Y int `json:"y"`
			} `json:"position,omitempty"`
			VictimDamageDealt []struct {
				Basic          bool   `json:"basic"`
				MagicDamage    int    `json:"magicDamage"`
				Name           string `json:"name"`
				ParticipantId  int    `json:"participantId"`
				PhysicalDamage int    `json:"physicalDamage"`
				SpellName      string `json:"spellName"`
				SpellSlot      int    `json:"spellSlot"`
				TrueDamage     int    `json:"trueDamage"`
				Type           string `json:"type"`
			} `json:"victimDamageDealt,omitempty"`
			VictimDamageReceived []struct {
				Basic          bool   `json:"basic"`
				MagicDamage    int    `json:"magicDamage"`
				Name           string `json:"name"`
				ParticipantId  int    `json:"participantId"`
				PhysicalDamage int    `json:"physicalDamage"`
				SpellName      string `json:"spellName"`
				SpellSlot      int    `json:"spellSlot"`
				TrueDamage     int    `json:"trueDamage"`
				Type           string `json:"type"`
			} `json:"victimDamageReceived,omitempty"`
			VictimId        int    `json:"victimId,omitempty"`
			KillType        string `json:"killType,omitempty"`
			AfterId         int    `json:"afterId,omitempty"`
			BeforeId        int    `json:"beforeId,omitempty"`
			GoldGain        int    `json:"goldGain,omitempty"`
			LaneType        string `json:"laneType,omitempty"`
			TeamId          int    `json:"teamId,omitempty"`
			BuildingType    string `json:"buildingType,omitempty"`
			TowerType       string `json:"towerType,omitempty"`
			KillerTeamId    int    `json:"killerTeamId,omitempty"`
			MonsterType     string `json:"monsterType,omitempty"`
			MultiKillLength int    `json:"multiKillLength,omitempty"`
			MonsterSubType  string `json:"monsterSubType,omitempty"`
			GameId          int64  `json:"gameId,omitempty"`
			WinningTeam     int    `json:"winningTeam,omitempty"`
		} `json:"events"`
		ParticipantFrames struct {
			Field1  MatchTimelineFrameParticipant `json:"1"`
			Field2  MatchTimelineFrameParticipant `json:"2"`
			Field3  MatchTimelineFrameParticipant `json:"3"`
			Field4  MatchTimelineFrameParticipant `json:"4"`
			Field5  MatchTimelineFrameParticipant `json:"5"`
			Field6  MatchTimelineFrameParticipant `json:"6"`
			Field7  MatchTimelineFrameParticipant `json:"7"`
			Field8  MatchTimelineFrameParticipant `json:"8"`
			Field9  MatchTimelineFrameParticipant `json:"9"`
			Field10 MatchTimelineFrameParticipant `json:"10"`
		} `json:"participantFrames"`
		Timestamp int `json:"timestamp"`
	} `json:"frames"`
	GameId       int64 `json:"gameId"`
	Participants []struct {
		ParticipantId int    `json:"participantId"`
		Puuid         string `json:"puuid"`
	} `json:"participants"`
}

type Champion struct {
	Id    string `json:"id"`
	Key   string `json:"key"`
	Name  string `json:"name"`
	Title string `json:"title"`
	Image struct {
		Full   string `json:"full"`
		Sprite string `json:"sprite"`
		Group  string `json:"group"`
		X      int    `json:"x"`
		Y      int    `json:"y"`
		W      int    `json:"w"`
		H      int    `json:"h"`
	} `json:"image"`
	Skins []struct {
		Id      string `json:"id"`
		Num     int    `json:"num"`
		Name    string `json:"name"`
		Chromas bool   `json:"chromas"`
	} `json:"skins"`
	Lore      string   `json:"lore"`
	Blurb     string   `json:"blurb"`
	Allytips  []string `json:"allytips"`
	Enemytips []string `json:"enemytips"`
	Tags      []string `json:"tags"`
	Partype   string   `json:"partype"`
	Info      struct {
		Attack     int `json:"attack"`
		Defense    int `json:"defense"`
		Magic      int `json:"magic"`
		Difficulty int `json:"difficulty"`
	} `json:"info"`
	Stats struct {
		Hp                   float64 `json:"hp"`
		Hpperlevel           float64 `json:"hpperlevel"`
		Mp                   float64 `json:"mp"`
		Mpperlevel           float64 `json:"mpperlevel"`
		Movespeed            float64 `json:"movespeed"`
		Armor                float64 `json:"armor"`
		Armorperlevel        float64 `json:"armorperlevel"`
		Spellblock           float64 `json:"spellblock"`
		Spellblockperlevel   float64 `json:"spellblockperlevel"`
		Attackrange          float64 `json:"attackrange"`
		Hpregen              float64 `json:"hpregen"`
		Hpregenperlevel      float64 `json:"hpregenperlevel"`
		Mpregen              float64 `json:"mpregen"`
		Mpregenperlevel      float64 `json:"mpregenperlevel"`
		Crit                 float64 `json:"crit"`
		Critperlevel         float64 `json:"critperlevel"`
		Attackdamage         float64 `json:"attackdamage"`
		Attackdamageperlevel float64 `json:"attackdamageperlevel"`
		Attackspeedperlevel  float64 `json:"attackspeedperlevel"`
		Attackspeed          float64 `json:"attackspeed"`
	} `json:"stats"`
	Spells []struct {
		Id          string `json:"id"`
		Name        string `json:"name"`
		Description string `json:"description"`
		Tooltip     string `json:"tooltip"`
		Leveltip    struct {
			Label  []string `json:"label"`
			Effect []string `json:"effect"`
		} `json:"leveltip"`
		Maxrank      int       `json:"maxrank"`
		Cooldown     []float64 `json:"cooldown"`
		CooldownBurn string    `json:"cooldownBurn"`
		Cost         []int     `json:"cost"`
		CostBurn     string    `json:"costBurn"`
		Datavalues   struct {
		} `json:"datavalues"`
		Effect     [][]float64   `json:"effect"`
		EffectBurn []string      `json:"effectBurn"`
		Vars       []interface{} `json:"vars"`
		CostType   string        `json:"costType"`
		Maxammo    string        `json:"maxammo"`
		Range      []int         `json:"range"`
		RangeBurn  string        `json:"rangeBurn"`
		Image      struct {
			Full   string `json:"full"`
			Sprite string `json:"sprite"`
			Group  string `json:"group"`
			X      int    `json:"x"`
			Y      int    `json:"y"`
			W      int    `json:"w"`
			H      int    `json:"h"`
		} `json:"image"`
		Resource string `json:"resource"`
	} `json:"spells"`
	Passive struct {
		Name        string `json:"name"`
		Description string `json:"description"`
		Image       struct {
			Full   string `json:"full"`
			Sprite string `json:"sprite"`
			Group  string `json:"group"`
			X      int    `json:"x"`
			Y      int    `json:"y"`
			W      int    `json:"w"`
			H      int    `json:"h"`
		} `json:"image"`
	} `json:"passive"`
	Recommended []interface{} `json:"recommended"`
}
