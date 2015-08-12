package utils

import (
	"appengine"
	"net/http"
	"os"
	"strings"
)

var rootPackageName = "simpletimeline"

func GetTemplatePath() (path string) {
	path, err := os.Getwd()
	if err != nil {
		return "templates/"
	}

	if appengine.IsDevAppServer() {
		return strings.Repeat("../", GetRootPackageOffset(path)) + "app/templates/"
	} else {
		// The production environment has a different directory structure
		return "templates/"
	}
}

func GetRootPackageOffset(path string) (offset int) {
	rootOffset := 0
	dirs := strings.Split(path, "/")
	for i, dir := range dirs {
		rootOffset = i + 1
		if dir == rootPackageName {
			break
		}
	}

	return len(dirs) - rootOffset
}

func CheckHandlerError(c appengine.Context, err error, w http.ResponseWriter, message string) (isError bool) {
	if err != nil {
		c.Errorf(message)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return true
	}
	return false
}
