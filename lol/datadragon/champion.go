package datadragon

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"sort"
	"strings"
)

type AbilityResourceType int

const (
	Mana           AbilityResourceType = 0
	Energy         AbilityResourceType = 1
	TryndamereFury AbilityResourceType = 4
	RenektonFury   AbilityResourceType = 6
	Rage           AbilityResourceType = 8
	Flow           AbilityResourceType = 11
)

func GetAllChampNames() ([]string, error) {
	resp, err := http.Get(fmt.Sprintf("%s/game/global/champions/champions.bin.json", cDragonUrl))
	if err != nil {
		return nil, err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("unexpected code: %d", resp.StatusCode)
	}
	var res map[string]interface{}

	err = json.Unmarshal(body, &res)
	if err != nil {
		return nil, fmt.Errorf("error while marshaling champion data: %v", err)
	}

	var champs []string
	for _, i := range res {
		m, ok := i.(map[string]interface{})
		if !ok {
			continue
		}
		if val, ok := m["__type"]; ok {
			s := val.(string)
			if s == "Champion" {
				name := m["name"].(string)
				if name == "TFTChampion" {
					continue
				}
				champs = append(champs, name)
			}
		}
	}
	sort.Strings(champs)
	return champs, nil
}

func GetChamp(name string) (*Champion, error) {
	resp, err := http.Get(fmt.Sprintf("%s/game/data/characters/%s/%s.bin.json", cDragonUrl, strings.ToLower(name), strings.ToLower(name)))
	if err != nil {
		return nil, err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("unexpected code: %d", resp.StatusCode)
	}

	var res map[string]interface{}

	err = json.Unmarshal(body, &res)
	if err != nil {
		return nil, fmt.Errorf("error while marshaling champion data: %v", err)
	}
	val, ok := res[fmt.Sprintf("Characters/%s/CharacterRecords/Root", name)]

	if !ok {
		return nil, fmt.Errorf("error character record root not found in response")
	}

	champ := Champion{Character: CharacterRecord{}}

	err = JsonMapToInterface(val.(map[string]interface{}), &champ.Character)
	if err != nil {
		return nil, fmt.Errorf("error while marshaling champion data: %v", err)
	}

	spells := append(champ.Character.SpellNames, champ.Character.MCharacterPassiveSpell)
	for i, spellName := range spells {
		path := fmt.Sprintf("Characters/%s/Spells/%s", name, spellName)
		if i == len(spells)-1 {
			path = spellName
		}
		val = GetValueFromBinMap(path, res)
		if val == nil {
			fmt.Printf("error spell: %s not found in response\n", spellName)
			continue
		}
		spell := Spell{}

		_, ok = val.(map[string]interface{})
		if !ok {
			return nil, fmt.Errorf("error casting spell to interface")
		}
		err = JsonMapToInterface(val.(map[string]interface{}), &spell)
		if err != nil {
			return nil, fmt.Errorf("error while spell data: %v", err)
		}
		champ.Spells = append(champ.Spells, spell)
	}

	for _, abilityName := range champ.Character.MAbilities {
		val, ok = res[abilityName]
		if !ok {
			return nil, fmt.Errorf("error ability: %s not found in response", abilityName)
		}
		ability := Ability{}

		_, ok = val.(map[string]interface{})
		if !ok {
			return nil, fmt.Errorf("error casting ability to interface")
		}
		err = JsonMapToInterface(val.(map[string]interface{}), &ability)
		if err != nil {
			return nil, fmt.Errorf("error while marshaling ability data: %v", err)
		}
		champ.Abilities = append(champ.Abilities, ability)
	}

	return &champ, nil
}
