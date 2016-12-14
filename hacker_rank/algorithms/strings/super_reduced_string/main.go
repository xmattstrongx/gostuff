package main

import (
	"fmt"
	"strings"
)

type superReducedStringer struct {
	original string
	reduced  string
}

func main() {
	srstring := superReducedStringer{}

	srstring.readInData()
	srstring.performSuperReduction()
	srstring.printSuperReducedString()
}

func (s *superReducedStringer) readInData() {
	fmt.Scanf("%s", &s.original)
}

func (s *superReducedStringer) performSuperReduction() {
	start := 0
	s.reduced = s.original
	for i := start; i < len(s.reduced)-1; {
		if s.reduced[i] == s.reduced[i+1] {
			s.reduced = strings.Replace(s.reduced, (string(s.reduced[i]) + string(s.reduced[i+1])), "", 1)
			i = 0
		} else {
			i++
		}
	}
}

func (s *superReducedStringer) printSuperReducedString() {
	if s.reduced == "" {
		s.reduced = "Empty String"
	}
	fmt.Printf("%s", s.reduced)
}
