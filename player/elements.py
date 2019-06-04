#!/usr/bin/env python3

from enum import Flag, auto

class ElementRefs(Flag):
    lightning = auto()
    fire = auto()
    earth = auto()
    water = auto()
    spirit = auto()
    external = fire | earth
    internal = water | spirit
    any = lightning | fire | earth | water | spirit
    connected_lightning = fire | earth | water | spirit
    connected_fire = lightning | spirit
    connected_earth = lightning | water
    connected_water = lightning | earth
    connected_spirit = lightning | fire

    @staticmethod
    def elements():
        return [ElementRefs.lightning, ElementRefs.fire, ElementRefs.earth, ElementRefs.water, ElementRefs.spirit]

    @staticmethod
    def ideas():
        return [ElementRefs.external, ElementRefs.internal, ElementRefs.any]

    @staticmethod
    def getConnected(e):
        connected = {
            ElementRefs.lightning: ElementRefs.connected_lightning,
            ElementRefs.fire: ElementRefs.connected_fire,
            ElementRefs.earth: ElementRefs.connected_earth,
            ElementRefs.water: ElementRefs.connected_water,
            ElementRefs.spirit: ElementRefs.connected_spirit,
        }
        return connected[e]

class BeadCounter:
    def __init__(self, state=None, init=0):
        from elements import ElementRefs as E
        self.beads = {e:init for e in E.elements()}
        if state is not None:
            self.beads.update(state)

    def __repr__(self):
        return "BeadCounter:{}".format(self.beads)

    def meets(self, other):
        for k,v in other.items():
            if self[k] < v:
                return False
        return True

    def negate(self):
        new_beads = BeadCounter()
        for b,v in self.items():
            new_beads[b] = -v
        return new_beads

    def items(self):
        return self.beads.items()

    def __getitem__(self, key):
        from elements import ElementRefs as E
        assert(key in E)
        if key in E.ideas():
            return sum([v for b,v in self.beads.items() if b & key])
        else:
            return self.beads[key]

    def modifyCheckAll(self, beads):
        for b,v in beads.items():
            self.modifyCheckSingle(b, v)
        return True

    def modifyCheckSingle(self, key, value):
        from elements import ElementRefs as E
        count = self[E.any]
        # Check that we have at least 0 beads of given element after change
        if self[key] + value < 0:
            raise NotEnoughBeadsException()
        elif count + value < 0:
            raise NotEnoughBeadsException()
        else:
            return True
 
    def _unchecked_modify(self, key, delta):
        self[key] += delta

    def addAll(self, beads):
        for b,v in beads.items():
            self.addSingle(b, v)

    def addSingle(self, key, value=1):
        from elements import ElementRefs as E
        assert(value >= 0)
        assert(key in E.elements())
        self.modifyCheckSingle(key, value)
        self._unchecked_modify(key, value)

    def removeAll(self, beads):
        for b,v in beads.items():
            self.removeSingle(b,v)

    def removeSingle(self, key, value=1):
        from elements import ElementRefs as E
        assert(value >= 0)
        assert(key in E.elements())
        self.modifyCheckSingle(key, -value)
        self._unchecked_modify(key, -value)

    def __setitem__(self, key, val):
        from elements import ElementRefs as E
        assert(key in E.elements())
        self.beads[key] = val

    def __eq__(self, other):
        return self.beads == other.beads

class Sphere(BeadCounter):
    MAX_BEADS = 5

    def __repr__(self):
        out = {e.name: v for e,v in self.items()}
        return "Sphere: L:({lightning}) F:({fire}) E:({earth}) W:({water}) S:({spirit})".format(**out)

    def modifyCheckSingle(self, key, value):
        super().modifyCheckSingle(key, value)
        from elements import ElementRefs as E

        count = self[E.any]
        if count + value > Sphere.MAX_BEADS:
            raise TooManyBeadsException()
        return True

class NotEnoughBeadsException(Exception):
    pass

class TooManyBeadsException(Exception):
    pass

class CostType(Flag):
    SUSTAIN = auto()
    OUT = auto()
    IN = auto()
    MOVE = IN | OUT

from collections import namedtuple

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
    assert(cost.type & CostType.MOVE)
    if cost.type == CostType.OUT:
        return SingleExplicitMoveCost(cost.sphere, ElementRefs.getConnected(cost.sphere), cost.beads)
    else:
        assert(cost.type == CostType.IN)
        return SingleExplicitMoveCost(ElementRefs.getConnected(cost.sphere), cost.sphere, cost.beads)
