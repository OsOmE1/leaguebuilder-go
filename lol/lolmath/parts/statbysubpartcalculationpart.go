package parts

import (
	"fmt"
	"leaguebuilder/lol/data"
	"math"
)

type StatBySubPartCalculationPart struct {
	Stat        data.StatType
	StatFormula data.StatFormulaType
	Subpart     *IGameCalculationPart
}

func (p StatBySubPartCalculationPart) GetValue(ctx *data.CalculationContext) float64 {
	if p.Subpart == nil {
		return 0
	}
	return (*p.Subpart).GetValue(ctx)
}

func (p StatBySubPartCalculationPart) String(ctx *data.CalculationContext) string {
	return fmt.Sprintf("%d%% %s", int(math.Round((*p.Subpart).GetValue(ctx))), p.Stat)
}

func (p *StatBySubPartCalculationPart) FromJson(m map[string]interface{}, ctx *data.ParseContext) {
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
}

func (p StatBySubPartCalculationPart) GetType() data.CalculationPartType {
	return data.StatBySubPartCalculationPart
}
