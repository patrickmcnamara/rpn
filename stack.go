package main

import "math/big"

type stack []*big.Rat

func (s *stack) Push(n *big.Rat) {
	*s = append(*s, n)
}

func (s *stack) Pop() (*big.Rat, error) {
	n := new(big.Rat)
	if len(*s) > 0 {
		n = (*s)[len(*s)-1]
		*s = (*s)[:len(*s)-1]
		return n, nil
	}
	return n, errPopEmptyStack
}

func (s *stack) Peek() (*big.Rat, error) {
	n := new(big.Rat)
	if len(*s) > 0 {
		n = (*s)[len(*s)-1]
		return n, nil
	}
	return n, errPeekEmptyStack
}

func (s *stack) Duplicate() error {
	n, err := s.Peek()
	if err != nil {
		return errDupEmptyStack
	}
	s.Push(n)
	return nil
}

func (s *stack) Reverse() error {
	n, err := s.Pop()
	if err != nil {
		return errRevEmptyStack
	}
	m, err := s.Pop()
	if err != nil {
		s.Push(n)
		return errRevLenOneStack
	}
	s.Push(n)
	s.Push(m)
	return nil
}

func (s *stack) Clear() {
	*s = make(stack, 0)
}
