// Package common provides common functionality for the application.
package common

import (
	"fmt"
	"runtime"
)

// Version information set by link flags during build. We fall back to these sane
// default values when we build outside the Makefile context (e.g. go run, go build, or go test).
var (
	version   = "99.99.99"             // value from VERSION file
	buildDate = "1970-01-01T00:00:00Z" // output from `date -u +'%Y-%m-%dT%H:%M:%SZ'`
	gitCommit = ""                     // output from `git rev-parse HEAD`
)

// Version represents the version information of the application.
type Version struct {
	Version   string
	BuildDate string
	GitCommit string
	GoVersion string
	Compiler  string
	Platform  string
}

// String returns the version as a string
func (v Version) String() string {
	return v.Version
}

// GetVersion returns the version information
func GetVersion() Version {
	var versionStr string

	// otherwise formulate a version string based on as much metadata
	// information we have available.
	versionStr = "v" + version
	if len(gitCommit) >= 7 {
		versionStr += "+unknown"
	}

	return Version{
		Version:   versionStr,
		BuildDate: buildDate,
		GitCommit: gitCommit,
		GoVersion: runtime.Version(),
		Compiler:  runtime.Compiler,
		Platform:  fmt.Sprintf("%s/%s", runtime.GOOS, runtime.GOARCH),
	}
}
