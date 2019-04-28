package main

import (
	"flag"
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

type GemDraw struct {
	Type   string
	Sphere string
	Gems   []string
}

type Cards struct {
	Cards  []MyCard          `yaml:"cards"`
	Rules  []string          `yaml:"rules_order"`
	Colors map[string]string `yaml:"colors"`
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func text_to_svg(in_string string, scale float32, colors map[string]string) string {
	gems := parse_all(in_string, colors)
	result := draw_gems_svg(gems, scale)
	return result
}

func makePlayableCards(cards Cards, template *template.Template, output string) {
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

	cards := Cards{}
	err = yaml.Unmarshal(data, &cards)
	check(err)

	t := template.Must(template.ParseFiles(paths...))

	makePlayableCards(cards, t, output)
}

func main() {
	var card_source = flag.String("card_outlines", "cards.yml", "YAML file describing cards")
	var html_template = flag.String("html_template", "default.html", "HTML template for cards")
	var output = flag.String("output", "cards.html", "HTML output for cards")
	flag.Parse()

	makeCardsFromTemplate(*card_source, *html_template, *output)
}
