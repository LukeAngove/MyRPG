#!/usr/bin/env python3
import pytest
import player
from elements import ElementRefs as E

# Actions that a player can do on their turn

# Initialize player state
def test_player_initial_state():
    b = player.Board()
    for e in E.elements():
        assert(b.beads_in(e, e) == 1)

# Initialize player state, check for 'any'
def test_player_initial_state_any():
    b = player.Board()
    for e in E.elements():
        assert(b.beads_in(e, E.any) == 1)

# Add beads to a sphere
def test_add_beads():
    b = player.Board()
    from elements import BeadCounter as BC
    b.add(E.lightning, {E.lightning:2})
    expected = {E.lightning:BC({E.lightning:3}),
                E.fire:BC({E.fire:1}),
                E.earth:BC({E.earth:1}),
                E.water:BC({E.water:1}),
                E.spirit:BC({E.spirit:1})}
    assert(b.spheres == expected)


# Move beads from one sphere to another
def test_move_beads():
    from player import Board
    from elements import Sphere, BeadCounter as BC, SingleExplicitMoveCost

    b = Board()
    move = SingleExplicitMoveCost(E.lightning, E.fire, BC({E.lightning:1}))
    b.move(move)

    expected = Board(state={E.lightning:Sphere(), E.fire:Sphere({E.lightning:1, E.fire:1})}, init=1)

    assert(b == expected)

# Move beads from one sphere to another, when not enough beads of type given in source
def test_move_beads_not_enough():
    from player import Board
    from elements import BeadCounter as BC
    from elements import NotEnoughBeadsException, SingleExplicitMoveCost

    b = Board()
    move = SingleExplicitMoveCost(E.lightning, E.fire, BC({E.lightning:2}))

    with pytest.raises(NotEnoughBeadsException):
        b.move(move)

# Move beads from one sphere to another, when too many beads in target 
def test_move_beads_not_enough_of_element():
    from player import Board
    from elements import BeadCounter as BC
    from elements import NotEnoughBeadsException, SingleExplicitMoveCost

    b = Board()
    b.add(E.lightning, BC({E.lightning:3, E.fire:1}))
    move = SingleExplicitMoveCost(E.lightning, E.fire, BC({E.fire:2}))

    with pytest.raises(NotEnoughBeadsException):
        b.move(move)

# Move beads from one sphere to another, when not enough of an element
def test_move_beads_too_many():
    from elements import BeadCounter as BC
    from elements import TooManyBeadsException, SingleExplicitMoveCost

    b = player.Board()
    b.add(E.lightning, BC({E.lightning:4}))
    move = SingleExplicitMoveCost(E.lightning, E.fire, BC({E.lightning:5}))
    with pytest.raises(TooManyBeadsException):
        b.move(move)

# Play a card
def test_play_card():
    from player import Card, Deck, Player, Board
    from elements import SingleExplicitMoveCost, SingleCost, BeadCounter as BC, CostType as CT

    card = Card("test", SingleCost(CT.OUT, E.lightning, BC()))
    d = Deck([card])
    p = Player(d)

    p.draw(1)
    payed_cost = SingleExplicitMoveCost(E.lightning, E.fire, BC())
    print(payed_cost)
    res_card = p.play(0, payed_cost)

    expected_board = Board(init=1)
    expected_card = card

    assert(p.board == expected_board)
    assert(res_card == expected_card)

# Play a card that moves beads
def test_play_card_with_beads():
    from player import Card, Deck, Player, Board
    from elements import SingleCost, Sphere, BeadCounter as BC, CostType as CT
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

def test_deck_draw():
    card = player.Card("test", {E.lightning:1})
    deck = player.Deck(init=[card])

    expected = [card]
    actual = deck.draw(1)

    assert(actual == expected)

def test_deck_2_card_draw():
    card1 = player.Card("test1", {E.lightning:1})
    card2 = player.Card("test2", {E.fire:1})
    deck = player.Deck(init=[card1, card2])

    expected = [card1]
    actual = deck.draw(1)
    assert(actual == expected)

    expected = [card2]
    actual = deck.draw(1)
    assert(actual == expected)

def test_deck_2_card_draw_2():
    card1 = player.Card("test1", {E.lightning:1})
    card2 = player.Card("test2", {E.fire:1})
    deck = player.Deck(init=[card1, card2])

    expected = [card1, card2]
    actual = deck.draw(2)
    assert(actual == expected)
