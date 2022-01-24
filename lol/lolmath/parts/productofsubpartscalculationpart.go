package parts

import (
	"fmt"
	"leaguebuilder/lol/data"
)

type ProductOfSubPartsCalculationPart struct {
	Part1 *IGameCalculationPart
	Part2 *IGameCalculationPart
}

func (p ProductOfSubPartsCalculationPart) GetValue(ctx *data.CalculationContext) float64 {
	if (*p.Part1) == nil || (*p.Part2) == nil {
		return 0
	}
	return (*p.Part1).GetValue(ctx) + (*p.Part2).GetValue(ctx)
}

func (p ProductOfSubPartsCalculationPart) String(ctx *data.CalculationContext) string {
	if (*p.Part1) == nil || (*p.Part2) == nil {
		return "?"
	}
	return fmt.Sprintf("%s + %s", (*p.Part1).String(ctx), (*p.Part2).String(ctx))
}

func (p *ProductOfSubPartsCalculationPart) FromJson(m map[string]interface{}, ctx *data.ParseContext) {
	if val, ok := m["mPart1"]; ok {
		tp := PartFromJson(val.(map[string]interface{}), ctx)
		p.Part1 = &tp
	}
	if val, ok := m["mPart2"]; ok {
		tp := PartFromJson(val.(map[string]interface{}), ctx)
		p.Part2 = &tp
	}
}

func (p ProductOfSubPartsCalculationPart) GetType() data.CalculationPartType {
	return data.ProductOfSubPartsCalculationPart
}
