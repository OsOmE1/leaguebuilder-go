package parts

import (
	"fmt"
	"leaguebuilder/lol/data"
	"strconv"
)

type ByCharLevelInterpolationCalculationPart struct {
	StartValue float64
	EndValue   float64
	Bool1      bool
	Bool2      bool
}

func (p ByCharLevelInterpolationCalculationPart) GetValue(ctx *data.CalculationContext) float64 {
	return p.StartValue + (p.EndValue-p.StartValue)*(float64(ctx.CharacterLevel)/18)
}

func (p ByCharLevelInterpolationCalculationPart) String(ctx *data.CalculationContext) string {
	vs := strconv.FormatFloat(p.GetValue(ctx), 'f', -1, 64)
	return fmt.Sprintf("%s", vs)
}

func (p *ByCharLevelInterpolationCalculationPart) FromJson(m map[string]interface{}, ctx *data.ParseContext) {
	if val, ok := m["mStartValue"]; ok {
		p.StartValue = val.(float64)
	}
	if val, ok := m["mEndValue"]; ok {
		p.EndValue = val.(float64)
	}
	set := false
	for s, i := range m {
		if s == "mStartValue" || s == "mEndValue" || s == "__type" {
			continue
		}
		if val, ok := i.(bool); ok {
			if set {
				p.Bool2 = val
				break
			}
			p.Bool1 = val
			set = true
		}
	}
}

func (p ByCharLevelInterpolationCalculationPart) GetType() data.CalculationPartType {
	return data.ByCharLevelInterpolationCalculationPart
}
