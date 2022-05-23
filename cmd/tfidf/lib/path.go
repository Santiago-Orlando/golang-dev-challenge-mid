package lib

import "strings"

func Path(file string) string {

	path := strings.Split(file, "/")
	dir := strings.Join(path[:len(path)-1], "/")

	if len(dir) == 0 {
		return ""
	}

	if dir[0:1] == "/" {
		dir = dir[1:]
	}

	return dir
}
