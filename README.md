# RPN calculator

![latest version](https://img.shields.io/github/v/tag/patrickmcnamara/rpn?label=latest%20version)
![last commit](https://img.shields.io/github/last-commit/patrickmcnamara/rpn)
![top language](https://img.shields.io/github/languages/top/patrickmcnamara/rpn)
![license](https://img.shields.io/github/license/patrickmcnamara/rpn?label=license)

I wrote this because I don't like `dc`. This is far worse, however.

It operates by pushing and popping numbers on a stack.

## Installation

Run `go get -u github.com/patrickmcnamara/rpn`.

## Usage

Run `rpn`, assuming your $GOPATH is in your $PATH, and you can begin entering
numbers, commands and operations.

### Commands

| input | command        | description                         |
| ----- | -------------- | ----------------------------------- |
| _NUM_ | number         | pushes number to top of stack       |
| `f`   | fraction       | prints the top number as a fraction |
| `d`   | decimal        | prints the top number as a decimal  |
| `sf`  | stack fraction | prints the stack as fractions       |
| `sd`  | stack decimal  | prints the stack as decimals        |
| `P`   | pop            | pops the top number                 |
| `D`   | duplicate      | duplicates the top number           |
| `R`   | reverse        | reverses the top two numbers        |
| `C`   | clear          | clears the stack                    |
| `S`   | save           | saves the stack to disk             |
| `L`   | load           | loads the stack from disk           |

The `S` and `L` commands save and load a file called `rpn` in the default cache
directory for the current user. This is typically `~/.config/` for GNU/Linux.

### Operations

There are two kinds of supported operations, unary and binary. Unary operations
pop one number off the stack as an operand. Binary operations pop two numbers
off the stack as operands.

#### Unary operations

| input | operation | description          |
| ----- | --------- | -------------------- |
| `n`   | negate    | negates the number   |
| `a`   | absolute  | absolutes the number |
| `i`   | inverse   | inverses the number  |

#### Binary operations

| input | operation | description            |
| ----- | --------- | ---------------------- |
| `+`   | add       | adds the numbers       |
| `-`   | subtract  | subtracts the numbers  |
| `*`   | multiply  | multiplies the numbers |
| `/`   | divide    | divides the numbers    |

## License

This project is licensed under the GNU General Public License v3.0.
