#!/usr/bin/env python3

from enum import Flag, auto

class CostType(Flag):
    SUSTAIN = auto()
    OUT = auto()
    IN = auto()
    MOVE = IN | OUT

from collections import namedtuple

def cost_from_string(string):
    import re
    from elements import shorthand as E
    to_match = re.compile("cost{{[{syms}]->[{syms}]}}".compile(syms=E.symbols()))

SingleCost = namedtuple("SingleCost", ["type", "sphere", "beads"])
# Cost is an iterable of SingleCosts

class SingleExplicitMoveCost:
    def __init__(self, source, dest, beads):
        self.source = source
        self.dest = dest
        self.beads = beads

    def __repr__(self):
        return "S: {}, D: {}, B: {}".format(self.source, self.dest, self.beads)

    def meets(self, other):
        if isinstance(other, SingleCost):
            o = singleCostToSingleExplicitMoveCost(other)
        else:
            o = other
        if self.source & o.source and self.dest & o.dest and self.beads.meets(o.beads):
            return True
        else:
            return False

def singleCostToSingleExplicitMoveCost(cost):
    from elements import ElementRefs

    assert(cost.type & CostType.MOVE)
    if cost.type == CostType.OUT:
        return SingleExplicitMoveCost(cost.sphere, ElementRefs.getConnected(cost.sphere), cost.beads)
    else:
        assert(cost.type == CostType.IN)
        return SingleExplicitMoveCost(ElementRefs.getConnected(cost.sphere), cost.sphere, cost.beads)
