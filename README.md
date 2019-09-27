# RPN calculator

![latest version](https://img.shields.io/github/v/tag/patrickmcnamara/rpn?label=latest%20version)
![last commit](https://img.shields.io/github/last-commit/patrickmcnamara/rpn)
![top language](https://img.shields.io/github/languages/top/patrickmcnamara/rpn)
![licence](https://img.shields.io/github/license/patrickmcnamara/rpn?label=licence)

I wrote this because I don't like `dc`.
This is far worse, however.

It operates by pushing and popping numbers on a stack.

## Installation

Run `go get -u github.com/patrickmcnamara/rpn`.

## Usage

Run `rpn`, assuming your $GOPATH is in your $PATH, and you can begin entering commands and operations.

### Commands

| input | command   | description                             |
|-------|-----------|-----------------------------------------|
| *NUM* | number    | pushes number to top of stack           |
| `f`   | fraction  | prints the top number as a fraction     |
| `d`   | decimal   | prints the top number as a decimal      |
| `s`   | stack     | prints the stack as a list of fractions |
| `P`   | pop       | pops the top number                     |
| `D`   | duplicate | duplicates the top number               |
| `R`   | reverse   | reverses the top two numbers            |
| `C`   | clear     | clears the stack                        |
| `S`   | save      | saves the stack to disk                 |
| `L`   | load      | loads the stack from disk               |
| `Q`   | quit      | quits `rpn`                             |

The `S` and `L` commands save and load a file called `rpn` in the default cache directory for the current user.
This is typically `~/.config/` for GNU/Linux.

As well as the `Q` command, an EOF to `stdin` will exit the program.

### Operations

There are two kinds of supported operations, unary and binary.
Unary operations pop one number off the stack as an operand.
Binary operations pop two numbers off the stack as operands.

### Unary operations

| input | operation | description          |
|-------|-----------|----------------------|
| `n`   | negate    | negates the number   |
| `a`   | absolute  | absolutes the number |
| `i`   | inverse   | inverses the number  |

### Binary operations

| input | operation | description            |
|-------|-----------|------------------------|
| `+`   | add       | adds the numbers       |
| `-`   | subtract  | subtracts the numbers  |
| `*`   | multiply  | multiplies the numbers |
| `/`   | divide    | divides the numbers    |

## Licence

This project is licenced under the European Union Public Licence v1.2.
