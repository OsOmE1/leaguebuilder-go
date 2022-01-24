package parts

import (
	"leaguebuilder/lol/data"
)

type SumOfSubPartsCalculationPart struct {
	SubParts []IGameCalculationPart
}

func (p SumOfSubPartsCalculationPart) GetValue(ctx *data.CalculationContext) float64 {
	res := 0.0
	for _, part := range p.SubParts {
		res += part.GetValue(ctx)
	}
	return res
}

func (p SumOfSubPartsCalculationPart) String(ctx *data.CalculationContext) string {
	str := ""
	for i, part := range p.SubParts {
		str += part.String(ctx)
		if i != len(p.SubParts)-1 {
			str += " + "
		}
	}
	return str
}

func (p *SumOfSubPartsCalculationPart) FromJson(m map[string]interface{}, ctx *data.ParseContext) {
	if _, ok := m["mSubpart"].(map[string]map[string]interface{}); ok {
		for _, i := range m["mSubpart"].(map[string]map[string]interface{}) {
			p.SubParts = append(p.SubParts, PartFromJson(i, ctx))
		}
		return
	}
	if val, ok := m["mSubparts"].([]interface{}); ok {
		for _, i := range val {
			p.SubParts = append(p.SubParts, PartFromJson(i.(map[string]interface{}), ctx))
		}
		return
	}
	p.SubParts = append(p.SubParts, PartFromJson(m["mSubpart"].(map[string]interface{}), ctx))
}

func (p SumOfSubPartsCalculationPart) GetType() data.CalculationPartType {
	return data.SumOfSubPartsCalculationPart
}
