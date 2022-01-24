package data

import "fmt"

type DamageType int
type EffectType int
type StatType int
type StatFormulaType int
type AbilityResourceType int

const (
	Physical DamageType = 0
	Magic    DamageType = 1
	True     DamageType = 2
)

const (
	Heal                 EffectType = 0
	Shield               EffectType = 1
	BonusAttackSpeed     EffectType = 2
	Mana                 EffectType = 3
	BonusMoveSpeedEffect EffectType = 4
	BonusArmor           EffectType = 5
	BonusMagicResist     EffectType = 6
	BonusAP              EffectType = 7
	BonusAD              EffectType = 8
	BonusLifeSteal       EffectType = 9

	CC EffectType = 10
)

var DmgMap = map[string]DamageType{"physicalDamage": Physical, "magicDamage": Magic, "trueDamage": True}
var EffMap = map[string]EffectType{"speed": BonusMoveSpeedEffect, "status": CC, "healing": Heal,
	"shield": Shield, "scaleArmor": BonusArmor,
	"scaleMR":   BonusMagicResist,
	"scaleAD":   BonusAD,
	"scaleAP":   BonusAP,
	"lifeSteal": BonusLifeSteal,
	"scaleMana": Mana,
}

const (
	AbilityPower                 StatType = 0
	Armor                        StatType = 1
	Attack                       StatType = 2
	AttackSpeed                  StatType = 3
	AttackWindupTime             StatType = 4
	MagicResist                  StatType = 5
	MoveSpeed                    StatType = 6
	CritChance                   StatType = 7
	CritDamage                   StatType = 8
	CooldownReduction            StatType = 9
	AbilityHaste                 StatType = 10
	MaxHealth                    StatType = 11
	CurrentHealth                StatType = 12
	PercentMissingHealth         StatType = 13
	Unknown14                    StatType = 14
	LifeSteal                    StatType = 15
	OmniVamp                     StatType = 17
	PhysicalVamp                 StatType = 18
	MagicPenetrationFlat         StatType = 19
	MagicPenetrationPercent      StatType = 20
	BonusMagicPenetrationPercent StatType = 21
	MagicLethality               StatType = 22
	ArmorPenetrationFlat         StatType = 23
	ArmorPenetrationPercent      StatType = 24
	BonusArmorPenetrationPercent StatType = 25
	PhysicalLethality            StatType = 26
	Tenacity                     StatType = 27
	AttackRange                  StatType = 28
	HealthRegenRate              StatType = 29
	ResourceRegenRate            StatType = 30
	Unknown31                    StatType = 31
	Unknown32                    StatType = 32
	DodgeChance                  StatType = 33
)

const (
	Base     StatFormulaType = 0
	Bonus    StatFormulaType = 1
	Total    StatFormulaType = 2
	Unknown3 StatFormulaType = 3
	Unknown4 StatFormulaType = 4
)

var HardCC = []string{"Knock", "Charm", "Fear", "Taunt", "Sleep", "Stun", "Suppress"}
var PartialCC = []string{"Blind", "Cripple", "Disarm", "Drowsy", "Ground", "Pull", "Nearsight", "Polymorph", "Root", "Silence", "Slow"}

func (a AbilityResourceType) String() string {
	return fmt.Sprintf("AbilityResource(%d)", a)
}

func (s StatType) String() string {
	switch s {
	case AbilityPower:
		return "AbilityPower"
	case Armor:
		return "Armor"
	case Attack:
		return "Attack"
	case AttackSpeed:
		return "AttackSpeed"
	case AttackWindupTime:
		return "AttackWindupTime"
	case MagicResist:
		return "MagicResist"
	case MoveSpeed:
		return "MoveSpeed"
	case CritChance:
		return "CritChance"
	case CritDamage:
		return "CritDamage"
	case CooldownReduction:
		return "CooldownReduction"
	case AbilityHaste:
		return "AbilityHaste"
	case MaxHealth:
		return "MaxHealth"
	case CurrentHealth:
		return "CurrentHealth"
	case PercentMissingHealth:
		return "PercentMissingHealth"
	case Unknown14:
		return "Unknown14"
	case LifeSteal:
		return "LifeSteal"
	case OmniVamp:
		return "OmniVamp"
	case PhysicalVamp:
		return "PhysicalVamp"
	case MagicPenetrationFlat:
		return "MagicPenetrationFlat"
	case MagicPenetrationPercent:
		return "MagicPenetrationPercent"
	case BonusMagicPenetrationPercent:
		return "BonusMagicPenetrationPercent"
	case MagicLethality:
		return "MagicLethality"
	case ArmorPenetrationFlat:
		return "ArmorPenetrationFlat"
	case ArmorPenetrationPercent:
		return "ArmorPenetrationPercent"
	case BonusArmorPenetrationPercent:
		return "BonusArmorPenetrationPercent"
	case PhysicalLethality:
		return "PhysicalLethality"
	case Tenacity:
		return "Tenacity"
	case AttackRange:
		return "AttackRange"
	case HealthRegenRate:
		return "HealthRegenRate"
	case ResourceRegenRate:
		return "ResourceRegenRate"
	case Unknown31:
		return "Unknown31"
	case Unknown32:
		return "Unknown32"
	case DodgeChance:
		return "DodgeChance"
	default:
		return ""
	}
}

type CalculationType int
type CalculationPartType int

const (
	UnknownCalculation          CalculationType = 0
	IGameCalculationType        CalculationType = 1
	GameCalculationType         CalculationType = 2
	GameCalculationConditional  CalculationType = 3
	GameCalculationModifiedType CalculationType = 4
)

const (
	UnknownPart                                 CalculationPartType = 0
	IGameCalculationPartType                    CalculationPartType = 1
	EffectValueCalculationPart                  CalculationPartType = 2
	NamedDataValueCalculationPart               CalculationPartType = 3
	CustomReductionMultiplierCalculationPart    CalculationPartType = 4
	ProductOfSubPartsCalculationPart            CalculationPartType = 5
	SumOfSubPartsCalculationPart                CalculationPartType = 6
	NumberCalculationPart                       CalculationPartType = 7
	IGameCalculationPartWithStats               CalculationPartType = 8
	StatByCoefficientCalculationPart            CalculationPartType = 9
	StatBySubPartCalculationPart                CalculationPartType = 10
	StatByNamedDataValueCalculationPart         CalculationPartType = 11
	SubPartScaledProportionalToStat             CalculationPartType = 12
	AbilityResourceByCoefficientCalculationPart CalculationPartType = 13
	IGameCalculationPartWithBuffCounter         CalculationPartType = 14
	BuffCounterByCoefficientCalculationPart     CalculationPartType = 15
	BuffCounterByNamedDataValueCalculationPart  CalculationPartType = 16
	IGameCalculationPartByCharLevel             CalculationPartType = 17
	ByCharLevelInterpolationCalculationPart     CalculationPartType = 18
	ByCharLevelBreakpointsCalculationPart       CalculationPartType = 19
	Breakpoint                                  CalculationPartType = 20
	ByCharLevelFormulaCalculationPart           CalculationPartType = 21
	CooldownMultiplierCalculationPart           CalculationPartType = 22
)

func CalculationTypeFromString(s string) CalculationType {
	switch s {
	case "IGameCalculation":
		return IGameCalculationType
	case "GameCalculation":
		return GameCalculationType
	case "GameCalculationConditional":
		return GameCalculationConditional
	case "GameCalculationModified":
		return GameCalculationModifiedType
	default:
		return UnknownCalculation
	}
}

func (c CalculationType) String() string {
	switch c {
	case UnknownCalculation:
		return "UnknownCalculation"
	case IGameCalculationType:
		return "IGameCalculation"
	case GameCalculationType:
		return "GameCalculation"
	case GameCalculationConditional:
		return "GameCalculationConditional"
	case GameCalculationModifiedType:
		return "GameCalculationModified"
	}
	return ""
}

func (f CalculationPartType) String() string {
	switch f {
	case UnknownPart:
		return "UnknownPart"
	case IGameCalculationPartType:
		return "IGameCalculationPart"
	case EffectValueCalculationPart:
		return "EffectValueCalculationPart"
	case NamedDataValueCalculationPart:
		return "NamedDataValueCalculationPart"
	case CustomReductionMultiplierCalculationPart:
		return "CustomReductionMultiplierCalculationPart"
	case ProductOfSubPartsCalculationPart:
		return "ProductOfSubPartsCalculationPart"
	case SumOfSubPartsCalculationPart:
		return "SumOfSubPartsCalculationPart"
	case NumberCalculationPart:
		return "NumberCalculationPart"
	case IGameCalculationPartWithStats:
		return "IGameCalculationPartWithStats"
	case StatByCoefficientCalculationPart:
		return "StatByCoefficientCalculationPart"
	case StatBySubPartCalculationPart:
		return "StatBySubPartCalculationPart"
	case StatByNamedDataValueCalculationPart:
		return "StatByNamedDataValueCalculationPart"
	case SubPartScaledProportionalToStat:
		return "SubPartScaledProportionalToStat"
	case AbilityResourceByCoefficientCalculationPart:
		return "AbilityResourceByCoefficientCalculationPart"
	case IGameCalculationPartWithBuffCounter:
		return "IGameCalculationPartWithBuffCounter"
	case BuffCounterByCoefficientCalculationPart:
		return "BuffCounterByCoefficientCalculationPart"
	case BuffCounterByNamedDataValueCalculationPart:
		return "BuffCounterByNamedDataValueCalculationPart"
	case IGameCalculationPartByCharLevel:
		return "IGameCalculationPartByCharLevel"
	case ByCharLevelInterpolationCalculationPart:
		return "ByCharLevelInterpolationCalculationPart"
	case ByCharLevelBreakpointsCalculationPart:
		return "ByCharLevelBreakpointsCalculationPart"
	case Breakpoint:
		return "Breakpoint"
	case ByCharLevelFormulaCalculationPart:
		return "ByCharLevelFormulaCalculationPart"
	case CooldownMultiplierCalculationPart:
		return "CooldownMultiplierCalculationPart"
	}
	return ""
}

func FormulaPartTypeFromString(s string) CalculationPartType {
	switch s {
	case "IGameCalculationPart":
		return IGameCalculationPartType
	case "EffectValueCalculationPart":
		return EffectValueCalculationPart
	case "NamedDataValueCalculationPart":
		return NamedDataValueCalculationPart
	case "CustomReductionMultiplierCalculationPart":
		return CustomReductionMultiplierCalculationPart
	case "ProductOfSubPartsCalculationPart":
		return ProductOfSubPartsCalculationPart
	case "SumOfSubPartsCalculationPart":
		return SumOfSubPartsCalculationPart
	case "NumberCalculationPart":
		return NumberCalculationPart
	case "IGameCalculationPartWithStats":
		return IGameCalculationPartWithStats
	case "StatByCoefficientCalculationPart":
		return StatByCoefficientCalculationPart
	case "StatBySubPartCalculationPart":
		return StatBySubPartCalculationPart
	case "StatByNamedDataValueCalculationPart":
		return StatByNamedDataValueCalculationPart
	case "SubPartScaledProportionalToStat":
		return SubPartScaledProportionalToStat
	case "AbilityResourceByCoefficientCalculationPart":
		return AbilityResourceByCoefficientCalculationPart
	case "IGameCalculationPartWithBuffCounter":
		return IGameCalculationPartWithBuffCounter
	case "BuffCounterByCoefficientCalculationPart":
		return BuffCounterByCoefficientCalculationPart
	case "BuffCounterByNamedDataValueCalculationPart":
		return BuffCounterByNamedDataValueCalculationPart
	case "IGameCalculationPartByCharLevel":
		return IGameCalculationPartByCharLevel
	case "ByCharLevelInterpolationCalculationPart":
		return ByCharLevelInterpolationCalculationPart
	case "ByCharLevelBreakpointsCalculationPart":
		return ByCharLevelBreakpointsCalculationPart
	case "Breakpoint":
		return Breakpoint
	case "ByCharLevelFormulaCalculationPart":
		return ByCharLevelFormulaCalculationPart
	default:
		return UnknownPart
	}
}
