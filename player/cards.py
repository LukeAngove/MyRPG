#!/usr/bin/env python3

class Card:
    def __init__(self, name, cost):
        self.name = name
        self.cost = cost

    def __repr__(self):
        return "Card: {{name: \"{}\", cost: {}}}".format(self.name, self.cost)

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

