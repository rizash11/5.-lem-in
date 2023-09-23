package main

import "fmt"

func (app *application) findOptimalPaths() {
	var path []string
	app.pathSearching(path, app.startRoom)
	if len(app.paths) == 0 {
		app.errorLog.Fatalln("no paths from start to end")
	}

	app.sortPaths()

	for _, path := range app.paths {
		if app.optimal(path) {
			app.optimalPaths = append(app.optimalPaths, path)
		}
	}
}

func (app *application) pathSearching(path []string, roomName string) {
	path = append(path, roomName)
	fmt.Printf("%p: %v\n", &path, path)

	for _, link := range app.rooms[roomName].links {
		if app.rooms[link] == app.rooms[app.endRoom] {
			path = append(path, link)
			fmt.Printf("%p: %v\n", &path, path)

			correctPath := make([]string, len(path))
			copy(correctPath, path)

			app.paths = append(app.paths, correctPath[1:])

			path = path[:len(path)-1]
		} else if !app.presentInPath(link, path) {
			app.pathSearching(path, link)
		}
	}

}

func (app *application) presentInPath(link string, path []string) bool {
	for _, roomName := range path {
		if app.rooms[link] == app.rooms[roomName] {
			return true
		}
	}

	return false
}

func (app *application) optimal(path []string) bool {
	for _, link := range path {
		for _, optimalPath := range app.optimalPaths {
			if app.presentInPath(link, optimalPath) {
				return false
			}
		}
	}

	return true
}

func (app *application) sortPaths() {
	for i := range app.paths {
		for j := i + 1; j < len(app.paths); j++ {
			if len(app.paths[i]) > len(app.paths[j]) {
				app.paths[i], app.paths[j] = app.paths[j], app.paths[i]
			}
		}
	}
}

func (app *application) moveAnts() {
	antsOnPaths := make([]int, len(app.optimalPaths))

	for _, ant := range app.ants {
		for i := len(app.optimalPaths) - 1; i >= 0; i-- {
			if antsOnPaths[i]+len(app.optimalPaths[i]) <= antsOnPaths[0]+len(app.optimalPaths[0]) {
				antsOnPaths[i]++
				ant.Order = -antsOnPaths[i]
				ant.Path = app.optimalPaths[i]
				break
			}
		}
	}

	for i, j := 1, len(antsOnPaths)-1; i < antsOnPaths[j]+len(app.optimalPaths[j]); i++ {
		for k, ant := range app.ants {
			ant.move(k + 1)
		}
		fmt.Println()
	}
}

func (ant *Ant) move(antID int) {
	ant.Order++
	if ant.Order >= 0 && ant.Order < len(ant.Path) {
		fmt.Printf("L%d-%s ", antID, ant.Path[ant.Order])
	}
}
