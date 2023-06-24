package main

func (app *application) checkForInputErrors() (startRoom, endRoom string) {
	// no start/end rooms, >1 start/end room
	startCounter := 0
	endCounter := 0
	for name, room := range app.rooms {
		if room.roomType == "start" && app.rooms[startRoom] != room {
			startRoom = name
			startCounter++
		} else if room.roomType == "end" && app.rooms[endRoom] != room {
			endRoom = name
			endCounter++
		}
	}

	if startCounter != 1 || endCounter != 1 {
		app.errorLog.Fatalln("invalid number of start/end rooms")
	}

	if app.rooms[startRoom] == app.rooms[endRoom] {
		app.errorLog.Fatalln("start and end room have the same coordinates")

	}

	return startRoom, endRoom
}
