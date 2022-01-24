package parts

import (
	"fmt"
	"leaguebuilder/lol/data"
	"strconv"
)

type SubPartScaledProportionalToStat struct {
	Stat        data.StatType
	StatFormula data.StatFormulaType
	Subpart     *IGameCalculationPart
	Ratio       float64
}

func (p SubPartScaledProportionalToStat) GetValue(ctx *data.CalculationContext) float64 {
	if p.Subpart == nil {
		return 0
	}
	res := (*p.Subpart).GetValue(ctx)
	return ((ctx.Stats.Values[p.Stat] * p.Ratio) + 1) * res // TODO fix statformula implementation
}

func (p SubPartScaledProportionalToStat) String(ctx *data.CalculationContext) string {
	vs := strconv.FormatFloat(p.GetValue(ctx), 'f', -1, 64)
	return fmt.Sprintf("%s", vs)
}

func (p *SubPartScaledProportionalToStat) FromJson(m map[string]interface{}, ctx *data.ParseContext) {
	if val, ok := m["mStat"]; ok {
		p.Stat = data.StatType(val.(float64))
	}
	if val, ok := m["mStatFormula"]; ok {
		p.StatFormula = data.StatFormulaType(val.(float64))
	}
	if val, ok := m["mSubpart"]; ok {
		tp := PartFromJson(val.(map[string]interface{}), ctx)
		p.Subpart = &tp
	}
	if val, ok := m["mRatio"]; ok {
		p.Ratio = val.(float64)
	}
}

func (p SubPartScaledProportionalToStat) GetType() data.CalculationPartType {
	return data.SubPartScaledProportionalToStat
}
