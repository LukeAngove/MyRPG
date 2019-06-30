#!/usr/bin/env python3

def test_load_cards():
    cards_yaml = '''
- title: Flame
  tags: [Physical, Fire, Spell, Attack, Aggressive]
  cost: "cost{(f)->a}"
  rules:
    Test: challenge{ffw}
    Threshold: Target's PD
    Target: Single
    Range: Line of Sight
    Damage: SM + cost{A:(f)}
  flavor: "Not sure about long descriptions yet"
- title: Pain
  tags: [Mental, Spell, Attack, Aggressive]
  cost: "cost{(w)->a}"
  rules:
    Test: challenge{wwl}
    Threshold: Target's MD
    Target: Single
    Range: Line of Sight
    Damage: SM + cost{A:(w)}
  flavor: "Not sure about long descriptions yet"'''
    from cards import load_cards, Card
    from costs import SingleCost as SC, CostType as CT
    from elements import shorthand as E

    actual = load_cards(cards_yaml)
    expected = [
        Card("Flame", cost=[SC(CT.OUT, E.f, [E.a])], tags={"Physical", "Fire", "Spell", "Attack", "Aggressive"},
          rules={"Test": "challenge{ffw}",
                 "Threshold": "Target's PD",
                 "Target": "Single",
                 "Range": "Line of Sight",
                 "Damage": "SM + cost{A:(f)}"},
          flavor= "Not sure about long descriptions yet"),
        Card("Pain", cost=[SC(CT.OUT, E.w, [E.a])], tags={"Mental", "Spell", "Attack", "Aggressive"},
          rules={"Test": "challenge{wwl}",
                 "Threshold": "Target's MD",
                 "Target": "Single",
                 "Range": "Line of Sight",
                 "Damage": "SM + cost{A:(w)}"},
          flavor="Not sure about long descriptions yet"),
    ]

    assert(actual == expected)

def test_create_card():
    from cards import Card

    card_def = {
        "title":"A card",
        "cost":[],
        "tags":[],
        "rules":{}
        }

    expected = Card(card_def["title"], card_def["cost"], [], {})

    actual = Card(**card_def)

    assert(actual == expected)

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
