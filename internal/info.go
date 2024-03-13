package internal

import "fmt"

var (
	version   string
	buildTime string
)

func BuildInfo() string {
	return fmt.Sprintf("golagen version %s, build %s", version, buildTime)
}

func GetVersion() string {
	return version
}
