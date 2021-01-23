package main

import (
	"fmt"
	"github.com/l0k18/spore/cmd/spore"
	"github.com/l0k18/spore/pkg/log"
	"github.com/l0k18/spore/version"
	"os"
)

func main() {
	version.URL = URL
	version.GitRef = GitRef
	version.GitCommit = GitCommit
	version.BuildTime = BuildTime
	version.Tag = Tag
	version.Get = GetVersion
	log.SetLogLevel("debug")
	Debug(version.Get())
	s := spore.New()
	result := s.Main()
	if result != 0 {
		os.Exit(result)
	}
}

var (
	URL       string
	GitRef    string
	GitCommit string
	BuildTime string
	Tag       string
)

func GetVersion() string {
	var err error
	var wd string
	if wd, err = os.Getwd(); Check(err) {
	}
	return fmt.Sprintf(
		`application information:
	working directory: %s
	command: '%s'
	repo: %s
	branch: %s
	commit: %s
	built: %s
	tag: %s
`,
		wd, os.Args[0], URL, GitRef, GitCommit, BuildTime, Tag,
	)
}
