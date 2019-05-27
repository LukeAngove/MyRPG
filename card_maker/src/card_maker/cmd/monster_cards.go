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
}

func join(a MonsterCard, b MonsterCard) MonsterCard {
	res := MonsterCard{}

	res.Title = a.Title + " " + b.Title

	res.Attributes = a.Attributes
	for key, _ := range b.Attributes {
		res.Attributes[key] += b.Attributes[key]
	}

	res.Rules = a.Rules
	for key, val := range b.Rules {
		if src_val, ok := res.Rules[key]; ok {
			res.Rules[key] = src_val + val
		} else {
			res.Rules[key] = val
		}
	}

	res.Abilities = append(a.Abilities, b.Abilities...)
	res.Tags = append(a.Tags, b.Tags...)

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

func genCombinations(mod CardModifierHolder) CardModifierHolder {
	return mod
}

func generateCardsFromTemplate(generator CardGenerator) []MonsterCard{
	res := []MonsterCard{}

	this_card := generator.Base

	for _, overlay := range generator.Overlays {
		keys := make([]string, 0, len(overlay.Modifiers))
		for k := range overlay.Modifiers {
			keys = append(keys, k)
		}

		this_card = join(this_card, overlay.Modifiers[keys[0]])
		this_card.Title += " " + keys[0]
	}

	res = append(res, this_card)

	return res
}

func (cards *MonsterCards) Preprocess() {
	(*cards).Cards = generateCardsFromTemplate((*cards).Generator)
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