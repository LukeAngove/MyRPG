package main

import (
	"fmt"
	"regexp"
	"strings"
)

func IsUpper(c byte) bool {
	return 'A' <= c && c <= 'Z'
}

func ToLower(c byte) byte {
	return strings.ToLower(string(c))[0]
}

type ChallengeDraw struct {
	Spheres [3]string
}
type GemDraw struct {
	Type   string
	Sphere string
	Gems   []string
	IsX    bool
}

func make_challenge(challenge string, colors map[string]string) ChallengeDraw {
	res := ChallengeDraw{}
	for i, g := range challenge {
		res.Spheres[i] = colors[string(g)]
	}
	return res
}

func make_gem(gtype string, sphere string, gems string, colors map[string]string) GemDraw {
	sphere_color := colors[sphere]
	gem_colors := make([]string, len(gems))
	isX := false

	for i, g := range gems {
		g_col := byte(g)
		if IsUpper(g_col) {
			g_col = ToLower(g_col)
			isX = true
		}
		gem_colors[i] = colors[string(g_col)]
	}

	return GemDraw{
		gtype,
		sphere_color,
		gem_colors,
		isX,
	}
}

func parse_single(the_string string, colors map[string]string) GemDraw {
	colors_bytes := make([]byte, len(colors))
	for k, _ := range colors {
		colors_bytes = append(colors_bytes, []byte(k)[0])
	}

	colors_str := string(colors_bytes)
	upper_colors_str := string(strings.ToUpper(colors_str))

	gem_str := fmt.Sprintf("[%s]|[%s]+", upper_colors_str, colors_str)
	sphere_str := fmt.Sprintf("[%s]", colors_str)

	in_colors := fmt.Sprintf(`(%s):\((%s)\)`, gem_str, sphere_str)
	into_colors := fmt.Sprintf(`(%s)->\((%s)\)`, gem_str, sphere_str)
	out_colors := fmt.Sprintf(`\((%s)\)->(%s)`, sphere_str, gem_str)

	in_colors_re := regexp.MustCompile(in_colors)
	into_colors_re := regexp.MustCompile(into_colors)
	out_colors_re := regexp.MustCompile(out_colors)

	gemDraw := GemDraw{
		"None",
		"None",
		[]string{"None"},
		false,
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
