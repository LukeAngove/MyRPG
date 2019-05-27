package cardgenerator

import (
	"bytes"
	"fmt"

	svg "github.com/ajstarks/svgo"
)

const sphereInternalBorder = 10.0
const sphereWidth = 94.0
const gemWidth = 60.0
const gemXWidth = 70.0
const arrowWidth = 60.0
const arrowOverlap = arrowWidth / 6.0
const lineWidth = 5.0

type Drawer struct {
	canvas    *svg.SVG
	outbuffer bytes.Buffer
	lineWidth float32
	offset    Pos
}

type Pos struct {
	x float32
	y float32
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

func draw_gem(canvas *svg.SVG, color string, position Pos, radius float32, lineWidth float32) {
	centre := Pos{position.x, position.y}
	canvas.Circle(int(centre.x), int(centre.y), int(radius),
		fmt.Sprintf("fill:%s;stroke:%s;stroke-width:%f", color, "black", lineWidth))
}

func draw_gem_X(canvas *svg.SVG, color string, position Pos, lineWidth float32, isX bool) {
	if isX {
		draw_gem(canvas, color, position, gemXWidth/2.0, lineWidth)
		font_size := gemXWidth - 2*lineWidth
		style := fmt.Sprintf("font-size:%f;text-anchor:middle;alignment-baseline:central", font_size)
		if color == "black" {
			style = style + "fill=silver"
		}
		canvas.Text(int(position.x), int(position.y+font_size/2-2*lineWidth), "X", style)
	} else {
		draw_gem(canvas, color, position, gemWidth/2.0, lineWidth)
	}
}

func draw_sphere(canvas *svg.SVG, color string, position Pos, lineWidth float32) {
	gem_size_factor := float32(sphereWidth / 2.0)
	centre := Pos{position.x + gem_size_factor, 50}
	canvas.Circle(int(centre.x), int(centre.y), int(gem_size_factor),
		fmt.Sprintf("fill:%s;stroke:%s;stroke-width:%f", color, "black", lineWidth))
}

func shift(array []Pos, offset Pos) []Pos {
	res := make([]Pos, len(array))
	for i, p := range array {
		res[i].x = p.x + offset.x
		res[i].y = p.y + offset.y
	}
	return res
}

func draw_arrow(canvas *svg.SVG, position Pos, lineWidth float32) {
	arrow := []Pos{{0, 20}, {30, 20}, {10, 0}, {30, 0}, {60, 30}, {30, 60}, {10, 60}, {30, 40}, {0, 40}}

	position = Pos{position.x, 20}
	arrow = shift(arrow, position)
	pts_x, pts_y := convert_pos(arrow)
	format := fmt.Sprintf("fill:%s;stroke-width:%f;stroke:%s", "gray", lineWidth, "black")
	canvas.Polygon(pts_x, pts_y, format)
}

func calc_width(size float32) float32 {
	width := float32(0.0)
	fudge := float32(1.0)
	width = (sphereWidth+arrowOverlap*2.0)*size*fudge + lineWidth*2
	return width
}

func draw_gems(canvas *svg.SVG, colors []string, position Pos, lineWidth float32, IsX bool) {
	number := len(colors)
	if len(colors) != 1 && IsX {
		panic("More than one gem color, and using X! Not a valid color set!")
	}

	topRow := float32(sphereInternalBorder + gemWidth/2.0 + 1)
	//bottomRow := float32(sphereWidth - (sphereInternalBorder + gemWidth/2.0))
	bottomRow := float32(sphereWidth - (gemWidth/2.0 + sphereInternalBorder/2))
	colOffset := float32(10)
	switch number {
	case 1:
		pos := Pos{position.x, 50}
		draw_gem_X(canvas, colors[0], pos, lineWidth, IsX)
		break
	case 2:
		pos1 := Pos{position.x, topRow}
		pos2 := Pos{position.x, bottomRow}
		draw_gem_X(canvas, colors[0], pos1, lineWidth, false)
		draw_gem_X(canvas, colors[1], pos2, lineWidth, false)
		break
	case 3:
		pos1 := Pos{position.x - colOffset, bottomRow}
		pos2 := Pos{position.x + colOffset, bottomRow}
		pos3 := Pos{position.x + 0, topRow}
		// Draw 3 first, so that it's on the bottom
		draw_gem_X(canvas, colors[2], pos3, lineWidth, false)
		draw_gem_X(canvas, colors[0], pos1, lineWidth, false)
		draw_gem_X(canvas, colors[1], pos2, lineWidth, false)
		break
	}
}

func (drawer *Drawer) draw_sphere(sphere_color string) {
	draw_sphere(drawer.canvas, sphere_color, drawer.offset, drawer.lineWidth)
}

func (drawer *Drawer) draw_gems(gem_colors []string, IsX bool) {
	offset := Pos{drawer.offset.x + sphereWidth/2.0, drawer.offset.y}
	draw_gems(drawer.canvas, gem_colors, offset, drawer.lineWidth, IsX)
}

func (drawer *Drawer) draw_arrow() {
	draw_arrow(drawer.canvas, drawer.offset, drawer.lineWidth)
}

func (drawer *Drawer) increment(inc float32) {
	drawer.offset.x += inc
}

func (drawer *Drawer) moveSphereWidth(ratio float32) {
	drawer.offset.x += (sphereWidth) * ratio
}

func (drawer *Drawer) draw_challenge(sphere_colors [3]string) {
	drawer.draw_sphere(sphere_colors[1])
	drawer.moveSphereWidth(2.0 / 3.0)
	drawer.draw_sphere(sphere_colors[2])
	drawer.moveSphereWidth(-1.0 / 3.0)
	drawer.draw_sphere(sphere_colors[0])
	drawer.moveSphereWidth(2.0 / 3.0)
}

func (drawer *Drawer) draw_gems_in_sphere(sphere_color string, gem_colors []string, IsX bool) {
	drawer.draw_sphere(sphere_color)
	drawer.draw_gems(gem_colors, IsX)
}

func (drawer *Drawer) draw_in(sphere_color string, gem_colors []string, IsX bool) {
	drawer.increment(arrowOverlap)
	drawer.draw_gems_in_sphere(sphere_color, gem_colors, IsX)
	drawer.moveSphereWidth(1.0)
	drawer.increment(arrowOverlap)
}

func (drawer *Drawer) draw_into(sphere_color string, gem_colors []string, IsX bool) {
	drawer.increment(arrowOverlap)
	drawer.draw_gems_in_sphere(sphere_color, gem_colors, IsX)
	drawer.increment(-arrowOverlap)
	drawer.draw_arrow()
	drawer.moveSphereWidth(1.0)
	drawer.increment(arrowOverlap * 2)
}

func (drawer *Drawer) draw_out(sphere_color string, gem_colors []string, IsX bool) {
	drawer.increment(arrowOverlap)
	drawer.draw_gems_in_sphere(sphere_color, gem_colors, IsX)
	drawer.moveSphereWidth(1.0)
	drawer.increment(arrowOverlap - arrowWidth)
	drawer.draw_arrow()
	drawer.increment(arrowWidth)
}

func NewDrawer(rel_width float32) *Drawer {
	rel_height := float32(100.0)

	act_height := float32(1.0)
	act_width := float32(act_height * rel_width / rel_height)

	d := new(Drawer)
	canvas := svg.New(&d.outbuffer)
	canvas.Start(0, 0, fmt.Sprintf("viewbox=\"0 0 %d %d\"style=\"height:%fem;width:%fem;position:relative;top:.25em\"", int(rel_width), int(rel_height), act_height, act_width))
	d.canvas = canvas
	d.lineWidth = lineWidth
	d.offset = Pos{d.lineWidth, 0}
	return d
}

func (drawer *Drawer) Finish() string {
	drawer.canvas.End()
	return drawer.outbuffer.String()
}

func draw_gems_svg(gems []GemDraw, scale float32) string {
	width := calc_width(float32(len(gems)))
	drawer := NewDrawer(width)

	for _, g := range gems {
		if g.Type == "into" {
			drawer.draw_into(g.Sphere, g.Gems, g.IsX)
		} else if g.Type == "in" {
			drawer.draw_in(g.Sphere, g.Gems, g.IsX)
		} else if g.Type == "out" {
			drawer.draw_out(g.Sphere, g.Gems, g.IsX)
		} else {
			panic(fmt.Sprintf("Invalid gem: %s", g))
		}
	}

	return drawer.Finish()
}

func draw_challenge_svg(challenge ChallengeDraw, scale float32) string {
	drawer := NewDrawer((sphereWidth * 5.0 / 3.0) + lineWidth*2)
	drawer.draw_challenge(challenge.Spheres)
	return drawer.Finish()
}
