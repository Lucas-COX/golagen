package internal

import "path/filepath"

var cachePath string = ".golagen"
var templateRepoUrl string = "git@github.com:Lucas-COX/golagen-templates.git"

func GetCachePath() string {
	absPath, _ := filepath.Abs(cachePath)
	return absPath
}

func GetTemplateRepoUrl() string {
	return templateRepoUrl
}
