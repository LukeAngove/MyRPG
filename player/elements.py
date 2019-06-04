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
