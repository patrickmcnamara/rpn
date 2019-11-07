package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scnr := bufio.NewScanner(os.Stdin)
	stck := make(stack, 0)

	chk := func(err error) {
		if err != nil {
			fmt.Fprintln(os.Stderr, fmt.Errorf("rpn: %w", err))
		}
	}

	for scnr.Scan() {
		input := scnr.Text()
		if input == "" || input[0] == '#' {
			continue
		} else if cmd, ok := commands[input]; ok {
			chk(cmd(&stck))
		} else if unOp, ok := unaryOperators[input]; ok {
			chk(unaryOperator(&stck, unOp))
		} else if binOp, ok := binaryOperators[input]; ok {
			chk(binaryOperator(&stck, binOp))
		} else {
			chk(pushNumber(&stck, input))
		}
	}
}
