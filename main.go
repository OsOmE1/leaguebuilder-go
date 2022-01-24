package main

import (
	"fmt"
	"leaguebuilder/lol"
	"leaguebuilder/lol/datadragon"
)

func main() {
	allChamps, err := datadragon.GetAllChampNames()
	if err != nil {
		panic(err)
	}
	fc, err := datadragon.GetFontConfig("en_us")
	if err != nil {
		panic(err)
	}
	//apiChamp, err := datadragon.GetChamp("Jinx")
	//if err != nil {
	//	panic(err)
	//}
	//champ, err := lol.ChampionFromApi(apiChamp, fc)
	//if err != nil {
	//	panic(err)
	//}
	//ci := champ.Instance(3)
	//for _, spell := range champ.Spells {
	//	fmt.Println(spell.Name)
	//	fmt.Println(spell.ResolveSpellText(ci, 1))
	//	fmt.Println()
	//}
	var types []string
	for _, c := range allChamps {
		fmt.Println(c)
		apiChamp, err := datadragon.GetChamp(c)
		if err != nil {
			panic(err)
		}

		champ, err := lol.ChampionFromApi(apiChamp, fc)
		if err != nil {
			panic(err)
		}
		//ci := champ.Instance(1)
		for _, spell := range champ.Spells {
			types = append(types, spell.GetRatios()...)
		}
	}
	types = datadragon.RemoveDuplicateStr(types)
	for _, i := range types {
		fmt.Println(i)
	}
}
