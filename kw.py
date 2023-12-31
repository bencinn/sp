list = [

    "and",
    "array",
    "begin",
    "case",
    "const",
    "div",
    "do",
    "downto",
    "else",
    "end",
    "file",
    "for",
    "function",
    "goto",
    "if",
    "in",
    "label",
    "mod",
    "nil",
    "not",
    "of",
    "or",
    "packed",
    "procedure",
    "program",
    "record",
    "repeat",
    "set",
    "then",
    "to",
    "type",
    "until",
    "var",
    "while",
    "with",
]

for i in list:
    print(f"case \"{i}\":")
    print(f"    ttype = kw_{i}")

print("-------")
for id, type in enumerate(list):
    print(f"case kw_{type}:")
    print(f"    return \"{type}\"")

