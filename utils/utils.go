package utils

import (
  "os"
  "strings"
)

var rootPackageName = "simpletimeline"

func GetTemplatePath() (path string){
  path, err := os.Getwd()
  if err != nil {
    return "templates/"  
  }

  return strings.Repeat("../", GetRootPackageOffset(path)) + "app/templates/"
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