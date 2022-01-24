package data

import "leaguebuilder/lol/datadragon"

type Stats struct {
	AbilityResource struct {
		Type  AbilityResourceType
		Value float64
	}
	Values map[StatType]float64
}

func StatsFromCharacterRecord(record datadragon.CharacterRecord) Stats {
	return Stats{
		AbilityResource: struct {
			Type  AbilityResourceType
			Value float64
		}{},
		Values: map[StatType]float64{
			AbilityPower:                 0,
			Armor:                        record.BaseArmor,
			Attack:                       record.BaseDamage,
			AttackSpeed:                  record.AttackSpeed,
			AttackWindupTime:             0,
			MagicResist:                  record.BaseSpellBlock,
			MoveSpeed:                    record.BaseMoveSpeed,
			CritChance:                   0,
			CritDamage:                   0,
			CooldownReduction:            0,
			AbilityHaste:                 0,
			MaxHealth:                    record.BaseHP,
			CurrentHealth:                record.BaseHP,
			PercentMissingHealth:         0,
			LifeSteal:                    0,
			OmniVamp:                     0,
			PhysicalVamp:                 0,
			MagicPenetrationFlat:         0,
			MagicPenetrationPercent:      0,
			BonusMagicPenetrationPercent: 0,
			MagicLethality:               0,
			ArmorPenetrationFlat:         0,
			ArmorPenetrationPercent:      0,
			BonusArmorPenetrationPercent: 0,
			PhysicalLethality:            0,
			Tenacity:                     0,
			AttackRange:                  record.AttackRange,
			HealthRegenRate:              record.BaseStaticHPRegen,
			ResourceRegenRate:            0,
			DodgeChance:                  0,
		},
	}
}
