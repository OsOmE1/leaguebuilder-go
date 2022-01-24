package parts

import (
	"fmt"
	"leaguebuilder/lol/data"
	"strconv"
)

type NumberCalculationPart struct {
	Number float64
}

func (p NumberCalculationPart) GetValue(ctx *data.CalculationContext) float64 {
	return p.Number
}

func (p NumberCalculationPart) String(ctx *data.CalculationContext) string {
	vs := strconv.FormatFloat(p.Number, 'f', -1, 64)
	return fmt.Sprintf("%s", vs)
}

func (p *NumberCalculationPart) FromJson(m map[string]interface{}, ctx *data.ParseContext) {
	if val, ok := m["mNumber"]; ok {
		p.Number = val.(float64)
	}
}

func (p NumberCalculationPart) GetType() data.CalculationPartType {
	return data.NumberCalculationPart
}
