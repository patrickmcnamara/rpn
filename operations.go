package main

import (
	"math/big"
)

var unaryOperators = map[string]func(x *big.Rat) (*big.Rat, error){
	"n": func(x *big.Rat) (*big.Rat, error) { return new(big.Rat).Neg(x), nil },
	"a": func(x *big.Rat) (*big.Rat, error) { return new(big.Rat).Abs(x), nil },
	"i": func(x *big.Rat) (*big.Rat, error) { return new(big.Rat).Inv(x), nil },
}

var binaryOperators = map[string]func(x, y *big.Rat) (*big.Rat, error){
	"+": func(x, y *big.Rat) (*big.Rat, error) { return new(big.Rat).Add(x, y), nil },
	"-": func(x, y *big.Rat) (*big.Rat, error) { return new(big.Rat).Sub(x, y), nil },
	"*": func(x, y *big.Rat) (*big.Rat, error) { return new(big.Rat).Mul(x, y), nil },
	"/": func(x, y *big.Rat) (*big.Rat, error) {
		if y.Cmp(big.NewRat(0, 1)) == 0 {
			return nil, errDivByZero
		}
		return new(big.Rat).Quo(x, y), nil
	},
}

func unaryOperator(s *stack, unOp func(x *big.Rat) (*big.Rat, error)) (err error) {
	x, err := s.Pop()
	if err != nil {
		return
	}
	n, err := unOp(x)
	if err != nil {
		s.Push(x)
		return
	}
	s.Push(n)
	return
}

func binaryOperator(s *stack, binOp func(x, y *big.Rat) (*big.Rat, error)) (err error) {
	y, err := s.Pop()
	if err != nil {
		return
	}
	x, err := s.Pop()
	if err != nil {
		s.Push(y)
		return
	}
	n, err := binOp(x, y)
	if err != nil {
		s.Push(x)
		s.Push(y)
		return
	}
	s.Push(n)
	return
}

func pushNumber(s *stack, input string) error {
	n, ok := new(big.Rat).SetString(input)
	if !ok {
		return errInvalidInput
	}
	s.Push(n)
	return nil
}
