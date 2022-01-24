package parts

import (
	"leaguebuilder/lol/data"
)

type CooldownMultiplierCalculationPart struct {
}

func (p CooldownMultiplierCalculationPart) GetValue(ctx *data.CalculationContext) float64 {
	return 0
}
func (p CooldownMultiplierCalculationPart) String(ctx *data.CalculationContext) string {
	return ""
}
func (p *CooldownMultiplierCalculationPart) FromJson(m map[string]interface{}, ctx *data.ParseContext) {
}
func (p CooldownMultiplierCalculationPart) GetType() data.CalculationPartType {
	return data.CooldownMultiplierCalculationPart
}
