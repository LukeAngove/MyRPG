package main

import (
	"bytes"
	"fmt"

	svg "github.com/ajstarks/svgo"
)

type Drawer struct {
	canvas    *svg.SVG
	outbuffer bytes.Buffer
	scale     float32
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

func draw_gem(canvas *svg.SVG, color string, position Pos, scale float32) {
	offset := 0.57 * scale // To make gem fit inside sphere
	position = Pos{position.x + offset, position.y + offset + 1}
	gem_size_factor := 2.0 * scale
	outer_line_width := float32(2.0)
	centre := Pos{position.x + gem_size_factor + 1, position.y + gem_size_factor + 1}
	canvas.Circle(int(centre.x), int(centre.y), int(gem_size_factor),
		fmt.Sprintf("fill:%s;stroke:%s;stroke_width:%d", color, "black", int(outer_line_width)))
}

func draw_gem_X(canvas *svg.SVG, color string, position Pos, scale float32, isX bool) {
	if isX {
		draw_gem(canvas, color, Pos{position.x - 0.5*scale, position.y - 0.5*scale}, scale*1.2)
		offset := 0.57 * scale // To make gem fit inside sphere
		style := fmt.Sprintf("font-size:%f", scale*4)
		if color == "black" {
			style = style + "fill=silver"
		}
		canvas.Text(int(position.x+offset+1.4*scale), int(position.y+offset+4*scale), "X", style)
	} else {
		draw_gem(canvas, color, position, scale)
	}
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
	arrow := []Pos{{0, 1}, {1.5, 1}, {0.5, 0}, {1.5, 0}, {3.0, 1.5}, {1.5, 3}, {0.5, 3}, {1.5, 2}, {0, 2}}

	outer_line_width := float32(2.0)

	position = Pos{position.x, position.y + 1.6*scale}
	arrow = grow(arrow, position, scale)
	pts_x, pts_y := convert_pos(arrow)
	format := fmt.Sprintf("fill:%s;stroke_width:%d;stroke:%s", "gray", int(outer_line_width), "black")
	canvas.Polygon(pts_x, pts_y, format)
}

func calc_width(size float32, scale float32) float32 {
	width := float32(0.0)
	width = 8.0 * size * scale

	// Add width for spacing
	width += 0.5 * float32(size-1)

	return width
}

func draw_gems(canvas *svg.SVG, colors []string, position Pos, scale float32, IsX bool) {
	number := len(colors)
	if len(colors) != 1 && IsX {
		panic("More than one gem color, and using X! Not a valid color set!")
	}
	switch number {
	case 1:
		pos := Pos{position.x + 0.5*scale, position.y + 0.7*scale}
		draw_gem_X(canvas, colors[0], pos, 0.75*scale, IsX)
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

func (drawer *Drawer) draw_gems(gem_colors []string, IsX bool) {
	draw_gems(drawer.canvas, gem_colors, drawer.offset, drawer.scale, IsX)
}

func (drawer *Drawer) draw_arrow() {
	draw_arrow(drawer.canvas, drawer.offset, drawer.scale)
}

func (drawer *Drawer) increment(inc float32) {
	drawer.offset.x += inc * drawer.scale
}

func (drawer *Drawer) draw_challenge(sphere_colors [3]string) {
	drawer.draw_sphere(sphere_colors[1])
	drawer.increment(4)
	drawer.draw_sphere(sphere_colors[2])
	drawer.increment(-2)
	drawer.draw_sphere(sphere_colors[0])
	drawer.increment(6)
}

func (drawer *Drawer) draw_gems_in_sphere(sphere_color string, gem_colors []string, IsX bool) {
	drawer.draw_sphere(sphere_color)
	drawer.draw_gems(gem_colors, IsX)
}

func (drawer *Drawer) draw_in(sphere_color string, gem_colors []string, IsX bool) {
	drawer.increment(1)
	drawer.draw_gems_in_sphere(sphere_color, gem_colors, IsX)
	drawer.increment(7)
}

func (drawer *Drawer) draw_into(sphere_color string, gem_colors []string, IsX bool) {
	drawer.increment(1)
	drawer.draw_gems_in_sphere(sphere_color, gem_colors, IsX)
	drawer.increment(-1)
	drawer.draw_arrow()
	drawer.increment(8)
}

func (drawer *Drawer) draw_out(sphere_color string, gem_colors []string, IsX bool) {
	drawer.increment(1)
	drawer.draw_gems_in_sphere(sphere_color, gem_colors, IsX)
	drawer.increment(4)
	drawer.draw_arrow()
	drawer.increment(3)
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

func draw_gems_svg(gems []GemDraw, scale float32) string {
	width := calc_width(float32(len(gems)), scale)

	drawer := NewDrawer(width, scale)

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
	width := calc_width(1.5, scale)

	drawer := NewDrawer(width, scale)
	drawer.draw_challenge(challenge.Spheres)
	return drawer.Finish()
}
