package parts

import (
	"fmt"
	"leaguebuilder/lol/data"
	"leaguebuilder/lol/datadragon"
	"strconv"
)

type NamedDataValueCalculationPart struct {
	DataValue *data.DataValue
}

func (p NamedDataValueCalculationPart) GetValue(ctx *data.CalculationContext) float64 {
	if p.DataValue == nil {
		return 0
	}
	return p.DataValue.Values[ctx.SpellRank]
}

func (p NamedDataValueCalculationPart) String(ctx *data.CalculationContext) string {
	if p.DataValue == nil || p.DataValue.Values == nil {
		return "?"
	}
	vs := strconv.FormatFloat(p.DataValue.Values[ctx.SpellRank], 'f', -1, 64)
	return fmt.Sprintf("%s", vs)
}

func (p *NamedDataValueCalculationPart) FromJson(m map[string]interface{}, ctx *data.ParseContext) {
	name := m["mDataValue"].(string)
	for s, value := range ctx.DataValues {
		if s == name || fmt.Sprintf("{%s}", datadragon.GetBinHash(s)) == name {
			p.DataValue = &value
		}
	}
}

func (p NamedDataValueCalculationPart) GetType() data.CalculationPartType {
	return data.NamedDataValueCalculationPart
}
