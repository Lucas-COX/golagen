package internal

import "fmt"

var (
	version   string
	buildTime string
)

func BuildInfo() string {
	return fmt.Sprintf("golagen %s, build %s", version, buildTime)
}

func GetVersion() string {
	return version
}
