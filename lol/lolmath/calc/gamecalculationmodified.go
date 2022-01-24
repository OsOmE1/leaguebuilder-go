package calc

import (
	"leaguebuilder/lol/data"
	"leaguebuilder/lol/lolmath/parts"
)

type GameCalculationModified struct {
	name                    string
	Multiplier              *parts.IGameCalculationPart
	OverrideSpellLevel      int
	ModifiedGameCalculation *GameCalculation
}

func (gc GameCalculationModified) Name() string {
	return gc.name
}

func (gc GameCalculationModified) Value(ctx *data.CalculationContext) float64 {
	if gc.Multiplier == nil {
		return 0
	}
	return gc.ModifiedGameCalculation.Value(ctx) * (*gc.Multiplier).GetValue(ctx)
}
func (gc GameCalculationModified) String(ctx *data.CalculationContext) string {
	if gc.Multiplier == nil {
		return "?"
	}
	return gc.ModifiedGameCalculation.String(ctx) + " * " + (*gc.Multiplier).String(ctx)
}
func (gc *GameCalculationModified) FromJson(name string, m map[string]interface{}, ctx *data.ParseContext) {
	gc.name = name

	name = m["mModifiedGameCalculation"].(string)
	if _, ok := m["mMultiplier"]; ok {
		tp := parts.PartFromJson(m["mMultiplier"].(map[string]interface{}), ctx)
		gc.Multiplier = &tp
	}
	if i, ok := ctx.ApiSpell.MSpell.MSpellCalculations[name]; ok {
		gcal := GameCalculation{}
		gcal.FromJson(name, i.(map[string]interface{}), ctx)
		gc.ModifiedGameCalculation = &gcal
	}
	gc.OverrideSpellLevel = -1
	if _, ok := m["mOverrideSpellLevel"]; ok {
		gc.OverrideSpellLevel = int(m["mOverrideSpellLevel"].(float64))
	}
}
func (gc GameCalculationModified) Type() data.CalculationType {
	return data.GameCalculationModifiedType
}
