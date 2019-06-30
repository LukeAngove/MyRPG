#!/usr/bin/env python3

def set_in(tokens):
    from costs import CostType as CT
    tokens.type = CT.IN
    return tokens

def set_out(tokens):
    from costs import CostType as CT
    tokens.type = CT.OUT
    return tokens

def set_sustain(tokens):
    from costs import CostType as CT
    tokens.type = CT.SUSTAIN
    return tokens

def sphere_to_element(tokens):
    from elements import shorthand as E
    tokens = [E.sh[tokens[0]]]
    return tokens

def beads_to_elements(tokens):
    from elements import shorthand as E
    res = []
    for t in tokens[0]:
        res.append(E.sh[t])
    tokens = [res]
    return tokens

def make_parser():
    from pyparsing import Word, Suppress
    from elements import shorthand as E

    io = Suppress("->")
    sustain_pos = Suppress(":")
    element_refs = str(list(E.symbols()))
    beads = Word(element_refs)('beads').setParseAction(beads_to_elements)
    sphere = Suppress("(") + Word(element_refs, exact=1)('sphere').setParseAction(sphere_to_element) + Suppress(")")
    move_in = beads + io + sphere
    move_out = sphere + io + beads
    sustain = beads + sustain_pos + sphere

    move_in.setParseAction(set_in)
    move_out.setParseAction(set_out)
    sustain.setParseAction(set_sustain)

    cost = move_in | move_out | sustain

    return cost

_parser = make_parser()

def parse_cost(text):
    from costs import SingleCost

    splits = text.split("|")

    costs = []
    for s in splits:
        tokens = _parser.parseString(s)
        costs.append(SingleCost(tokens.type, tokens.sphere, tokens.beads))
    return costs