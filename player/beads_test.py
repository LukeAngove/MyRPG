#!/usr/bin/env python3
import pytest

# Initialize board state
def test_player_initial_state():
    from elements import ElementRefs as E
    from beads import Board

    b = Board()
    for e in E.elements():
        assert(b.beads_in(e, e) == 1)

# Initialize player state, check for 'any'
def test_board_initial_state_any():
    from elements import ElementRefs as E
    from beads import Board

    b = Board()
    for e in E.elements():
        assert(b.beads_in(e, E.any) == 1)

# Add beads to a sphere
def test_add_beads():
    from beads import Board, BeadCounter as BC
    from elements import ElementRefs as E

    b = Board()
    b.add(E.lightning, {E.lightning:2})
    expected = {E.lightning:BC({E.lightning:3}),
                E.fire:BC({E.fire:1}),
                E.earth:BC({E.earth:1}),
                E.water:BC({E.water:1}),
                E.spirit:BC({E.spirit:1})}
    assert(b.spheres == expected)


# Move beads from one sphere to another
def test_move_beads():
    from beads import Board, Sphere, BeadCounter as BC
    from costs import SingleExplicitMoveCost
    from elements import ElementRefs as E

    b = Board()
    move = SingleExplicitMoveCost(E.lightning, E.fire, BC({E.lightning:1}))
    b.move(move)

    expected = Board(state={E.lightning:Sphere(), E.fire:Sphere({E.lightning:1, E.fire:1})}, init=1)

    assert(b == expected)

# Move beads from one sphere to another, when not enough beads of type given in source
def test_move_beads_not_enough():
    from beads import Board, NotEnoughBeadsException, BeadCounter as BC
    from costs import SingleExplicitMoveCost
    from elements import ElementRefs as E

    b = Board()
    move = SingleExplicitMoveCost(E.lightning, E.fire, BC({E.lightning:2}))

    with pytest.raises(NotEnoughBeadsException):
        b.move(move)

# Move beads from one sphere to another, when too many beads in target 
def test_move_beads_not_enough_of_element():
    from beads import Board, BeadCounter as BC, NotEnoughBeadsException
    from costs import SingleExplicitMoveCost
    from elements import ElementRefs as E

    b = Board()
    b.add(E.lightning, BC({E.lightning:3, E.fire:1}))
    move = SingleExplicitMoveCost(E.lightning, E.fire, BC({E.fire:2}))

    with pytest.raises(NotEnoughBeadsException):
        b.move(move)

# Move beads from one sphere to another, when not enough of an element
def test_move_beads_too_many():
    from beads import TooManyBeadsException, Board, BeadCounter as BC
    from costs import SingleExplicitMoveCost
    from elements import ElementRefs as E

    b = Board()
    b.add(E.lightning, BC({E.lightning:4}))
    move = SingleExplicitMoveCost(E.lightning, E.fire, BC({E.lightning:5}))
    with pytest.raises(TooManyBeadsException):
        b.move(move)

