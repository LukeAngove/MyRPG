#!/usr/bin/env python3

import svgwrite

def do_rotate_90(pt, centre):
    '''
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
    '''
    newPt = (centre[0]+centre[1]-pt[1], centre[1]+pt[0]-centre[0])
    return newPt

def do_symetry(pts, centre):
    pts_size = len(pts)
    start = 0
    for i in range(3):
        for i in range(pts_size):
            pts.append(do_rotate_90(pts[start+i], centre))
        start += pts_size

class DrawSymbol:
    def __init__(self, filename, scale=10):
        self.drawer = Drawer(filename, scale)
        self.pos = 0
        self.pos_move = scale*6
        self.pos_minor_move = scale*3.3
        self.height = scale*6

    def __enter__(self):
        self.drawer.__enter__()
        return self

    def __exit__(self, exc_type, exc_val, exc_tb):
        return self.drawer.__exit__(exc_type, exc_val, exc_tb)

    def increment(self):
        self.pos += self.pos_move

    def minor_increment(self):
        self.pos += self.pos_minor_move

    def draw_in(self, sphere_color, gem_colors):
        self.drawer.draw_sphere(sphere_color, position=(self.pos, 0))
        self.drawer.draw_gems(gem_colors, position=(self.pos, 0))
        self.increment()

    def draw_into(self, sphere_color, gem_colors):
        self.drawer.draw_gems(gem_colors, position=(self.pos, 0))
        self.increment()
        self.drawer.draw_arrow(position=(self.pos, 0))
        self.minor_increment()
        self.drawer.draw_sphere(sphere_color, position=(self.pos, 0))
        self.increment()

    def draw_out(self, sphere_color, gem_colors):
        self.drawer.draw_sphere(sphere_color, position=(self.pos, 0))
        self.increment()
        self.drawer.draw_arrow(position=(self.pos, 0))
        self.minor_increment()
        self.drawer.draw_gems(gem_colors, position=(self.pos, 0))
        self.increment()
    
    def draw_desc(self):
        self.drawer.draw_desc((0,0), (100,100), "This is some text", 10)

    def tostring(self):
        self.drawer.svg.attribs['height'] = self.height
        self.drawer.svg.attribs['width'] = self.pos
        return self.drawer.tostring()

class Drawer:
    def __init__(self, filename, scale=10):
        self.svg = svgwrite.Drawing(filename=filename, profile='tiny', debug=True, height=5*scale)
        self.scale=scale

    def octagon(self, color, position, corner_size, flat_size, outer_line_width):
        edging = outer_line_width / 2
        total_width = edging*2 + corner_size*2 + flat_size
        centre = (position[0]+(total_width/2), position[1]+total_width/2)

        c1 = (edging+position[0], edging + corner_size + position[1])
        c2 = (edging + corner_size+position[0], edging+position[1])
        pts_arr = [c1, c2]
        do_symetry(pts_arr, centre)
        return pts_arr

    def octagon_internal_corners(self, octagon_pts):
        for i in range(int(len(octagon_pts)/2)-1, -1 ,-1):
            if i % 2 == 1:
                octagon_pts.insert(i*2+1, (octagon_pts[i*2][1-(i%2)], octagon_pts[i*2+1][i%2]))
            else:
                octagon_pts.insert(i*2+1, (octagon_pts[i*2+1][i%2], octagon_pts[i*2][1-(i%2)]))
        return octagon_pts

    def draw_gems(self, colors, position, scale=None):
        if scale is None:
            scale = self.scale
        number = len(colors)
        if number is 1:
            self.draw_gem(colors[0], position=(position[0]+0.5*scale, position[1]+0.7*scale), scale=0.75*scale)
        elif number is 2:
            pos1 = (position[0]+0.6*scale, position[1])
            pos2 = (position[0]+0.6*scale, position[1]+1.2*scale)
            self.draw_gem(colors[0], position=pos1, scale=0.75*scale)
            self.draw_gem(colors[1], position=pos2, scale=0.75*scale)
        elif number is 3:
            pos1 = (position[0]+0.6*scale, position[1])
            pos2 = (position[0]+1.2*scale, position[1]+0.9*scale)
            pos3 = (position[0]+0.0*scale, position[1]+0.9*scale)
            # Draw 3 first, so that it's on the bottom
            self.draw_gem(colors[2], position=pos3, scale=0.75*scale)
            self.draw_gem(colors[0], position=pos1, scale=0.75*scale)
            self.draw_gem(colors[1], position=pos2, scale=0.75*scale)

    def draw_gem(self, color, position, scale=None):
        if scale is None:
            scale = self.scale
        offset = 0.57*scale # To make gem fit inside sphere
        position = (position[0]+offset, position[1]+offset+1)
        outer_line_width = 2
        corner_size = 1*scale
        flat_size = 2*scale
        edging = outer_line_width / 2

        octagon_pts = self.octagon(color, position, corner_size, flat_size, outer_line_width)
        poly = self.svg.polygon(points=octagon_pts, fill=color, stroke_width=outer_line_width, stroke="black")
        self.svg.add(poly)

        octagon_pts = self.octagon_internal_corners(octagon_pts)

        for i in range(4):
            self.svg.add(self.svg.polyline(points=octagon_pts[i*3:i*3+3], fill="none", stroke_width=edging, stroke="black"))

        self.svg.add(self.svg.rect(insert=(position[0]+edging+corner_size, position[1]+edging+corner_size), size=(flat_size, flat_size), stroke_width=edging, stroke="black", fill="none"))

    def draw_sphere(self, color, position, scale=None):
        if scale is None:
            scale = self.scale
        gem_size_factor = 2.6*scale
        centre = (position[0]+gem_size_factor, position[1]+gem_size_factor+2)
        self.svg.add(self.svg.circle(center=centre, r=gem_size_factor, fill=color, stroke="black", stroke_width=2))

    def draw_arrow(self, position, scale=None):
        if scale is None:
            scale = self.scale
        arrow_mid = [(0, 1), (2, 1)]
        arrow_point = [(1, 0), (2, 1), (1, 2)]

        position = (position[0], position[1]+1.9*scale)
        def grow(array, pos, scale):
            res = [(x*scale+pos[0], y*scale+pos[1]) for x,y in array]
            return res

        arrow_mid = grow(arrow_mid, position, scale)
        arrow_point = grow(arrow_point, position, scale)
        self.svg.add(self.svg.polyline(points=arrow_mid, fill="none", stroke_width=2, stroke="black"))
        self.svg.add(self.svg.polyline(points=arrow_point, fill="none", stroke_width=2, stroke="black"))

    def draw_desc(self, position, size, text, scale=None):
        if scale is None:
            scale = self.scale
        self.svg.add(self.svg.textArea(text=text, insert=position, size=size, font_size=scale, fill="black"))

    def __enter__(self):
        return self
    
    def tostring(self):
        return self.svg.tostring()

    def __exit__(self, exc_type, exc_val, exc_tb):
        pass

def parse_all(the_string):
    the_string = the_string.replace(" ", "")
    return [parse_single(s) for s in the_string.split("|")]

def parse_single(the_string):
    import re
    colors = {
        "a" : "grey",
        "i" : "white",
        "t" : "black",
        "l" : "yellow",
        "f" : "fire",
        "e" : "brown",
        "w" : "blue",
        "s" : "green"
    }

    colors_str = "".join(colors.keys())
    in_colors_re = re.compile("([{colors}]+):\(([{colors}])\)".format(colors=colors_str))
    into_colors_re = re.compile("([{colors}]+)->\(([{colors}])\)".format(colors=colors_str))
    out_colors_re = re.compile("\(([{colors}])\)->([{colors}]+)".format(colors=colors_str))

    groups = None
    if "->(" in the_string:
        groups = into_colors_re.search(the_string)
        scolor = colors[groups[2]]
        gcolors = [colors[c] for c in groups[1]]
        gtype = "into"
    elif ")->" in the_string:
        groups = out_colors_re.search(the_string)
        scolor = colors[groups[1]]
        gcolors = [colors[c] for c in groups[2]]
        gtype = "out"
    elif ":(" in the_string:
        groups = in_colors_re.search(the_string)
        scolor = colors[groups[2]]
        gcolors = [colors[c] for c in groups[1]]
        gtype = "in"
    else:
        print("Failed to match with {}".format(the_string))
        scolors = ""
        gcolors = [""]
        gtype = ""
    
    return {
        'args': {
            'sphere_color': scolor,
            'gem_colors': gcolors
        },
        'type': gtype
        }

def text_to_svg(in_string, scale=10):
    to_draw = parse_all(in_string)
    with DrawSymbol("not_saved_anywhere.svg", scale=scale) as gem:
        for d in to_draw:
            if d["type"] is "into":
                gem.draw_into(**d["args"])
            elif d["type"] is "in":
                gem.draw_in(**d["args"])
            elif d["type"] is "out":
               gem.draw_out(**d["args"])
        res = gem.tostring()
    return res

if __name__ == "__main__":
    print(text_to_svg("al ->  (a)|w:(e)|(i) -> w", scale=10))