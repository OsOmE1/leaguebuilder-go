package parts

import (
	"leaguebuilder/lol/data"
)

type IGameCalculationPart interface {
	GetValue(*data.CalculationContext) float64
	String(*data.CalculationContext) string
	FromJson(map[string]interface{}, *data.ParseContext)
	GetType() data.CalculationPartType
}

type GameCalculationPart struct {
}

func (gp GameCalculationPart) GetValue(ctx *data.CalculationContext) float64 {
	return 0
}
func (gp GameCalculationPart) String(ctx *data.CalculationContext) string {
	return ""
}
func (gp GameCalculationPart) GetType() data.CalculationPartType {
	return data.UnknownPart
}
func (gp GameCalculationPart) FromJson(m map[string]interface{}, ctx *data.ParseContext) {}

func PartFromJson(m map[string]interface{}, ctx *data.ParseContext) IGameCalculationPart {
	t := data.FormulaPartTypeFromString(m["__type"].(string))
	var p IGameCalculationPart
	switch t {
	case data.UnknownPart:
		return p
	case data.EffectValueCalculationPart:
		p = &EffectValueCalculationPart{}
		break
	case data.NamedDataValueCalculationPart:
		p = &NamedDataValueCalculationPart{}
		break
	case data.CustomReductionMultiplierCalculationPart:
		p = &CustomReductionMultiplierCalculationPart{}
		break
	case data.ProductOfSubPartsCalculationPart:
		p = &ProductOfSubPartsCalculationPart{}
		break
	case data.SumOfSubPartsCalculationPart:
		p = &SumOfSubPartsCalculationPart{}
		break
	case data.NumberCalculationPart:
		p = &NumberCalculationPart{}
		break
	case data.StatByCoefficientCalculationPart:
		p = &StatByCoefficientCalculationPart{}
		break
	case data.StatBySubPartCalculationPart:
		p = &StatBySubPartCalculationPart{}
		break
	case data.StatByNamedDataValueCalculationPart:
		p = &StatByNamedDataValueCalculationPart{}
		break
	case data.SubPartScaledProportionalToStat:
		p = &SubPartScaledProportionalToStat{}
		break
	case data.AbilityResourceByCoefficientCalculationPart:
		p = &AbilityResourceByCoefficientCalculationPart{}
		break
	case data.BuffCounterByCoefficientCalculationPart:
		p = &BuffCounterByCoefficientCalculationPart{}
		break
	case data.BuffCounterByNamedDataValueCalculationPart:
		p = &BuffCounterByNamedDataValueCalculationPart{}
		break
	case data.ByCharLevelInterpolationCalculationPart:
		p = &ByCharLevelInterpolationCalculationPart{}
		break
	case data.ByCharLevelBreakpointsCalculationPart:
		p = &ByCharLevelBreakpointsCalculationPart{}
		break
	case data.ByCharLevelFormulaCalculationPart:
		p = &ByCharLevelFormulaCalculationPart{}
		break
	case data.CooldownMultiplierCalculationPart:
		p = &CooldownMultiplierCalculationPart{}
		break
	}
	p.FromJson(m, ctx)
	return p
}
