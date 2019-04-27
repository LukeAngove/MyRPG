# Mechanics

## Motion
A motion is decribed by one or more of:
- Move some number of gems into some sphere: \<letter> -> (\<letter>)
- Have some number of gems in a sphere: \<letter>:(\<letter>)
- Move some number of gems out of some sphere: (\<letter>)->\<letter>

Where the letters are:

| letter | Element      | Colour         |
| ------ | ------------ | -------------- |
| a      | Any          | Any            |
| i      | Any Internal | Blue or Purple |
| t      | Any External | Red or Brown   |
| f      | Fire         | Red            |
| w      | Water        | Blue           |
| e      | Earth        | Brown          |
| l      | Lightning    | Yellow         |

Costs are separated by a '|'.
So for example, a cost of:
**a -> (f) | w:(e) | (e) -> l**

Would mean:
- Move any gem into the Fire Sphere
- One Water gem must be in the Earth Sphere
- One Lightning gem must be moved out of the Earth Sphere

Multiple gems can be listed at once, for example:
**ae -> (l)**

Would mean:
Move an Earth and any other colour gem into the Lightning Sphere.

Multiple spheres must be separate entries, e.g.:
**a -> (f) | a -> (w)**

Multiple spheres of the same kind must be listed, e.g.:
**aa -> (f)**

To move 2 of any gem into the Fire sphere.

Some abilities may have different effects, depending on
the number of gems used. This is denoted using an 'x', 
so an ability that allows the player to move some number of
Lightning gems out of the Fire Sphere would be written as:
**(f) -> Xl**

There are rare cases when all gems of a type must be moved out of a sphere.
This is denoted with a capital 'A', e.g.:
**(f) -> Al**

Note that motions in and out CAN be chained; if one action requires moving gems out of a sphere, and they move out into another sphere, and another action requires move the same gems into that same sphere, then the actions can count the same motion for both criteria.

Motions in and out CANNOT be counted twice; if two actions require the same movement (e.g.: both require moving a red gem out of the yellow sphere), then both actions need to be taken independantly.

A note on gems in spheres (**a:(a)); this is used in *sustained*
cards. Sustained cards are kept active constanly, and their
cost is ongoing. When considering cost, sustained cards
are considered to be ALWAYS played; gems used to pay the
cost for a sustained card cannot be used for any other purpose.
Movements (**a->(a)** or **(a)->a**) may be costs on sustained
cards, but these costs are always a one-off, at the time of
playing the card. Sustained costs are only associated with
**a:(a)**.

For example, if there are two cards with costs:
- **a -> (w) | ww:(w)**
- **l -> (w) | w:(w) | (l) -> f**

Then the minimum gems and movements required to play both cards would be:
- Move one yellow gem into the fire sphere for the second card
- Have 3 blue gems in the blue sphere
- Move one fire gem from the yellow sphere to the fire sphere

Notice that a total of 5 gems are noted here, even though 6
are used in the cards. The fire gem being moved out of the
lightning sphere can also count as the 'any' gem being moved
into the water sphere, if the player choses to do so, for the
first card. This is because we are making uses of 2 different
aspects of the movement. We could NOT use the movement of the
lightning gem into the water sphere in the second card to count
as the 'any' into the water sphere for the first card, as this
would be counting the same action twice. Also note that we
need 3 water gems in the water sphere. We cannot count gems
kept in spheres for two cards. Even if we moved a
water gem for the 'any' gem, into water sphere, we could not
count it for any of the three water gems in the water sphere
when we are chaining. This is because the cards are being
activated simultaneously. If we were to activate the first
card, completely resolve it's action, and then activate the
second, then we would need a minimum of only 2 water gems in
the water sphere, as once the first card is resolved, both
gems become 'available' for use again.

