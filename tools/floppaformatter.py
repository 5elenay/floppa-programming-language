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

keys = input_file.split(" ")
count = 0
formatted = ""

for key in keys:
    formatted += f"{key} "
    count += 1
    if count == 16:
        formatted += "\n"
        count = 0

try:
    with open(sys.argv[2], "w+", encoding="utf-8") as f:
        f.write(formatted)
except Exception as e:
    print(f"Failed to write output file. ({e.__class__.__name__})")
    exit(2)

print("Formatted floppa code successfully, check {0} file".format(sys.argv[2]))