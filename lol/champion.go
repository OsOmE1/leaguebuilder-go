package lol

import (
	"leaguebuilder/lol/data"
	"leaguebuilder/lol/datadragon"
)

type Champion struct {
	Name string
	Tags []string

	AbilityResource data.AbilityResourceType
	BaseStats       data.Stats
	Spells          []Spell
}

func ChampionFromApi(c *datadragon.Champion, fc *datadragon.FontConfig) (Champion, error) {
	var spells []Spell
	for _, s := range c.Spells {
		spell, err := SpellFromApiSpell(s, fc)
		if err != nil {
			return Champion{}, err
		}
		spells = append(spells, spell)
	}
	return Champion{
		Name:            c.Character.Name,
		Tags:            []string{c.Character.CharacterToolData.SearchTags, c.Character.CharacterToolData.SearchTagsSecondary},
		AbilityResource: 0,
		BaseStats:       data.StatsFromCharacterRecord(c.Character),
		Spells:          spells,
	}, nil
}

type ChampionInstance struct {
	Champion *Champion
	Stats    data.Stats
	Items    []struct{}
	Level    int
	Buffs    []struct{}
}

func (c *Champion) Instance(level int) *ChampionInstance {
	return &ChampionInstance{
		Champion: c,
		Stats:    c.BaseStats,
		Items:    []struct{}{},
		Level:    level,
		Buffs:    []struct{}{},
	}
}

func (c *ChampionInstance) CalculationContext(si *SpellInstance) *data.CalculationContext {
	return &data.CalculationContext{
		CharacterLevel: c.Level,
		Stats:          data.Stats{},
		DataValues:     si.Spell.DataValues,
		EffectValues:   si.Spell.EffectAmounts,
		SpellRank:      si.Rank,
		Multiplier:     1.0,
	}
}
