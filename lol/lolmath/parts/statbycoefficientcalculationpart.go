package parts

import (
	"fmt"
	"leaguebuilder/lol/data"
	"math"
)

type StatByCoefficientCalculationPart struct {
	Stat        data.StatType
	StatFormula data.StatFormulaType
	Coefficient float64
}

func (p StatByCoefficientCalculationPart) GetValue(ctx *data.CalculationContext) float64 {
	return p.Coefficient
}

func (p StatByCoefficientCalculationPart) String(ctx *data.CalculationContext) string {
	return fmt.Sprintf("%d%% %s", int(math.Round(p.Coefficient*100)), p.Stat)
}

func (p *StatByCoefficientCalculationPart) FromJson(m map[string]interface{}, ctx *data.ParseContext) {
	if val, ok := m["mStat"]; ok {
		p.Stat = data.StatType(val.(float64))
	}
	if val, ok := m["mStatFormula"]; ok {
		p.StatFormula = data.StatFormulaType(val.(float64))
	}
	p.Coefficient = m["mCoefficient"].(float64)
}

func (p StatByCoefficientCalculationPart) GetType() data.CalculationPartType {
	return data.StatByCoefficientCalculationPart
}
