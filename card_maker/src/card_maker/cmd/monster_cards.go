package main

import (
	"io/ioutil"
	"os"
	"text/template"

	"gopkg.in/yaml.v2"
)

type MonsterCard struct {
	Title     string
	Image     string
	Cost      string
	Rules     map[string]string
	Abilities []map[string]string
	Flavor    string
}

type Cards struct {
	Cards     []MonsterCard     `yaml:"cards"`
	Rules     []string          `yaml:"rules_order"`
	Abilities []string          `yaml:"abilities_order"`
	Colors    map[string]string `yaml:"colors"`
}

func makeMonsterCards(cards Cards, template *template.Template, output string) {
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

	makeMonsterCards(cards, t, output)
}
