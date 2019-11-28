package main

import "fmt"

type animal struct {
	active   bool
	house    string
	species  string
	name     string
	life     int
	strength int
	speed    int
	lethal   int
	flight   bool
	swim     bool
}

func GetLion() animal {
	l := animal{}
	l.active = true
	l.house = "Africa"
	l.species = "cat"
	l.name = "Lion"
	l.life = 10
	l.strength = 10
	l.speed = 8
	l.lethal = 10
	l.flight = false
	l.swim = false
	return l
}

func (a animal) PrintAnimal() {
	fmt.Printf("Name: %v \n Information: %+v /n", a.name, a)

}
