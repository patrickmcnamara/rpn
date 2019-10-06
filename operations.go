package main

import "math/big"

var unaryOperators = map[string]func(x *big.Rat) *big.Rat{
	"n": func(x *big.Rat) *big.Rat { return new(big.Rat).Neg(x) },
	"a": func(x *big.Rat) *big.Rat { return new(big.Rat).Abs(x) },
	"i": func(x *big.Rat) *big.Rat { return new(big.Rat).Inv(x) },
}

var binaryOperators = map[string]func(x, y *big.Rat) *big.Rat{
	"+": func(x, y *big.Rat) *big.Rat { return new(big.Rat).Add(x, y) },
	"-": func(x, y *big.Rat) *big.Rat { return new(big.Rat).Sub(x, y) },
	"*": func(x, y *big.Rat) *big.Rat { return new(big.Rat).Mul(x, y) },
	"/": func(x, y *big.Rat) *big.Rat { return new(big.Rat).Quo(x, y) },
}

func unaryOperator(s *stack, unOp func(x *big.Rat) *big.Rat) error {
	x, err := s.Pop()
	if err != nil {
		return err
	}
	s.Push(unOp(x))
	return nil
}

func binaryOperator(s *stack, binOp func(x, y *big.Rat) *big.Rat) (err error) {
	defer func() {
		if recover() != nil {
			err = errDivByZero
		}
	}()
	y, err := s.Pop()
	if err != nil {
		return err
	}
	x, err := s.Pop()
	if err != nil {
		return err
	}
	s.Push(binOp(x, y))
	return err
}

func pushNumber(s *stack, input string) error {
	n, ok := new(big.Rat).SetString(input)
	if !ok {
		return errInvalidInput
	}
	s.Push(n)
	return nil
}
