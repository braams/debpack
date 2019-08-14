package main

import (
	"fmt"
	"github.com/braams/debpack"
)

var Version = "0.0.0"

func main() {
	fmt.Println(Version)
	p := debpack.NewPackage("debpack", "1.2.4")
	p.Build("cmd/debpack.go")
	p.MaintainerName = "maintainer"
	p.MaintainerEmail = "maintainer@email.com"
	p.SetDefaultFilenames()
	p.SetDefaultFiles()
	p.AddControls()
	p.AddFiles()
	p.MarkConfig()
	p.Pack()

}
