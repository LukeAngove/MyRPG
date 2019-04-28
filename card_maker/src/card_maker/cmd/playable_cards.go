package main

import (
	"io/ioutil"
	"os"
	"text/template"

	"gopkg.in/yaml.v2"
)

type MyCard struct {
	Title  string
	Image  string
	Cost   string
	Rules  map[string]string
	Flavor string
}

type PlayableCards struct {
	Cards  []MyCard          `yaml:"cards"`
	Rules  []string          `yaml:"rules_order"`
	Colors map[string]string `yaml:"colors"`
}

func text_to_svg(in_string string, scale float32, colors map[string]string) string {
	gems := parse_all(in_string, colors)
	result := draw_gems_svg(gems, scale)
	return result
}

func makePlayableCards(cards PlayableCards, template *template.Template, output string) {
	for i, c := range cards.Cards {
		cards.Cards[i].Cost = text_to_svg(c.Cost, 4, cards.Colors)
	}

	file, err := os.Create(output)
	check(err)
	err = template.Execute(file, cards)
	check(err)
}

func makeCardsFromTemplate(yaml_desc string, html_template string, output string) {
	paths := []string{
		html_template,
	}

	data, err := ioutil.ReadFile(yaml_desc)
	check(err)

	cards := PlayableCards{}
	err = yaml.Unmarshal(data, &cards)
	check(err)

	t := template.Must(template.ParseFiles(paths...))

	makePlayableCards(cards, t, output)
}
