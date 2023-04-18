package pathseparator

import "path"

func SplitPath(input string) (dir, file string) {
	return path.Split(input)
}
