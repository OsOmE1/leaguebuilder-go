package calc

import (
	"fmt"
	"leaguebuilder/lol/data"
	"leaguebuilder/lol/lolmath/parts"
	"math"
	"strconv"
)

type GameCalculation struct {
	name             string
	FormulaParts     []parts.IGameCalculationPart
	DisplayAsPercent bool
	Precision        int

	Multiplier *parts.IGameCalculationPart
}

func (gc GameCalculation) Name() string {
	return gc.name
}

func (gc GameCalculation) Value(ctx *data.CalculationContext) float64 {
	if gc.Multiplier != nil {
		ctx.Multiplier *= (*gc.Multiplier).GetValue(ctx)
	}
	res := 0.0
	for _, part := range gc.FormulaParts {
		res += part.GetValue(ctx) * ctx.Multiplier
	}
	return 0
}

func (gc GameCalculation) String(ctx *data.CalculationContext) string {
	str := ""
	for i, part := range gc.FormulaParts {
		if part == nil {
			continue
		}
		str += part.String(ctx)
		if i != len(gc.FormulaParts)-1 {
			str += " + "
		}
	}
	if gc.DisplayAsPercent {
		v, err := strconv.ParseFloat(str, 64)
		if err != nil {
			return str
		}
		v = v * 100
		vs := strconv.FormatFloat(math.Round(v), 'f', -1, 64)
		return fmt.Sprintf("%s%%", vs)
	}
	return str
}

func (gc *GameCalculation) FromJson(name string, m map[string]interface{}, ctx *data.ParseContext) {
	gc.name = name

	ps := m["mFormulaParts"].([]interface{})
	for _, p := range ps {
		gc.FormulaParts = append(gc.FormulaParts, parts.PartFromJson(p.(map[string]interface{}), ctx))
	}
	if _, ok := m["mMultiplier"]; ok {
		tp := parts.PartFromJson(m["mMultiplier"].(map[string]interface{}), ctx)
		gc.Multiplier = &tp
	}
	if _, ok := m["mDisplayAsPercent"]; ok {
		gc.DisplayAsPercent = m["mDisplayAsPercent"].(bool)
	}
	if _, ok := m["mPrecision"]; ok {
		gc.Precision = int(m["mPrecision"].(float64))
	}
}
func (gc GameCalculation) Type() data.CalculationType {
	return data.GameCalculationType
}
