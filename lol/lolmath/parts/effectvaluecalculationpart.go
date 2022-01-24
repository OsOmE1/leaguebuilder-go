package parts

import (
	"fmt"
	"leaguebuilder/lol/data"
	"strconv"
)

type EffectValueCalculationPart struct {
	EffectIndex int
}

func (p EffectValueCalculationPart) GetValue(ctx *data.CalculationContext) float64 {
	if ctx.EffectValues == nil {
		return 0
	}
	return ctx.EffectValues[p.EffectIndex-1][ctx.SpellRank]
}
func (p EffectValueCalculationPart) String(ctx *data.CalculationContext) string {
	v := p.GetValue(ctx)
	vs := strconv.FormatFloat(v, 'f', -1, 64)
	return fmt.Sprintf("%s", vs)
}

func (p *EffectValueCalculationPart) FromJson(m map[string]interface{}, ctx *data.ParseContext) {
	p.EffectIndex = int(m["mEffectIndex"].(float64))
}

func (p EffectValueCalculationPart) GetType() data.CalculationPartType {
	return data.EffectValueCalculationPart
}
