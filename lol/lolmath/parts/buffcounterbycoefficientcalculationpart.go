package parts

import (
	"fmt"
	"leaguebuilder/lol/data"
	"math"
)

type BuffCounterByCoefficientCalculationPart struct {
	BuffName    string
	Coefficient float64
}

func (p BuffCounterByCoefficientCalculationPart) GetValue(ctx *data.CalculationContext) float64 {
	return p.Coefficient
}

func (p BuffCounterByCoefficientCalculationPart) String(ctx *data.CalculationContext) string {
	return fmt.Sprintf("%d%% %s", int(math.Round(p.Coefficient*100)), p.BuffName)
}

func (p *BuffCounterByCoefficientCalculationPart) FromJson(m map[string]interface{}, ctx *data.ParseContext) {
	if val, ok := m["mBuffName"]; ok {
		p.BuffName = val.(string)
	}
	if val, ok := m["mCoefficient"]; ok {
		p.Coefficient = val.(float64)
	}
}

func (p BuffCounterByCoefficientCalculationPart) GetType() data.CalculationPartType {
	return data.BuffCounterByCoefficientCalculationPart
}
