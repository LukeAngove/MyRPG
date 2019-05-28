package main

import (
	"cardgenerator"
	"sort"

	"gopkg.in/yaml.v2"
)

type MyCard struct {
	Title        string
	Image        string
	Cost         string
	Tags         []string
	Rules        map[string]string
	Flavor       string
	ElementColor string
	TypeColor    string
}

type Scores map[string]map[string]int
type TypeColors []map[string]string
type PlayableCards struct {
	Cards      []MyCard          `yaml:"cards"`
	Rules      []string          `yaml:"rules_order"`
	TagScores  Scores            `yaml:"tag_scores"`
	TypeColors TypeColors        `yaml:"card_type_colors"`
	Colors     map[string]string `yaml:"colors"`
}

func rankByValue(scores map[string]int) PairList {
	pl := make(PairList, len(scores))
	i := 0
	for k, v := range scores {
		pl[i] = Pair{k, v}
		i++
	}
	sort.Sort(sort.Reverse(pl))
	return pl
}

type Pair struct {
	Key   string
	Value int
}

type PairList []Pair

func (p PairList) Len() int           { return len(p) }
func (p PairList) Less(i, j int) bool { return p[i].Value < p[j].Value }
func (p PairList) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

func sort_card(card MyCard) MyCard {
	sort.Strings(card.Tags)
	return card
}

func card_color_score(card MyCard, scores Scores) map[string]int {
	score := map[string]int{"l": 0, "f": 0, "e": 0, "w": 0, "s": 0}

	for _, tag := range card.Tags {
		for k, _ := range score {
			// Check if tag is in list
			if val, ok := scores[k][tag]; ok {
				score[k] += val
			}
		}
	}
	return score
}

func card_type_color(card MyCard, type_colors TypeColors) string {
	for _, t := range type_colors {
		for _, tag := range card.Tags {
			if val, ok := t[tag]; ok {
				return val
			}
		}
	}
	return "invalid"
}

func preprocessCards(cards PlayableCards) PlayableCards {
	for i, c := range cards.Cards {
		cards.Cards[i] = sort_card(c)
		scores := card_color_score(c, cards.TagScores)
		colors := rankByValue(scores)
		card_element := colors[0].Key
		cards.Cards[i].ElementColor = cards.Colors[card_element]
		cards.Cards[i].TypeColor = card_type_color(c, cards.TypeColors)
	}
	return cards
}

func (cards *PlayableCards) Preprocess() {
	*cards = preprocessCards(*cards)
}

func (cards *PlayableCards) Load(yaml_bytes []byte) {
	err := yaml.Unmarshal(yaml_bytes, cards)
	cardgenerator.Check(err)
	cards.Preprocess()
}

func main() {
	cards := PlayableCards{}
	cardgenerator.MakeCardsFromTemplate(&cards)
}
