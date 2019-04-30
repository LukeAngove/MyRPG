package main

import (
	"io/ioutil"
	"os"
	"regexp"
	"strings"
	"text/template"

	"gopkg.in/yaml.v2"
)

type MyCard struct {
	Title  string
	Image  string
	Cost   string
	Tags   []string
	Rules  map[string]string
	Flavor string
}

type PlayableCards struct {
	Cards  []MyCard          `yaml:"cards"`
	Rules  []string          `yaml:"rules_order"`
	Colors map[string]string `yaml:"colors"`
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

	filePath := strings.Split(html_template, "/")
	fileName := filePath[len(filePath)-1]
	t, err := template.New(fileName).Funcs(template.FuncMap{
		"renderCost": func(str string) string {
			return cost_string_replace(str, 3.0, cards.Colors)
		},
		"renderChallenge": func(str string) string {
			return challenge_string_replace(str, 3.0, cards.Colors)
		},
	}).ParseFiles(html_template)

	makePlayableCards(cards, t, output)
}
