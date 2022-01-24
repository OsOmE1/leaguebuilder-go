from xxhash import xxh64_intdigest

def key_to_hash(key, bits=40):
    if isinstance(key, str):
        key = xxh64_intdigest(key)
    return key & ((1 << bits) - 1)

# 5a7d965023
tests = ["Spell_VexE_Name", "Spell_VexE_Summary", "Spell_VexE_Tooltip", "Spell_VexE_TooltipExtended"]
for t in tests:
    h = key_to_hash(t)
    h2 = key_to_hash(t.lower())
    print(f"{t}: {h:010x}")
    print(f"{t.lower()}: {h2:010x}")
