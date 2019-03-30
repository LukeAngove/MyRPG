# Mechanics

## Motion
A motion is decribed by one or more of:
- Move some number of gems into some sphere: \<letter> -> (\<letter>)
- Have some number of gems in a sphere: \<letter>:(\<letter>)
- Move some number of gems out of some sphere: (\<letter>)->\<letter>

Where the letters are:

| letter | Element | Colour |
|-|-|-|
| a | Any          | Any    |
| f | Fire         | Red    |
| w | Water        | Blue   |
| e | Earth        | Brown  |
| l | Lightning    | Yellow |

Costs are separated by a '|'.
So for example, a cost of:
**a -> (r) | w:(e) | (e) -> y**

Would mean:
- Move any gem into the Fire Sphere
- One Water gem must be in the Earth Sphere
- One Lightning gem must be moved out of the Earth Sphere

Multiple gems can be listed at once, for example:
**a,e -> (y)**

Would mean:
Move an Earth and any other colour gem into the Lightning Sphere.

Multiple spheres must be separate entries, e.g.:
**a -> (f) | a -> (w)**

Multiple spheres of the same kind must be listed, e.g.:
**a,a -> (f)**

To move 2 of any gem into the Fire sphere.

Some abilities may have different effects, depending on
the number of gems used. This is denoted using an 'x', 
so an ability that allows the player to move some number of
Lightning gems out of the Fire Sphere would be written as:
**(f) -> xl**

Note that motions in and out CAN be chained; if one action requires moving gems out of a sphere, and they move out into another sphere, and another action requires move the same gems into that same sphere, then the actions can count the same motion for both criteria.

Motions in and out CANNOT be counted twice; if two actions require the same movement (e.g.: both require moving a red gem out of the yellow sphere), then both actions need to be taken independantly.

