package main

func (app *application) pathSearching(path []string, roomName string) {
	path = append(path, roomName)

loop:
	for _, link := range app.rooms[roomName].links {
		if !app.alreadyVisited(link, path) {
			switch {
			case app.rooms[link] == app.rooms[app.endRoom]:
				path = append(path, link)

				correctPath := make([]string, len(path))
				copy(correctPath, path)

				app.paths = append(app.paths, correctPath)
				break loop
			default:
				app.pathSearching(path, link)
			}
		}
	}

}

func (app *application) alreadyVisited(link string, path []string) bool {
	for _, roomName := range path {
		if app.rooms[link] == app.rooms[roomName] {
			return true
		}
	}

	return false
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
