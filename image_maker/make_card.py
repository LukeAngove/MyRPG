#!/usr/bin/env python3

def make_card_entry(image, name, desc):
    return {
        'image': text_to_svg(image, scale=4),
        'desc': desc
    }

# Capture our current directory
from os import path
THIS_DIR = path.dirname(path.abspath(__file__))

def make_card(title, image, cost, rules, flavor_text):
    from make_gem import text_to_svg
    from jinja2 import Environment, FileSystemLoader
    env = Environment(loader=FileSystemLoader(THIS_DIR),
                         trim_blocks=True)
    return env.get_template('default.html.j2').render(
        title=title,
        image=image,
        cost=text_to_svg(cost, scale=4),
        rules=rules,
        flavor_text=flavor_text
    )

if __name__ == "__main__":
    print(make_card("Some Spell",
                    "../board.svg",
                    "al->(a)|w:(e)|(i)->w",
                    "This where the rules go",
                    "Here's a long description"))
