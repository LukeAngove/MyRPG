package main

import (
	"flag"
)

func main() {
	var card_source = flag.String("card_outlines", "playable.yml", "YAML file describing cards")
	var html_template = flag.String("html_template", "playable.html", "HTML template for cards")
	var output = flag.String("output", "cards.html", "HTML output for cards")
	flag.Parse()

	makeCardsFromTemplate(*card_source, *html_template, *output)
}
