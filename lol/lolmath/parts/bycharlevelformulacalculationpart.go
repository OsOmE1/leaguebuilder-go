package parts

import (
	"fmt"
	"leaguebuilder/lol/data"
	"strconv"
)

type ByCharLevelFormulaCalculationPart struct {
	Values []float64
}

func (p ByCharLevelFormulaCalculationPart) GetValue(ctx *data.CalculationContext) float64 {
	return p.Values[ctx.CharacterLevel]
}

func (p ByCharLevelFormulaCalculationPart) String(ctx *data.CalculationContext) string {
	vs := strconv.FormatFloat(p.GetValue(ctx), 'f', -1, 64)
	return fmt.Sprintf("%s", vs)
}

func (p *ByCharLevelFormulaCalculationPart) FromJson(m map[string]interface{}, ctx *data.ParseContext) {
	for _, v := range m["mValues"].([]interface{}) {
		p.Values = append(p.Values, v.(float64))
	}
}

func (p ByCharLevelFormulaCalculationPart) GetType() data.CalculationPartType {
	return data.ByCharLevelFormulaCalculationPart
}
