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

    @staticmethod
    def elements():
        return [ElementRefs.lightning, ElementRefs.fire, ElementRefs.earth, ElementRefs.water, ElementRefs.spirit]

    @staticmethod
    def ideas():
        return [ElementRefs.external, ElementRefs.internal, ElementRefs.any]

class BeadCounter:
    MAX_BEADS = 5

    def __init__(self, state=None, init=0):
        from elements import ElementRefs as E
        self.beads = {e:init for e in E.elements()}
        if state is not None:
            self.beads.update(state)

    def items(self):
        return self.beads.items()

    def __getitem__(self, key):
        from elements import ElementRefs as E
        assert(key in E)
        if key in E.ideas():
            return sum([v for b,v in self.beads.items() if b & key])
        else:
            return self.beads[key]

    def add(self, key, value=1):
        from elements import ElementRefs as E
        assert(value > 0)
        assert(key in E.elements())
        self.modifyCheck(value)
        self._unchecked_modify(key, value)

    def modifyCheck(self, value):
        from elements import ElementRefs as E
        count = self[E.any]
        if count + value > BeadCounter.MAX_BEADS:
            raise TooManyBeadsException()
        elif count + value < 0:
            raise NotEnoughBeadsException()
        else:
            return True
 
    def _unchecked_modify(self, key, delta):
        self[key] += delta

    def remove(self, key, value=1):
        from elements import ElementRefs as E
        assert(value > 0)
        assert(key in E.elements())
        self.modifyCheck(-value)
        self._unchecked_modify(key, -value)

    def __setitem__(self, key, val):
        from elements import ElementRefs as E
        assert(key in E.elements())
        self.beads[key] = val

    def __eq__(self, other):
        return self.beads == other.beads

class NotEnoughBeadsException(Exception):
    pass

class TooManyBeadsException(Exception):
    pass