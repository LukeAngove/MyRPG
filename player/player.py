#!/usr/bin/env python3

class Board:
    '''The player state: board and cards.'''

    def __init__(self, state=None, init=1):
        from elements import ElementRefs as E
        from elements import BeadCounter as BC
        self.spheres = {e:BC({e:init}) for e in E.elements()}
        if state:
            self.spheres.update(state)

    def __eq__(self, other):
        return self.spheres == other.spheres

    def beads_in(self, sphere, beads):
        '''Count the beads of the given type in the given sphere.'''
        return self.spheres[sphere][beads]

    def add(self, sphere, beads):
        for b,v in beads.items():
            self.spheres[sphere][b] += v

    def move(self, source, target, beads):
        for b,v in beads.items():
            from elements import NotEnoughBeadsException, TooManyBeadsException, BeadCounter, ElementRefs as E
            # We need to check both here first, so that the transaction is only carried out as a whole.
            self.spheres[source].modifyCheck(-v)
            self.spheres[target].modifyCheck(v)
            self.spheres[source].remove(b, v)
            self.spheres[target].add(b, v)

