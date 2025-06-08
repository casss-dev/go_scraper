package util

import (
	"os"
	"path/filepath"
	"strings"
)

// if the path contains the `~` character, expands to the user's home directory
func ExpandUserHomeDir(path string) string {
	dir := path
	if dir == "~" {
		homeDir, err := os.UserHomeDir()
		if err == nil {
			dir = homeDir
		}
	} else if len(dir) > 0 && dir[0] == '~' {
		homeDir, err := os.UserHomeDir()
		if err == nil {
			dir = filepath.Join(homeDir, dir[1:])
		}
	}
	return dir
}

// creates a json filename based on a search query
func FilenameFromQuery(query string) string {
	safeQuery := strings.ReplaceAll(query, " ", "_")
	safeQuery = strings.ReplaceAll(safeQuery, "/", "_")
	safeQuery = strings.ReplaceAll(safeQuery, "\\", "_")
	return safeQuery + ".json"
}
