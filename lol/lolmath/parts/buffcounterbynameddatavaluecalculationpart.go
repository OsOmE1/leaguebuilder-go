package parts

import (
	"fmt"
	"leaguebuilder/lol/data"
	"leaguebuilder/lol/datadragon"
	"strconv"
)

type BuffCounterByNamedDataValueCalculationPart struct {
	BuffName  string
	DataValue *data.DataValue
}

func (p BuffCounterByNamedDataValueCalculationPart) GetValue(ctx *data.CalculationContext) float64 {
	if p.DataValue == nil {
		return 0
	}
	return p.DataValue.Values[ctx.SpellRank]
}

func (p BuffCounterByNamedDataValueCalculationPart) String(ctx *data.CalculationContext) string {
	vs := strconv.FormatFloat(p.DataValue.Values[ctx.SpellRank], 'f', -1, 64)

	return fmt.Sprintf("%s %s", vs, p.BuffName)
}

func (p *BuffCounterByNamedDataValueCalculationPart) FromJson(m map[string]interface{}, ctx *data.ParseContext) {
	if val, ok := m["mBuffName"]; ok {
		p.BuffName = val.(string)
	}
	name := m["mDataValue"].(string)
	for s, value := range ctx.DataValues {
		if s == name || fmt.Sprintf("{%s}", datadragon.GetBinHash(s)) == name {
			p.DataValue = &value
		}
	}
}

func (p BuffCounterByNamedDataValueCalculationPart) GetType() data.CalculationPartType {
	return data.BuffCounterByNamedDataValueCalculationPart
}
