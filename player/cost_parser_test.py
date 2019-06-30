#!/usr/bin/env python3

def test_any_in_any():
    from costs import SingleCost, CostType as CT
    from elements import shorthand as E
    from cost_parser import parse_cost

    text = "a->(a)"
    expected = [SingleCost(CT.IN, E.a, [E.a])]

    actual = parse_cost(text)

    assert(actual == expected)

def test_firefire_in_water():
    from costs import SingleCost, CostType as CT
    from elements import shorthand as E
    from cost_parser import parse_cost

    text = "ff->(w)"
    expected = [SingleCost(CT.IN, E.w, [E.f, E.f])]

    actual = parse_cost(text)

    assert(actual == expected)

def test_earthspirit_out_inner():
    from costs import SingleCost, CostType as CT
    from elements import shorthand as E
    from cost_parser import parse_cost

    text = "(i)->es"
    expected = [SingleCost(CT.OUT, E.i, [E.e, E.s])]

    actual = parse_cost(text)

    assert(actual == expected)

