package cardgenerator

import (
	"os"
	"flag"
	"regexp"
	"text/template"
	"strings"
	"io/ioutil"
)

type ICards interface {
	Load([]byte)
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

func parseArgs() (string, string, string) {
	var card_source = flag.String("card_outlines", "playable.yml", "YAML file describing cards")
	var html_template = flag.String("html_template", "playable.html", "HTML template for cards")
	var output = flag.String("output", "cards.html", "HTML output for cards")
	flag.Parse()

	return *card_source, *html_template, *output
}

func readTemplate(html_template string) *template.Template {
	colors := map[string]string{
		"a": "grey",
		"i": "white",
		"t": "black",
		"l": "yellow",
		"f": "red",
		"e": "green",
		"w": "blue",
		"s": "purple",
	}
	filePath := strings.Split(html_template, "/")
	fileName := filePath[len(filePath)-1]
	t, err := template.New(fileName).Funcs(template.FuncMap{
		"renderCost": func(str string) string {
			return cost_string_replace(str, 1.0, colors)
		},
		"renderChallenge": func(str string) string {
			return challenge_string_replace(str, 1.0, colors)
		},
		"join": func(in []string) string {
			return strings.Join(in, ", ")
		},
	}).ParseFiles(html_template)
	Check(err)

	return t
}

func makeCards(cards ICards, template *template.Template, output string) {
	file, err := os.Create(output)
	Check(err)
	err = template.Execute(file, cards)
	Check(err)
}

func MakeCardsFromTemplate(cards ICards) {
	yaml_desc, html_template, output := parseArgs()

	data, err := ioutil.ReadFile(yaml_desc)
	Check(err)

	cards.Load(data)

	t := readTemplate(html_template)
	makeCards(cards, t, output)
}
