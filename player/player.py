#!/usr/bin/env python3

class Player:
    def __init__(self, deck):
        from cards import Hand
        from beads import Board
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
       