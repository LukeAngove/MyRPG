#!/usr/bin/env python3

def createExplicitMoveCostFromSingleCost(cost, board):
    '''Turns a single cost, which must be a move, into an explicit cost.
    Moving to lightning is preferred, if full, will chose the other
    connected sphere. If the source IS lightning, then choose first of
    fire, earth, water, spirit, whichever has space. For any, choose first
    possible of lightning, fire, earth, water, spirit. For internal, choose
    first possible of water and spirit, external, first of fire and earth.'''
    from elements import SingleExplicitMoveCost, ElementRefs as E, CostType as CT
    assert(cost.type & CT.MOVE)

    possible_connected = E.getConnected(cost.sphere)

    for e in E.elements():
        target_sphere = e & possible_connected
        if target_sphere and board.spheres[target_sphere].modifyCheckAll(cost.beads):
            if cost.type == CT.OUT:
                return SingleExplicitMoveCost(cost.sphere, target_sphere, cost.beads)
            else:
                assert(cost.type == CT.IN)
                return SingleExplicitMoveCost(target_sphere, cost.sphere, cost.beads)
    return None
