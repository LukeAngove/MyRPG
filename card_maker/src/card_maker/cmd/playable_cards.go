package main

import (
	"io/ioutil"
	"os"
	"regexp"
	"sort"
	"strings"
	"text/template"

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

func cost_string_replace(the_string string, scale float32, colors map[string]string) string {
	costs := regexp.MustCompile("cost{([a-zA-Z():\\->|]+)}")
	cost_to_svg := func(in_string string) string {
		in_string = costs.ReplaceAllString(in_string, "$1")
		gems := parse_all(in_string, colors)
		result := draw_gems_svg(gems, scale)
		return result
	}
	res := costs.ReplaceAllStringFunc(the_string, cost_to_svg)
	return res
}

func challenge_string_replace(the_string string, scale float32, colors map[string]string) string {
	costs := regexp.MustCompile("challenge{([a-z]{3})}")
	challenge_to_svg := func(in_string string) string {
		in_string = costs.ReplaceAllString(in_string, "$1")
		chal := make_challenge(in_string, colors)
		result := draw_challenge_svg(chal, scale)
		return result
	}
	res := costs.ReplaceAllStringFunc(the_string, challenge_to_svg)
	return res
}

func makePlayableCards(cards PlayableCards, template *template.Template, output string) {
	file, err := os.Create(output)
	check(err)
	err = template.Execute(file, cards)
	check(err)
}

func makeCardsFromTemplate(yaml_desc string, html_template string, output string) {
	data, err := ioutil.ReadFile(yaml_desc)
	check(err)

	cards := PlayableCards{}
	err = yaml.Unmarshal(data, &cards)
	check(err)

	cards = preprocessCards(cards)

	filePath := strings.Split(html_template, "/")
	fileName := filePath[len(filePath)-1]
	t, err := template.New(fileName).Funcs(template.FuncMap{
		"renderCost": func(str string) string {
			return cost_string_replace(str, 1.0, cards.Colors)
		},
		"renderChallenge": func(str string) string {
			return challenge_string_replace(str, 1.0, cards.Colors)
		},
		"join": func(in []string) string {
			return strings.Join(in, ", ")
		},
	}).ParseFiles(html_template)

	makePlayableCards(cards, t, output)
}
