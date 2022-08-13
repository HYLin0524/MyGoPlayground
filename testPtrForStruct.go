package main

import (
	"fmt"
)

type Store struct {
	No   int
	Name string
}

func addByVal(s Store) { // receive by value ( a copy )
	fmt.Println("==========pass by value==========")
	fmt.Printf("addr = %p, val = %v\n", &s, s.No)
	s.No += 1
}

func addByPtr(s *Store) { // receive a pointer ( an address reference )
	// `*s` -> means get the value from address `s`
	fmt.Println("==========pass by reference==========")
	fmt.Printf("addr = %p, val = %v\n", s, (*s).No)
	s.No += 10
}

func main() {
	s := &Store{
		No:   100,
		Name: "HALO SuperMarket",
	}
	fmt.Printf("initial address = %p, value = %v \n", s, s)
	addByVal(*s)
	addByPtr(s)
	fmt.Println("final result:", s.No)
}
