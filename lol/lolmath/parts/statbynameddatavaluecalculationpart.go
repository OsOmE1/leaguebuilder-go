package parts

import (
	"fmt"
	"leaguebuilder/lol/data"
	"leaguebuilder/lol/datadragon"
	"math"
)

type StatByNamedDataValueCalculationPart struct {
	Stat        data.StatType
	StatFormula data.StatFormulaType
	DataValue   *data.DataValue
}

func (p StatByNamedDataValueCalculationPart) GetValue(ctx *data.CalculationContext) float64 {
	if p.DataValue.Values == nil {
		return 0
	}
	return p.DataValue.Values[ctx.SpellRank]
}

func (p StatByNamedDataValueCalculationPart) String(ctx *data.CalculationContext) string {
	if p.DataValue == nil || p.DataValue.Values == nil {
		return "?"
	}
	return fmt.Sprintf("%d%% %s", int(math.Round(p.DataValue.Values[ctx.SpellRank]*100)), p.Stat)
}

func (p *StatByNamedDataValueCalculationPart) FromJson(m map[string]interface{}, ctx *data.ParseContext) {
	if val, ok := m["mStat"]; ok {
		p.Stat = data.StatType(val.(float64))
	}
	if val, ok := m["mStatFormula"]; ok {
		p.StatFormula = data.StatFormulaType(val.(float64))
	}
	name := m["mDataValue"].(string)
	for s, value := range ctx.DataValues {
		if s == name || fmt.Sprintf("{%s}", datadragon.GetBinHash(s)) == name {
			p.DataValue = &value
		}
	}
}

func (p StatByNamedDataValueCalculationPart) GetType() data.CalculationPartType {
	return data.StatByNamedDataValueCalculationPart
}
