package parts

import (
	"fmt"
	"leaguebuilder/lol/data"
	"math"
)

type AbilityResourceByCoefficientCalculationPart struct {
	Coefficient     float64
	AbilityResource data.AbilityResourceType
	StatFormula     data.StatFormulaType
}

func (p AbilityResourceByCoefficientCalculationPart) GetValue(ctx *data.CalculationContext) float64 {
	fmt.Println("ability resource values not implemented")
	return p.Coefficient * 0 // ctx.Champion.GetAbilityResourceValue(p.AbilityResource, p.StatFormula)
}

func (p AbilityResourceByCoefficientCalculationPart) String(ctx *data.CalculationContext) string {
	return fmt.Sprintf("%d%% %s", int(math.Round(p.Coefficient*100)), p.AbilityResource)
}

func (p *AbilityResourceByCoefficientCalculationPart) FromJson(m map[string]interface{}, ctx *data.ParseContext) {
	if val, ok := m["mCoefficient"]; ok {
		p.Coefficient = val.(float64)
	}
	if val, ok := m["mStatFormula"]; ok {
		p.StatFormula = data.StatFormulaType(val.(float64))
	}
	if val, ok := m["mAbilityResource"]; ok {
		p.AbilityResource = data.AbilityResourceType(val.(float64))
	}
}

func (p AbilityResourceByCoefficientCalculationPart) GetType() data.CalculationPartType {
	return data.AbilityResourceByCoefficientCalculationPart
}
