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
    from elements import shorthand as E

    b = Board()
    b.add(E.l, {E.l:2})
    expected = {E.l:BC({E.l:3}),
                E.f:BC({E.f:1}),
                E.e:BC({E.e:1}),
                E.w:BC({E.w:1}),
                E.s:BC({E.s:1})}
    assert(b.spheres == expected)


# Move beads from one sphere to another
def test_move_beads():
    from beads import Board, Sphere, BeadCounter as BC
    from costs import SingleExplicitMoveCost
    from elements import shorthand as E

    b = Board()
    move = SingleExplicitMoveCost(E.l, E.f, BC({E.l:1}))
    b.move(move)

    expected = Board(state={E.l:Sphere(), E.f:Sphere({E.l:1, E.f:1})}, init=1)

    assert(b == expected)

# Move beads from one sphere to another, when not enough beads of type given in source
def test_move_beads_not_enough():
    from beads import Board, NotEnoughBeadsException, BeadCounter as BC
    from costs import SingleExplicitMoveCost
    from elements import shorthand as E

    b = Board()
    move = SingleExplicitMoveCost(E.l, E.f, BC({E.l:2}))

    with pytest.raises(NotEnoughBeadsException):
        b.move(move)

# Move beads from one sphere to another, when too many beads in target 
def test_move_beads_not_enough_of_element():
    from beads import Board, BeadCounter as BC, NotEnoughBeadsException
    from costs import SingleExplicitMoveCost
    from elements import shorthand as E

    b = Board()
    b.add(E.l, BC({E.l:3, E.f:1}))
    move = SingleExplicitMoveCost(E.l, E.f, BC({E.f:2}))

    with pytest.raises(NotEnoughBeadsException):
        b.move(move)

# Move beads from one sphere to another, when not enough of an element
def test_move_beads_too_many():
    from beads import TooManyBeadsException, Board, BeadCounter as BC
    from costs import SingleExplicitMoveCost
    from elements import shorthand as E

    b = Board()
    b.add(E.l, BC({E.l:4}))
    move = SingleExplicitMoveCost(E.l, E.f, BC({E.l:5}))
    with pytest.raises(TooManyBeadsException):
        b.move(move)

