package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	app := application{
		rooms:    map[string]*Room{},
		errorLog: log.New(os.Stderr, "ERROR:\t", log.Ldate|log.Ltime|log.Lshortfile),
		infoLog:  log.New(os.Stdout, "INFO:\t", log.Ldate|log.Ltime),
	}

	app.readInput("input.txt")
	startRoom, endRoom := app.checkForInputErrors()

	for name, room := range app.rooms {
		fmt.Printf("%s: %v, %p\n", name, *room, room)
	}
	fmt.Println(startRoom, endRoom)
}
