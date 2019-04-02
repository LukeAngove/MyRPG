#!/usr/bin/env python3

def make_card_entry(image, name, desc):
    from make_gem import text_to_svg
    return {
        'image': text_to_svg(image, scale=4),
        'name': name,
        'desc': desc
    }

# Capture our current directory
from os import path
THIS_DIR = path.dirname(path.abspath(__file__))

def make_card(title, image, abilities, flavor_text):
    from jinja2 import Environment, FileSystemLoader
    env = Environment(loader=FileSystemLoader(THIS_DIR),
                         trim_blocks=True)
    return env.get_template('default.html.j2').render(
        title=title,
        image=image,
        abilities=abilities,
        flavor_text=flavor_text
    )

if __name__ == "__main__":
    print(make_card("My first card", "../board.svg", [
        make_card_entry("al->(a)|w:(e)|(i)->w",
                        "Shock",
                        "This is a thing")
        ], "Here's a long description"))
