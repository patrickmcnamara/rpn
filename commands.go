package main

import (
	"encoding/gob"
	"fmt"
	"math/big"
	"os"
	"strings"
)

var commands = map[string]func(s *stack) error{
	"f": func(s *stack) error {
		n, err := s.Peek()
		if err != nil {
			return err
		}
		fmt.Println(n)
		return nil
	},
	"d": func(s *stack) error {
		n, err := s.Peek()
		if err != nil {
			return err
		}
		fmt.Println(new(big.Float).SetPrec(256).SetRat(n).Text('g', 64))
		return nil
	},
	"sf": func(s *stack) error {
		var ss strings.Builder
		for i, n := range *s {
			ss.WriteString(n.RatString())
			if i != len(*s) {
				ss.WriteByte(' ')
			}
		}
		fmt.Println(ss.String())
		return nil
	},
	"sd": func(s *stack) error {
		var ss strings.Builder
		for i, n := range *s {
			ss.WriteString(new(big.Float).SetPrec(256).SetRat(n).Text('g', 4))
			if i != len(*s) {
				ss.WriteByte(' ')
			}
		}
		fmt.Println(ss.String())
		return nil
	},
	"P": func(s *stack) error {
		_, err := s.Pop()
		if err != nil {
			return err
		}
		return nil
	},
	"D": func(s *stack) error {
		err := s.Duplicate()
		if err != nil {
			return err
		}
		return nil
	},
	"R": func(s *stack) error {
		err := s.Reverse()
		if err != nil {
			return err
		}
		return nil
	},
	"C": func(s *stack) error {
		s.Clear()
		return nil
	},
	"S": func(s *stack) error {
		cacheDir, err := os.UserCacheDir()
		if err != nil {
			return err
		}
		fd, err := os.Create(cacheDir + "/rpn")
		if err != nil {
			return err
		}
		defer fd.Close()
		enc := gob.NewEncoder(fd)
		enc.Encode(s)
		return nil
	},
	"L": func(s *stack) error {
		cacheDir, err := os.UserCacheDir()
		if err != nil {
			return err
		}
		fd, err := os.Open(cacheDir + "/rpn")
		if err != nil {
			return err
		}
		defer fd.Close()
		dec := gob.NewDecoder(fd)
		dec.Decode(&s)
		return nil
	},
}
