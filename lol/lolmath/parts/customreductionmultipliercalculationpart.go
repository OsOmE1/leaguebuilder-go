package parts

import (
	"fmt"
	"leaguebuilder/lol/data"
	"math"
	"strconv"
	"strings"
)

type CustomReductionMultiplierCalculationPart struct {
	MaximumReductionPercent float64
	Part                    *IGameCalculationPart
}

func (p CustomReductionMultiplierCalculationPart) GetValue(ctx *data.CalculationContext) float64 {
	if p.Part == nil {
		return 0
	}
	res := (*p.Part).GetValue(ctx)
	clamp := math.Max(math.Min(res, p.MaximumReductionPercent), 0)
	return 1 - clamp
}

func (p CustomReductionMultiplierCalculationPart) String(ctx *data.CalculationContext) string {
	v := p.GetValue(ctx)
	vs := strconv.FormatFloat(v, 'f', -1, 64)
	return fmt.Sprintf("%s", vs)
}

func (p *CustomReductionMultiplierCalculationPart) FromJson(m map[string]interface{}, ctx *data.ParseContext) {
	var v string
	for s, _ := range m {
		if strings.Contains(s, "{") {
			v = s
			break
		}
	}
	if v != "" {
		tp := PartFromJson(m[v].(map[string]interface{}), ctx)
		p.Part = &tp
	}
	p.MaximumReductionPercent = m["mMaximumReductionPercent"].(float64)
}
func (p CustomReductionMultiplierCalculationPart) GetType() data.CalculationPartType {
	return data.CustomReductionMultiplierCalculationPart
}
