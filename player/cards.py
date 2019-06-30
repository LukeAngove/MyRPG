#!/usr/bin/env python3

class Card:
    def __init__(self, title=None, cost=None, tags=None, rules=None, flavor=None):
        self.name = title
        self.cost = cost
        self.rules = rules
        self.flavor = flavor

    def __repr__(self):
        return "Card: {{name: \"{}\", cost: {}}}".format(self.name, self.cost)

    def __eq__(self, other):
        return self.name == other.name and self.cost == other.cost and self.rules == other.rules and self.flavor == other.flavor

def load_cards(yaml_text):
    import yaml
    from cost_parser import parse_cost
    import re
    from elements import shorthand as E

    cost_re = "cost{{([\\(\\){}\\->:,|]*)}}".format("".join(E.symbols()))
    cost_path = re.compile(cost_re)

    cards_data = yaml.load(yaml_text, Loader=yaml.SafeLoader)
    for d in cards_data:
        cost_str = cost_path.findall(d['cost'])[0]
        d['cost'] = parse_cost(cost_str)
    cards = [Card(**d) for d in cards_data]

    return cards

class Hand:
    def __init__(self):
        self.cards = []

    def __repr__(self):
        return "Hand: {}".format(self.cards)

    def show(self):
        return self.cards

    def remove(self, index):
        return self.cards.pop(index)

    def inspect(self, index):
        return self.cards[index]

    def add(self, cards):
        self.cards.extend(cards)

class Deck:
    def __init__(self, init=None):
        if init is not None:
            self.cards = init
        else:
            self.cards = []

    def __repr__(self):
        return "Deck : {}".format(self.cards)

    def shuffle(self):
        pass

    def draw(self, num=1):
        res = self.cards[0:num]
        self.cards = self.cards[num:]
        return res

    def size(self):
        return len(self.cards)

