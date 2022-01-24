package parts

import (
	"fmt"
	"leaguebuilder/lol/data"
	"strconv"
)

type ByCharLevelBreakpointsCalculationPart struct {
	Level1Value float64
	Scale       float64
	BreakPoints map[int]float64
}

func (p ByCharLevelBreakpointsCalculationPart) GetValue(ctx *data.CalculationContext) float64 {
	v := p.Level1Value
	perLevel := p.Scale
	for i := 1; i < ctx.CharacterLevel; i++ {
		if new, ok := p.BreakPoints[i+1]; ok {
			perLevel = new
		}
		v += perLevel
	}
	return v
}

func (p ByCharLevelBreakpointsCalculationPart) String(ctx *data.CalculationContext) string {
	vs := strconv.FormatFloat(p.GetValue(ctx), 'f', -1, 64)
	return fmt.Sprintf("%s", vs)
}

func (p *ByCharLevelBreakpointsCalculationPart) FromJson(m map[string]interface{}, ctx *data.ParseContext) {
	if val, ok := m["mLevel1Value"]; ok {
		p.Level1Value = val.(float64)
	}
	for s, i := range m {
		if s == "mLevel1Value" || s == "mBreakpoints" || s == "__type" {
			continue
		}
		if val, ok := i.(float64); ok {
			p.Scale = val
		}
	}
	if _, ok := m["mBreakpoints"]; !ok {
		return
	}
	breakpoints := m["mBreakpoints"].([]interface{})
	p.BreakPoints = map[int]float64{}

	for _, b := range breakpoints {
		breakpoint := b.(map[string]interface{})
		v := ""
		for s, _ := range breakpoint {
			if s == "__type" || s == "mLevel" {
				continue
			}
			v = s
			break
		}
		if _, ok := breakpoint["mLevel"]; !ok {
			if v != "" {
				p.Scale = breakpoint[v].(float64)
			}
			break
		}
		level := int(breakpoint["mLevel"].(float64))
		if v != "" {
			p.BreakPoints[level] = breakpoint[v].(float64)
			break
		}
		p.BreakPoints[level] = 0
	}
}

func (p ByCharLevelBreakpointsCalculationPart) GetType() data.CalculationPartType {
	return data.ByCharLevelBreakpointsCalculationPart
}
