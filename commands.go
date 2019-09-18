package main

import (
	"encoding/gob"
	"fmt"
	"os"
)

var commands = map[string]func(s *stack) error{
	"p": func(s *stack) error {
		n, err := s.Peek()
		if err != nil {
			return err
		}
		fmt.Println(n)
		return nil
	},
	"s": func(s *stack) error {
		fmt.Println(*s)
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
		defer fd.Close()
		if err != nil {
			return err
		}
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
		defer fd.Close()
		if err != nil {
			return err
		}
		dec := gob.NewDecoder(fd)
		dec.Decode(&s)
		return nil
	},
}
