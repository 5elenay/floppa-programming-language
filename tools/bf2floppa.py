import re, sys

if len(sys.argv) < 3:
    print("ERROR: Can't find input or output argument.")
    exit(1)

try:
    with open(sys.argv[1], "r", encoding="utf-8") as f:
        input_file = f.read()
except Exception as e:
    print(f"Failed to load input file. ({e.__class__.__name__})")
    exit(1)

result = re.findall(r"\+\+?\+?\+?|--?-?-?|\[|\]|>|<|\.|,", input_file)

commands = {
    "+": "ker",
    "-": "fo",
    ">": "go",
    "<": "no",
    "[": "flop",
    "]": "hoe",
    ",": "caracal",
    ".": "floppa",
    "++": "keer",
    "--": "foo",
    "+++": "keerr",
    "---": "fooo",
    "++++": "keeerr",
    "----": "foooo"
}

compiled = [commands[i] for i in result]

try:
    with open(sys.argv[2], "w+", encoding="utf-8") as f:
        f.write(' '.join(compiled))
except Exception as e:
    print(f"Failed to write output file. ({e.__class__.__name__})")
    exit(2)

print("Turned brainf*ck code to floppa programming language successfully, check {0} file".format(sys.argv[2]))