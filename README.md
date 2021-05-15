# Floppa Programming Language
Created just for fun. But if you want to contribute, why not?

Floppa p.l. inspired by the brainf\*ck programming language. So you can convert your brainf\*ck code to floppa code.

# Compiling
You must have installed go before using compiler. If you have, run the `./compiler/main.go` program. Compiler takes two argument, first argument is input file (.flop file), second argument is output file (.go file). The compiler is written in golang, so you actually can compile your .flop code to machine language with using `go build <file.go>`

# Converting Brainf*ck Code to Floppa Code
You can use simple python code for that. You must have python3.6>= for use this code. Converter: `./tools/bf2floppa.py`. Same with compiler, takes two argument. First one is input file (.bf), Second one is output file (.flop).

# Hello World
Here an example hello world program written in floppa programming language (`./example/helloworld.flop`);
```flop
keeerr keerr flop go keeerr keeerr keer no fo hoe go keer floppa 
no keerr flop go keeerr keeerr keer no fo hoe go fo floppa keeerr 
keerr floppa floppa keerr floppa no keeerr keerr flop go foooo fo 
foooo foo no fo hoe go foo floppa no keeerr keeerr keerr flop go 
keeerr ker no fo hoe go floppa no keeerr keer flop go keeerr no 
fo hoe go floppa keerr floppa foooo foo floppa foooo foooo floppa 
no ker keeerr ker flop go foooo foooo fooo no fo hoe go fo flop 
no ker go fo hoe no floppa fooo flop foooo foooo fooo hoe
```

# Commands
```
ker => Add one byte to current cell.
fo => Remove one byte from current cell.
go => Goto next cell.
no => Goto last cell.
flop => Start loop. (will continue until no byte left.)
hoe => Loop finish.
caracal => Get one byte input.
floppa => Print current line as ascii.

These commands added for help while adding / removing byte:
keer => Add two byte to current cell.
foo => Remove two byte from current cell.
keerr => Add three byte to current cell.
fooo => Remove three byte from current cell.
keeerr => Add four byte to current cell.
foooo => Remove four byte from current cell.
```
For brainf*ck;
```py
{
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
```

# What Can I Do Whit This?
In theory, you can do everything. If you are masochist, try write a floppa programming language interpreted that created with floppa programming language!