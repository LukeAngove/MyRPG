#!/usr/bin/env python3

import pytest

# Actions that a player can do on their turn

# Play a card
def test_play_card():
    from cards import Card, Deck
    from player import Player
    from beads import Board, BeadCounter as BC
    from costs import SingleExplicitMoveCost, SingleCost, CostType as CT
    from elements import ElementRefs as E

    card = Card("test", SingleCost(CT.OUT, E.lightning, BC()))
    d = Deck([card])
    p = Player(d)

    p.draw(1)
    payed_cost = SingleExplicitMoveCost(E.lightning, E.fire, BC())
    res_card = p.play(0, payed_cost)

    expected_board = Board(init=1)
    expected_card = card

    assert(p.board == expected_board)
    assert(res_card == expected_card)

# Play a card that moves beads
def test_play_card_with_beads():
    from cards import Card, Deck
    from player import Player
    from beads import Board, Sphere, BeadCounter as BC
    from costs import SingleExplicitMoveCost, SingleCost, CostType as CT
    from elements import ElementRefs as E
    from helpers import createExplicitMoveCostFromSingleCost as helpCost

    card = Card("test", SingleCost(CT.OUT, E.lightning, BC({E.fire: 2})))
    d = Deck([card])
    p = Player(d)
    p.addBeads(E.lightning, BC({E.fire: 2}))

    p.draw(1)
    card_cost = p.inspect(0).cost
    try_cost = helpCost(card_cost, p.board)
    res_card = p.play(0, try_cost)

    expected_board = Board(init=1, state={E.fire: Sphere({E.fire:3})})
    expected_card = card

    assert(p.board == expected_board)
    assert(res_card == expected_card)
