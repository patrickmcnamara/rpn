package main

import "errors"

var (
	errDivByZero      = errors.New("division by zero")
	errPopEmptyStack  = errors.New("popping empty stack")
	errPeekEmptyStack = errors.New("peeking empty stack")
	errDupEmptyStack  = errors.New("duplicating empty stack")
	errRevEmptyStack  = errors.New("reversing empty stack")
	errRevLenOneStack = errors.New("reversing length one stack")
	errInvalidInput   = errors.New("invalid input")
)
