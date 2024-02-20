package stack

import "fmt"

type Stack []int

func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

func (s *Stack) NUM(item int) {
	*s = append(*s, item)
	fmt.Println("input: ", item)
	fmt.Println(*s)
}

func (s *Stack) POP() int {
	if s.IsEmpty() {
		panic("-")
	} else {
		index := len(*s) - 1
		element := (*s)[index]
		*s = (*s)[:index]
		fmt.Println(*s)
		return element
	}
}

func (s *Stack) GETMIN() int {
	if s.IsEmpty() {
		panic("-")
	}
	min := (*s)[0]
	for i := 1; i < len(*s); i++ {
		if (*s)[i] < min {
			min = (*s)[i]
		}
	}
	return min
}

func (s *Stack) ADD() {
	if len(*s) < 2 {
		panic("-")
	}
	op2 := s.POP()
	op1 := s.POP()
	result := op1 + op2
	s.NUM(result)
}

func (s *Stack) SUB() {
	if len(*s) < 2 {
		panic("-")
	}
	op2 := s.POP()
	op1 := s.POP()
	result := op1 - op2
	s.NUM(result)
}

func (s *Stack) MULT() {
	if len(*s) < 2 {
		panic("-")
	}
	op2 := s.POP()
	op1 := s.POP()
	result := op1 * op2
	s.NUM(result)
}

func (s *Stack) DIV() {
	if len(*s) < 2 {
		panic("-")
	}
	op2 := s.POP()
	op1 := s.POP()
	result := op1 / op2
	s.NUM(result)
}
func (s *Stack) MOD() {
	if len(*s) < 2 {
		panic("-")
	}
	op2 := s.POP()
	op1 := s.POP()
	result := op1 % op2
	s.NUM(result)
}
