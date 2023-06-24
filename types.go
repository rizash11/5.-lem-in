package main

import "log"

type application struct {
	rooms              map[string]*Room
	errorLog           *log.Logger
	infoLog            *log.Logger
	startRoom, endRoom string
	paths              [][]string
	optimalPaths       [][]string
	ants               []*Ant
}

type Room struct {
	links       []string
	coordinates [2]int
	roomType    string
}

type Ant struct {
	Order int
	Path  []string
}
