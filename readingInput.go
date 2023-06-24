package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func (app *application) readInput(inputFile string) {
	f, err := os.Open(inputFile)
	if err != nil {
		app.errorLog.Fatalln(err)
	}

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		antNumber, err := strconv.Atoi(scanner.Text())
		if err != nil {
			continue
		} else if antNumber <= 0 {
			app.errorLog.Fatalln("invalid number of ants")
		}

		app.ants = make([]*Ant, antNumber)
		for i := range app.ants {
			app.ants[i] = &Ant{}
		}
		break
	}

	for scanner.Scan() {
		input := scanner.Text()

		switch {
		case input == "##start":
			scanner.Scan()
			input := scanner.Text()

			// if the next line is not a Room, skip lines until you find a room
			if !app.NewRoom(input, "start") {
				for scanner.Scan() {
					input := scanner.Text()
					if app.NewRoom(input, "start") {
						break
					}
				}
			}
		case input == "##end":
			scanner.Scan()
			input := scanner.Text()

			// if the next line is not a Room, skip lines until you find a room
			if !app.NewRoom(input, "end") {
				for scanner.Scan() {
					input := scanner.Text()
					if app.NewRoom(input, "end") {
						break
					}
				}
			}
		default:
			if app.NewRoom(input, "regular") {
				continue
			}
			app.newLink(input)
		}
	}
}

func (app *application) NewRoom(input, roomType string) bool {
	input2 := strings.Split(input, " ")

	if len(input2) != 3 {
		return false
	}

	var coordinates [2]int
	var err error

	coordinates[0], err = strconv.Atoi(input2[1])
	if err != nil || coordinates[0] < 0 {
		app.errorLog.Println("invalid input:", input)
		return false
	}
	coordinates[1], err = strconv.Atoi(input2[2])
	if err != nil || coordinates[1] < 0 {
		app.errorLog.Println("invalid input:", input)
		return false
	}

	for _, room := range app.rooms {
		if room.coordinates == coordinates {
			app.rooms[input2[0]] = room
			return true
		}
	}

	app.rooms[input2[0]] = &Room{coordinates: coordinates, roomType: roomType}

	return true
}

func (app *application) newLink(input string) {
	input2 := strings.Split(input, "-")

	if len(input2) != 2 {
		return
	}

	room1, ok := app.rooms[input2[0]]
	if !ok {
		app.errorLog.Println("ignoring a link, link to non-existing room:", input)
		return
	}
	room2, ok := app.rooms[input2[1]]
	if !ok {
		app.errorLog.Println("ignoring a link, link to non-existing room:", input)
		return
	}
	if room1.coordinates == room2.coordinates {
		app.errorLog.Println("ignoring a link, room links to itself:", input)
		return
	}
	for _, link := range room1.links {
		if room2.coordinates == app.rooms[link].coordinates {
			app.errorLog.Println("ignoring a duplicate link:", input)
			return
		}
	}

	room1.links = append(room1.links, input2[1])
	room2.links = append(room2.links, input2[0])
}
