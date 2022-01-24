package lol

import (
	"fmt"
	"leaguebuilder/lol/data"
	"leaguebuilder/lol/datadragon"
	"leaguebuilder/lol/lolmath/calc"
	"strconv"
	"strings"
)

type Spell struct {
	Name    string
	Tooltip string

	DamageType             []data.DamageType
	SpellTags              []string
	EffectAmounts          [][]float64
	DataValues             map[string]data.DataValue
	SpellCalculations      []calc.IGameCalculation
	Coefficient            float64
	CastTime               float64
	CooldownTime           []float64
	DelayCastOffsetPercent float64
	DelayTotalTimePercent  float64
	CastRange              []float64
	CastRadius             []float64

	CC        []string
	PartialCC []string
}

type SpellInstance struct {
	Spell *Spell
	Rank  int
}

type Ratio struct {
	Stat   data.StatType
	Amount float64
}

type Effect struct {
	Damage data.DamageType
	Effect data.EffectType
	Amount float64
	Ratios []Ratio
}

func (s *Spell) GetRatios() (ratios []string) {
	groups := datadragon.GetGroups(`<(\w*)>`, s.Tooltip)
	groups = datadragon.RemoveDuplicateStr(groups)
	for _, group := range groups {
		_, valid := data.DmgMap[group]
		//if !valid {
		//	_, valid = data.EffMap[group]
		//}

		if !valid {
			continue
		}

		gs := datadragon.GetGroups(fmt.Sprintf(`<%s>(.*?)<\/%s>`, group, group), s.Tooltip)
		for _, g := range gs {
			vars := datadragon.GetGroups(`@(.*?)@`, g)
			for _, v := range vars {
				g = strings.Replace(g, fmt.Sprintf("@%s@", v), "", -1)
			}
			g = strings.TrimSpace(g)
			ratios = append(ratios, fmt.Sprintf("%s - %s", group, g))
		}
	}
	return ratios
}

func (s *Spell) GetDataValue(name string) *data.DataValue {
	if val, ok := s.DataValues[name]; ok {
		return &val
	}
	if val, ok := s.DataValues[fmt.Sprintf("{%s}", datadragon.GetBinHash(name))]; ok {
		return &val
	}
	return nil
}

func SpellFromApiSpell(s datadragon.Spell, fc *datadragon.FontConfig) (Spell, error) {
	spell := Spell{
		Name:                   s.MScriptName,
		SpellTags:              s.MSpell.MSpellTags,
		Coefficient:            s.MSpell.MCoefficient,
		CastTime:               s.MSpell.MCastTime,
		CooldownTime:           s.MSpell.CooldownTime,
		DelayCastOffsetPercent: s.MSpell.DelayCastOffsetPercent,
		DelayTotalTimePercent:  s.MSpell.DelayTotalTimePercent,
		CastRange:              s.MSpell.CastRange,
		CastRadius:             s.MSpell.CastRadius,
	}
	tt, err := fc.Get(strings.ToLower(s.MSpell.MClientData.MTooltipData.MLocKeys.KeyTooltip))
	if err != nil {
		fmt.Println(err)
		tt = "not found"
	}
	spell.Tooltip = tt

	if strings.Contains(spell.Tooltip, "physicalDamage") {
		spell.DamageType = append(spell.DamageType, data.Physical)
	}
	if strings.Contains(spell.Tooltip, "magicDamage") {
		spell.DamageType = append(spell.DamageType, data.Magic)
	}
	if strings.Contains(spell.Tooltip, "trueDamage") {
		spell.DamageType = append(spell.DamageType, data.True)
	}

	for _, cc := range data.HardCC {
		if strings.Contains(spell.Tooltip, fmt.Sprintf("<status>%s", cc)) {
			spell.CC = append(spell.CC, cc)
		}
	}
	for _, cc := range data.PartialCC {
		if strings.Contains(spell.Tooltip, fmt.Sprintf("<status>%s", cc)) {
			spell.PartialCC = append(spell.PartialCC, cc)
		}
	}

	for _, vs := range s.MSpell.MEffectAmount {
		spell.EffectAmounts = append(spell.EffectAmounts, vs.Value)
	}
	spell.DataValues = map[string]data.DataValue{}
	for _, value := range s.MSpell.MDataValues {
		dv := data.DataValue{
			Values:    value.MValues,
			BaseValue: value.MBaseP,
		}
		if value.MFormula == "" || value.MFormula == "P" {
			spell.DataValues[value.MName] = dv
			continue
		}

		dv.Formula = value.MFormula
		spell.DataValues[value.MName] = dv
	}
	pctx := data.ParseContext{
		DataValues: spell.DataValues,
		ApiSpell:   &s,
	}
	for name, i := range s.MSpell.MSpellCalculations {
		spcalc := i.(map[string]interface{})
		t := data.CalculationTypeFromString(spcalc["__type"].(string))
		var gc calc.IGameCalculation
		switch t {
		case data.GameCalculationType:
			gc = &calc.GameCalculation{}
		case data.GameCalculationModifiedType:
			gc = &calc.GameCalculationModified{}
		}
		if gc == nil {
			continue // TODO: Fix this
		}
		gc.FromJson(name, spcalc, &pctx)
		spell.SpellCalculations = append(spell.SpellCalculations, gc)
	}
	return spell, nil
}

func (s *Spell) ResolveSpellText(champ *ChampionInstance, rank int) string {
	tt := s.Tooltip
	tt = strings.Replace(tt, ".0*", "*", -1)
	for _, calculation := range s.SpellCalculations {
		ctx := champ.CalculationContext(s.Instance(rank))

		tt = ttStringReplace(tt, calculation.Name(), calculation.String(ctx))
	}
	for name, value := range s.DataValues {
		if value.Values == nil {
			continue
		}
		tt = ttStringReplace(tt, name, fmt.Sprintf("%.1f", value.Values[rank]))
	}
	for i, value := range s.EffectAmounts {
		if value == nil {
			continue
		}
		name := fmt.Sprintf("Effect%dAmount", i+1)
		tt = ttStringReplace(tt, name, fmt.Sprintf("%.1f", value[rank]))
	}
	return tt
}

func (s *Spell) Instance(level int) *SpellInstance {
	return &SpellInstance{
		Spell: s,
		Rank:  level,
	}
}

func ttStringReplace(s, name, value string) string {
	name = strings.Replace(name, "{", "\\{", -1)
	name = strings.Replace(name, "}", "\\}", -1)

	groups := datadragon.GetGroups(fmt.Sprintf(`@%s\*([-\.0-9]*)@`, name), s)
	if len(groups) > 0 {
		mult := groups[0]
		if mult[0] == '.' {
			mult = "0" + mult
		}
		f, _ := strconv.ParseFloat(mult, 64)
		v, err := strconv.ParseFloat(value, 64)
		vs := value
		if err == nil {
			v = f * v
			vs = strconv.FormatFloat(v, 'f', -1, 64)
		}
		if err != nil {
			return strings.Replace(s, fmt.Sprintf("@%s*%s@", name, groups[0]), fmt.Sprintf("(%s*%s)", vs, groups[0]), -1)
		}
		return strings.Replace(s, fmt.Sprintf("@%s*%s@", name, groups[0]), vs, -1)
	}
	v, err := strconv.ParseFloat(value, 64)
	if err != nil {
		return strings.Replace(s, fmt.Sprintf("@%s@", name), value, -1)
	}
	vs := strconv.FormatFloat(v, 'f', -1, 64)
	return strings.Replace(s, fmt.Sprintf("@%s@", name), vs, 10)
}
