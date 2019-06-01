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
    b = player.Board()
    from elements import BeadCounter as BC
    b.move(E.lightning, E.fire, {E.lightning:1})
    expected = player.Board(state={E.lightning:BC(), E.fire:BC({E.lightning:1, E.fire:1})}, init=1)
    assert(b == expected)

# Move beads from one sphere to another, when not enough beads of type given in source
def test_move_beads_not_enough():
    b = player.Board()
    from elements import BeadCounter as BC
    from elements import NotEnoughBeadsException
    with pytest.raises(NotEnoughBeadsException):
        b.move(E.lightning, E.fire, {E.lightning:2})

# Move beads from one sphere to another, when too many beads in target 
def test_move_beads_too_many():
    b = player.Board()
    b.add(E.lightning, {E.lightning:4})
    from elements import BeadCounter as BC
    from elements import TooManyBeadsException
    with pytest.raises(TooManyBeadsException):
        b.move(E.lightning, E.fire, {E.lightning:5})

# Play a card
@pytest.mark.xfail
def test_play_card():
    assert(False)

# Play a card that moves beads
@pytest.mark.xfail
def test_play_card_with_beads():
    assert(False)