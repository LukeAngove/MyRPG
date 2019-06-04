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

class Board:
    '''The player state: board and cards.'''

    def __init__(self, state=None, init=1):
        from elements import Sphere, ElementRefs as E

        self.spheres = {e:Sphere({e:init}) for e in E.elements()}
        if state:
            self.spheres.update(state)
            for e,s in self.spheres.items():
                assert(isinstance(s, Sphere))

    def __repr__(self):
        return "Board: {}".format(self.spheres)

    def __eq__(self, other):
        return self.spheres == other.spheres

    def beads_in(self, sphere, beads):
        '''Count the beads of the given type in the given sphere.'''
        return self.spheres[sphere][beads]

    def add(self, sphere, beads):
        self.spheres[sphere].addAll(beads)

    def move(self, cost):
        # We need to check both here first, so that the transaction is only carried out as a whole.
        self.spheres[cost.source].modifyCheckAll(cost.beads.negate())
        self.spheres[cost.dest].modifyCheckAll(cost.beads)

        # Checks are done, so execute the actual move
        self.spheres[cost.source].removeAll(cost.beads)
        self.spheres[cost.dest].addAll(cost.beads)

        return True

class Player:
    def __init__(self, deck):
        self.deck = deck
        self.hand = Hand()
        self.board = Board()

    def addBeads(self, sphere, beads):
        self.board.add(sphere, beads)

    def draw(self, num):
        cards = self.deck.draw(num)
        self.hand.add(cards)

    def inspect(self, idx):
        return self.hand.inspect(idx)

    def play(self, idx, pay):
        cost = self.inspect(idx).cost
        assert(pay.meets(cost))
        assert(self.board.move(pay))
        return self.hand.remove(idx)
       