#!/usr/bin/env python3

def test_deck_draw():
    from cards import Card, Deck
    from elements import ElementRefs as E
    card = Card("test", {E.lightning:1})
    deck = Deck(init=[card])

    expected = [card]
    actual = deck.draw(1)

    assert(actual == expected)

def test_deck_2_card_draw():
    from cards import Card, Deck
    from elements import ElementRefs as E
    card1 = Card("test1", {E.lightning:1})
    card2 = Card("test2", {E.fire:1})
    deck = Deck(init=[card1, card2])

    expected = [card1]
    actual = deck.draw(1)
    assert(actual == expected)

    expected = [card2]
    actual = deck.draw(1)
    assert(actual == expected)

def test_deck_2_card_draw_2():
    from cards import Card, Deck
    from elements import ElementRefs as E
    card1 = Card("test1", {E.lightning:1})
    card2 = Card("test2", {E.fire:1})
    deck = Deck(init=[card1, card2])

    expected = [card1, card2]
    actual = deck.draw(2)
    assert(actual == expected)
