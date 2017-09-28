package main

import (
	"strconv"
	"github.com/kr/pretty"
)

type Stack struct {
	Counter int
	Storage map[string]interface{}
}

func (s *Stack) Push(value interface{}) {
	position := strconv.Itoa(s.Counter)

	if(len(s.Storage) == 0){
		s.Storage = make(map[string]interface{})
	}

	s.Storage[position] = value
	s.Counter++
}

func (s *Stack) Pop() interface{} {
	if(s.Counter == 0){
		return nil
	}

	s.Counter--
	position := strconv.Itoa(s.Counter)
	result := s.Storage[position]
	delete(s.Storage, position)
	return result
}

func (s *Stack) Size() int {
	return s.Counter
}

func main(){	
	var stack Stack	
	stack.Push("hello")
	pretty.Println(stack)
	stack.Pop()
	pretty.Println(stack)
	stack.Push("world")
	pretty.Println(stack)
}