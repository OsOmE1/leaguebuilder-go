package calc

import (
	"leaguebuilder/lol/data"
)

type IGameCalculation interface {
	Name() string
	Value(*data.CalculationContext) float64
	String(*data.CalculationContext) string
	FromJson(string, map[string]interface{}, *data.ParseContext)
	Type() data.CalculationType
}
