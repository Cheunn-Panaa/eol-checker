package main

import (
	"github.com/cheunn-panaa/eol-checker/cmd"
	"github.com/cheunn-panaa/eol-checker/internal/utils"
)

var buildVersion = "v0.0.1"

func main() {
	utils.SetVersion(buildVersion)
	cmd.Execute()
}
