package data

import (
	"leaguebuilder/lol/datadragon"
)

type CalculationContext struct {
	CharacterLevel int
	Stats          Stats
	DataValues     map[string]DataValue
	EffectValues   [][]float64
	SpellRank      int
	Multiplier     float64
}

type ParseContext struct {
	ApiSpell   *datadragon.Spell
	DataValues map[string]DataValue
}
