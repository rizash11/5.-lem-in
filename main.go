package main

import (
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
	app.startRoom, app.endRoom = app.checkForInputErrors()

	app.findOptimalPaths()
	app.moveAnts()
}
