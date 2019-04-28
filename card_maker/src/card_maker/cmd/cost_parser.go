package main

import (
	"fmt"
	"regexp"
	"strings"
)

type GemDraw struct {
	Type   string
	Sphere string
	Gems   []string
}

func make_gem(gtype string, sphere string, gems string, colors map[string]string) GemDraw {
	sphere_color := colors[sphere]
	gem_colors := make([]string, len(gems))
	for i, g := range gems {
		gem_colors[i] = colors[string(g)]
	}

	return GemDraw{
		gtype,
		sphere_color,
		gem_colors,
	}
}

func parse_single(the_string string, colors map[string]string) GemDraw {
	colors_bytes := make([]byte, len(colors))
	for k, _ := range colors {
		colors_bytes = append(colors_bytes, []byte(k)[0])
	}

	colors_str := string(colors_bytes)

	in_colors_re := regexp.MustCompile(fmt.Sprintf(`([%s]+):\(([%s])\)`, colors_str, colors_str))
	into_colors_re := regexp.MustCompile(fmt.Sprintf(`([%s]+)->\(([%s])\)`, colors_str, colors_str))
	out_colors_re := regexp.MustCompile(fmt.Sprintf(`\(([%s])\)->([%s]+)`, colors_str, colors_str))

	gemDraw := GemDraw{
		"None",
		"None",
		[]string{"None"},
	}

	groups := in_colors_re.FindStringSubmatch(the_string)
	if groups != nil {
		gemDraw = make_gem("in", groups[2], groups[1], colors)
	} else {
		groups = into_colors_re.FindStringSubmatch(the_string)
		if groups != nil {
			gemDraw = make_gem("into", groups[2], groups[1], colors)
		} else {
			groups = out_colors_re.FindStringSubmatch(the_string)
			if groups != nil {
				gemDraw = make_gem("out", groups[1], groups[2], colors)
			} else {
				panic(fmt.Sprintf("No match for pattern: %s", the_string))
			}
		}
	}

	return gemDraw
}

func parse_all(the_string string, colors map[string]string) []GemDraw {
	the_string = strings.Replace(the_string, " ", "", -1)
	gems := strings.Split(the_string, "|")
	res := make([]GemDraw, len(gems))
	for i, g := range gems {
		res[i] = parse_single(g, colors)
	}
	return res
}
