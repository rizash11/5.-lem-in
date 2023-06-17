package main

// import (
// 	"fmt"
// 	"log"
// 	"os"
// 	"testing"
// )

// func TestNewRoom(t *testing.T) {
// 	appTest := application{
// 		rooms:    map[string]*Room{},
// 		errorLog: log.New(os.Stderr, "ERROR:\t", log.Ldate|log.Ltime|log.Lshortfile),
// 		infoLog:  log.New(os.Stdout, "INFO:\t", log.Ldate|log.Ltime),
// 	}

// 	testRooms := []string{
// 		"0 1 2",
// 		"1 5 4",
// 		"2 5 s",
// 		"3 -4 5",
// 	}

// 	for _, testRoom := range testRooms {
// 		appTest.NewRoom(testRoom, "test")
// 	}

// 	var correctRooms map[string]*Room = map[string]*Room{}
// 	correctRooms["0"] = &Room{
// 		[]string{},
// 		[2]int{1, 2},
// 		"test",
// 	}
// 	correctRooms["1"] = &Room{
// 		[]string{},
// 		[2]int{5, 4},
// 		"test",
// 	}

// 	for name, room := range appTest.rooms {
// 		if correctRooms[name].coordinates == room.coordinates {
// 			fmt.Println("room allocated correctly,", name+":", room.coordinates)
// 		} else {
// 			t.Errorf("NewRoom doesn't allocate coordinates correctly, %s should be %v, but it is %v ", name, correctRooms[name].coordinates, room.coordinates)
// 		}
// 	}

// }
