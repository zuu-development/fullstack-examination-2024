// Package main is the entry point of the application.
package main

import (
	"github.com/zuu-development/fullstack-examination-2024/cmd"
	_ "github.com/zuu-development/fullstack-examination-2024/docs"
)

// @title			fullstack-examination-2024 API
// @version		0.0.1
// @description	This is a server for fullstack-examination-2024.
// @license.name	Apache 2.0
// @host			localhost:8080
// @BasePath		/api/v1
// @schemes		http
func main() {
	cmd.Execute()
}
