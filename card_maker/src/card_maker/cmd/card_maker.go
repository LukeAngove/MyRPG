package main

import (
	"bytes"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strings"
	"text/template"

	"gopkg.in/yaml.v2"

	svg "github.com/ajstarks/svgo"
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

type Pos struct {
	x float32
	y float32
}

type Drawer struct {
	canvas    *svg.SVG
	outbuffer bytes.Buffer
	scale     float32
	offset    Pos
}

func generate_diagram() {
	width := 500
	height := 500
	canvas := svg.New(os.Stdout)
	canvas.Start(width, height)
	canvas.Circle(width/2, height/2, 100)
	canvas.Text(width/2, height/2, "Hello, SVG", "text-anchor:middle;font-size:30px;fill:white")
	canvas.End()
}

func do_rotate_90(pt Pos, centre Pos) Pos {
	/*
		    Translation:
		    [x']   [1  0 cx][x]
		    [y'] = [0  1 cy][y]
		    [1 ]   [0  0  1][1]

		    Rotation:
		    [x']   [0 -1  0][x]
		    [y'] = [1  0  0][y]
		    [1 ]   [0  0  1][1]

		    [x']   [1  0 cx'][0 -1  0][1  0 -cx][x]
		    [y'] = [0  1 cy'][1  0  0][0  1 -cy][y]
		    [1 ]   [0  0  1][0  0  1][0  0   1][1]
		    =>
		    [x']   [0 -1 cx'+cy][x]
		    [y'] = [1  0 cy'-cx][y]
		    [1 ]   [0  0     1][1]
		    =>
		    x' = cy+cx'-y = cx'+cy-y
			y' = x+cy'-cx = cy'+x-cx
	*/
	newPt := Pos{centre.x + centre.y - pt.y, centre.y + pt.x - centre.x}
	return newPt
}

func do_symetry(pts []Pos, centre Pos) []Pos {
	pts_size := len(pts)
	for i := 0; i < 3; i++ {
		for j := 0; j < pts_size; j++ {
			pts = append(pts, do_rotate_90(pts[i*pts_size+j], centre))
		}
	}
	return pts
}

func octagon(color string, position Pos, corner_size float32, flat_size float32, outer_line_width float32) []Pos {
	edging := outer_line_width / 2
	total_width := edging*2 + corner_size*2 + flat_size
	centre := Pos{position.x + (total_width / 2), position.y + total_width/2}

	c1 := Pos{edging + position.x, edging + corner_size + position.y}
	c2 := Pos{edging + corner_size + position.x, edging + position.y}
	pts_arr := []Pos{c1, c2}
	pts_arr = do_symetry(pts_arr, centre)
	return pts_arr
}

func octagon_internal_corners(octagon_pts []Pos) [][]Pos {
	size := len(octagon_pts)
	// We add 1 point for every 2, for the corners inbetween
	out := make([][]Pos, size/2)
	for i := range out {
		out[i] = make([]Pos, 3)
	}
	for i := 0; i < (size / 2); i++ {
		pos := Pos{0, 0}
		start := i * 2
		if i%2 == 1 {
			pos = Pos{octagon_pts[start].x, octagon_pts[start+1].y}
		} else {
			pos = Pos{octagon_pts[start+1].x, octagon_pts[start].y}
		}
		out[i][0] = octagon_pts[start+0]
		out[i][1] = pos
		out[i][2] = octagon_pts[start+1]
	}
	return out
}

func convert_pos(pos []Pos) ([]int, []int) {
	x := make([]int, len(pos))
	y := make([]int, len(pos))
	for i, p := range pos {
		x[i] = int(p.x)
		y[i] = int(p.y)
	}
	return x, y
}

func draw_gem(canvas *svg.SVG, color string, position Pos, scale float32) {
	offset := 0.57 * scale // To make gem fit inside sphere
	position = Pos{position.x + offset + 1, position.y + offset + 1}
	outer_line_width := float32(2.0)
	corner_size := 1 * scale
	flat_size := 2 * scale
	edging := outer_line_width / 2

	octagon_pts := octagon(color, position, corner_size, flat_size, outer_line_width)
	oct_x, oct_y := convert_pos(octagon_pts)
	canvas.Polygon(oct_x, oct_y, fmt.Sprintf("fill:%s;stroke_width:%d;stroke:%s", color, int(outer_line_width), "black"))

	corner_pts := octagon_internal_corners(octagon_pts)

	for _, p := range corner_pts {
		cnr_x, cnr_y := convert_pos(p)
		canvas.Polyline(cnr_x, cnr_y, fmt.Sprintf("fill:%s;stroke_width:%d;stroke:%s", "none", int(edging), "black"))
	}

	rect_x := int(position.x + edging + corner_size)
	rect_y := int(position.y + edging + corner_size)
	canvas.Rect(rect_x, rect_y, int(flat_size), int(flat_size),
		fmt.Sprintf("stroke_width:%d;stroke:%s;fill:%s", int(edging), "black", "none"))
}

func draw_sphere(canvas *svg.SVG, color string, position Pos, scale float32) {
	gem_size_factor := 2.6 * scale
	outer_line_width := float32(2.0)
	centre := Pos{position.x + gem_size_factor + 1, position.y + gem_size_factor + 2}
	canvas.Circle(int(centre.x), int(centre.y), int(gem_size_factor),
		fmt.Sprintf("fill:%s;stroke:%s;stroke_width:%d", color, "black", int(outer_line_width)))
}

func grow(array []Pos, offset Pos, scale float32) []Pos {
	res := make([]Pos, len(array))
	for i, p := range array {
		res[i].x = p.x*scale + offset.x
		res[i].y = p.y*scale + offset.y
	}
	return res
}

func draw_arrow(canvas *svg.SVG, position Pos, scale float32) {
	arrow_mid := []Pos{{0, 1}, {2, 1}}
	arrow_point := []Pos{{1, 0}, {2, 1}, {1, 2}}
	outer_line_width := float32(2.0)

	position = Pos{position.x, position.y + 1.9*scale}
	arrow_mid = grow(arrow_mid, position, scale)
	arrow_point = grow(arrow_point, position, scale)
	pts_mid_x, pts_mid_y := convert_pos(arrow_mid)
	pts_point_x, pts_point_y := convert_pos(arrow_point)
	format := fmt.Sprintf("fill:%s;stroke_width:%d;stroke:%s", "none", int(outer_line_width), "black")
	canvas.Polyline(pts_mid_x, pts_mid_y, format)
	canvas.Polyline(pts_point_x, pts_point_y, format)
}

func draw_bar(canvas *svg.SVG, position Pos, scale float32) {
	bar := []Pos{{0, 0.5}, {0, 7}}
	outer_line_width := float32(2.0)

	bar = grow(bar, position, scale)
	pts_bar_x, pts_bar_y := convert_pos(bar)
	format := fmt.Sprintf("fill:%s;stroke_width:%d;stroke:%s", "none", int(outer_line_width), "black")
	canvas.Polyline(pts_bar_x, pts_bar_y, format)
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

func calc_width(gems []GemDraw, scale float32) float32 {
	width := float32(0.0)
	for _, g := range gems {
		if g.Type == "into" {
			width += 6*scale*2 + 3.3*scale
		} else if g.Type == "in" {
			width += 6 * scale
		} else if g.Type == "out" {
			width += 6*scale*2 + 3.3*scale
		} else {
			panic(fmt.Sprintf("Invalid gem for width calc: %s", g))
		}
	}
	// Add width for spacing
	width += 0.5 * float32(len(gems)-1)

	return width
}

func draw_gems(canvas *svg.SVG, colors []string, position Pos, scale float32) {
	number := len(colors)
	switch number {
	case 1:
		pos := Pos{position.x + 0.5*scale, position.y + 0.7*scale}
		draw_gem(canvas, colors[0], pos, 0.75*scale)
		break
	case 2:
		pos1 := Pos{position.x + 0.6*scale, position.y}
		pos2 := Pos{position.x + 0.6*scale, position.y + 1.2*scale}
		draw_gem(canvas, colors[0], pos1, 0.75*scale)
		draw_gem(canvas, colors[1], pos2, 0.75*scale)
		break
	case 3:
		pos1 := Pos{position.x + 0.6*scale, position.y}
		pos2 := Pos{position.x + 1.2*scale, position.y + 0.9*scale}
		pos3 := Pos{position.x + 0.0*scale, position.y + 0.9*scale}
		// Draw 3 first, so that it's on the bottom
		draw_gem(canvas, colors[2], pos3, 0.75*scale)
		draw_gem(canvas, colors[0], pos1, 0.75*scale)
		draw_gem(canvas, colors[1], pos2, 0.75*scale)
		break
	}
}

func (drawer *Drawer) draw_sphere(sphere_color string) {
	draw_sphere(drawer.canvas, sphere_color, drawer.offset, drawer.scale)
}

func (drawer *Drawer) draw_gems(gem_colors []string) {
	draw_gems(drawer.canvas, gem_colors, drawer.offset, drawer.scale)
}

func (drawer *Drawer) draw_arrow() {
	draw_arrow(drawer.canvas, drawer.offset, drawer.scale)
}

func (drawer *Drawer) draw_bar() {
	draw_bar(drawer.canvas, drawer.offset, drawer.scale)
	drawer.bar_increment()
}

func (drawer *Drawer) big_increment() {
	drawer.offset.x += 6 * drawer.scale
}

func (drawer *Drawer) small_increment() {
	drawer.offset.x += 3.3 * drawer.scale
}

func (drawer *Drawer) bar_increment() {
	drawer.offset.x += 0.5 * drawer.scale
}

func (drawer *Drawer) draw_in(sphere_color string, gem_colors []string) {
	drawer.draw_sphere(sphere_color)
	drawer.draw_gems(gem_colors)
	drawer.big_increment()
}

func (drawer *Drawer) draw_into(sphere_color string, gem_colors []string) {
	drawer.draw_gems(gem_colors)
	drawer.big_increment()
	drawer.draw_arrow()
	drawer.small_increment()
	drawer.draw_sphere(sphere_color)
	drawer.big_increment()
}

func (drawer *Drawer) draw_out(sphere_color string, gem_colors []string) {
	drawer.draw_sphere(sphere_color)
	drawer.big_increment()
	drawer.draw_arrow()
	drawer.small_increment()
	drawer.draw_gems(gem_colors)
	drawer.big_increment()
}

func NewDrawer(width float32, scale float32) *Drawer {
	height := 6 * scale

	d := new(Drawer)
	canvas := svg.New(&d.outbuffer)
	canvas.Start(int(width), int(height))
	offset := Pos{0, 0}
	d.canvas = canvas
	d.scale = scale
	d.offset = offset
	return d
}

func (drawer *Drawer) Finish() string {
	drawer.canvas.End()
	return drawer.outbuffer.String()
}

func text_to_svg(in_string string, scale float32, colors map[string]string) string {
	gems := parse_all(in_string, colors)
	width := calc_width(gems, scale)

	drawer := NewDrawer(width, scale)

	for i, g := range gems {
		if i != 0 {
			drawer.draw_bar()
		}
		if g.Type == "into" {
			drawer.draw_into(g.Sphere, g.Gems)
		} else if g.Type == "in" {
			drawer.draw_in(g.Sphere, g.Gems)
		} else if g.Type == "out" {
			drawer.draw_out(g.Sphere, g.Gems)
		} else {
			panic(fmt.Sprintf("Invalid gem: %s", g))
		}
	}

	return drawer.Finish()
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

	for i, c := range cards.Cards {
		cards.Cards[i].Cost = text_to_svg(c.Cost, 4, cards.Colors)
	}

	file, err := os.Create(output)
	check(err)
	err = t.Execute(file, cards)
	check(err)
}

func main() {
	var card_source = flag.String("card_outlines", "cards.yml", "YAML file describing cards")
	var html_template = flag.String("html_template", "default.html", "HTML template for cards")
	var output = flag.String("output", "cards.html", "HTML output for cards")
	flag.Parse()

	makeCardsFromTemplate(*card_source, *html_template, *output)
}
