package main

import (
	"cardgenerator"
	"gopkg.in/yaml.v2"
)

type MonsterCard struct {
	Title     string
	Image     string
	Tags      []string
	Attributes map[string]int
	Rules     map[string]string
	Abilities []map[string]string
	Flavor    string
}

func NewMonsterCard() MonsterCard {
	card := MonsterCard{}
	card.Tags = []string{}
	card.Attributes = map[string]int{
		"L": 0,
		"F": 0,
		"E": 0,
		"W": 0,
		"S": 0,
	}
	card.Rules = map[string]string{}
	card.Abilities = []map[string]string{}
	card.Flavor = ""

	return card
}

type CardModifierHolder struct {
	MaxInclude int `yaml:"max_include"`
	MinInclude int `yaml:"min_include"`
    Modifiers map[string]MonsterCard `yaml:"items"`
}

type CardGenerator struct {
    Base MonsterCard `yaml:"base"`
	Overlays map[string]CardModifierHolder `yaml:"overlays"`
}

type MonsterCards struct {
	Generator CardGenerator   `yaml:"generator"`
	Cards     []MonsterCard
	Rules     []string          `yaml:"rules_order"`
	Abilities []string          `yaml:"abilities_order"`
	Overlays  []string          `yaml:"overlay_order"`
}

func join(a MonsterCard, b MonsterCard) MonsterCard {
	res := NewMonsterCard()

	res.Title = a.Title + " " + b.Title

	for key, _ := range a.Attributes {
		res.Attributes[key] += a.Attributes[key]
	}
	for key, _ := range b.Attributes {
		res.Attributes[key] += b.Attributes[key]
	}

	for key, val := range a.Rules {
		if src_val, ok := res.Rules[key]; ok {
			res.Rules[key] = src_val + val
		} else {
			res.Rules[key] = val
		}
	}
	for key, val := range b.Rules {
		if src_val, ok := res.Rules[key]; ok {
			res.Rules[key] = src_val + val
		} else {
			res.Rules[key] = val
		}
	}

	res.Abilities = append(res.Abilities, a.Abilities...)
	res.Abilities = append(res.Abilities, b.Abilities...)

	res.Tags = append(res.Tags, a.Tags...)
	res.Tags = append(res.Tags, b.Tags...)

	return res
}

func factorial(n int) int {
	res := 1

	for i := 1; i<=n; i++ {
		res *= i
	}

	return res
}

func combinations(n int, r int) int {
	return factorial(n) / (factorial(r) * factorial(n-r))
}

func getMultiplier(mod CardModifierHolder) int {
	res := 0
	for i := mod.MinInclude; i <= mod.MaxInclude; i++ {
		res += combinations(i, len(mod.Modifiers))
	}
	return res
}

func genCombinations(mod CardModifierHolder) map[string]MonsterCard {
	mods := mod.Modifiers

	if mod.MinInclude == 0 {
		mods[""] = MonsterCard{}
	}

	return mod.Modifiers
}

func applyOverlay(base_card MonsterCard, overlays CardModifierHolder) []MonsterCard {
	res := []MonsterCard{}

	modifiers := genCombinations(overlays)

	for key, overlay := range modifiers {
		new_card := join(base_card, overlay)
		new_card.Title += " " + key
		res = append(res, new_card)
	}

	return res
}

func generateCardsFromTemplate(cards MonsterCards) []MonsterCard {
	res := []MonsterCard{}
	generator := cards.Generator

	this_card := generator.Base
	res = append(res, this_card)

	for _, overlay_type := range cards.Overlays {
		overlay := generator.Overlays[overlay_type]

		new_res := []MonsterCard{}
		for _, card := range res {
			new_res = append(new_res, applyOverlay(card, overlay)...)
		}

		res = new_res
	}

	return res
}

func (cards *MonsterCards) Preprocess() {
	(*cards).Cards = generateCardsFromTemplate(*cards)
}

func (cards *MonsterCards) Load(yaml_bytes []byte) {
	err := yaml.Unmarshal(yaml_bytes, cards)
	cardgenerator.Check(err)
	cards.Preprocess()
}

func main() {
	cards := MonsterCards{}
	cardgenerator.MakeCardsFromTemplate(&cards)
}