#!/usr/bin/env python3

class Board:
    '''The player state: board and cards.'''

    def __init__(self, state=None, init=1):
        from elements import ElementRefs as E

        self.spheres = {e:Sphere({e:init}) for e in E.elements()}
        if state:
            self.spheres.update(state)
            for e,s in self.spheres.items():
                assert(isinstance(s, Sphere))

    def __repr__(self):
        return "Board: {}".format(self.spheres)

    def __eq__(self, other):
        return self.spheres == other.spheres

    def beads_in(self, sphere, beads):
        '''Count the beads of the given type in the given sphere.'''
        return self.spheres[sphere][beads]

    def add(self, sphere, beads):
        self.spheres[sphere].addAll(beads)

    def move(self, cost):
        # We need to check both here first, so that the transaction is only carried out as a whole.
        self.spheres[cost.source].modifyCheckAll(cost.beads.negate())
        self.spheres[cost.dest].modifyCheckAll(cost.beads)

        # Checks are done, so execute the actual move
        self.spheres[cost.source].removeAll(cost.beads)
        self.spheres[cost.dest].addAll(cost.beads)

        return True

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

