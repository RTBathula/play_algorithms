package main

import (
	"strconv"
	"github.com/kr/pretty"
)

type Queue struct {
	Counter int
	Storage map[string]interface{}
	LowestCount int
}

func (s *Queue) Enqueue(value interface{}) {
	position := strconv.Itoa(s.Counter)

	if(len(s.Storage) == 0){
		s.Storage = make(map[string]interface{})
	}

	s.Storage[position] = value
	s.Counter++
}

func (s *Queue) Dequeue() interface{} {
	if((s.Counter - s.LowestCount) == 0 ){
		return nil
	}

	position := strconv.Itoa(s.LowestCount)	
	result := s.Storage[position]
	s.LowestCount++
	delete(s.Storage, position)
	return result
}

func main(){
	var q Queue	
	q.Enqueue("hello")
	pretty.Println(q)
	q.Dequeue()
	pretty.Println(q)
	q.Enqueue("world")
	pretty.Println(q)
}