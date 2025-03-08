package main

import (
	"flag"
	"tagliatelle/pkg/settings"
	"tagliatelle/pkg/tagliatelle"

	log "github.com/sirupsen/logrus"
)

var (
	dryRun      bool
	pattern     string
	repo        string
	tag         string
	project     string
	environment string
	filePath    string
)

func main() {
	if err := settings.Load(); err != nil {
		log.WithError(err).
			Error("failed to load settings")
	}

	flag.StringVar(&repo, "repo", "", "name of git repository")
	flag.StringVar(&filePath, "file", "", "file path to update")
	flag.StringVar(&tag, "tag", "", "new tag to use for update")
	flag.StringVar(&pattern, "pattern", "", "regex pattern to find and replace tag")
	flag.StringVar(&project, "project", "", "project name")
	flag.StringVar(&environment, "environment", "", "environment name")
	flag.BoolVar(&dryRun, "dry-run", false, "enable dry run")
	flag.Parse()

	switch {
	case repo == "":
		invalid("repo")
	case tag == "":
		invalid("tag")
	case pattern == "":
		invalid("pattern")
	case project == "":
		invalid("project")
	case environment == "":
		invalid("environment")
	case filePath == "":
		invalid("filename")
	}

	opts := tagliatelle.Options{
		DryRun:      dryRun,
		GitRepo:     repo,
		Pattern:     pattern,
		Tag:         tag,
		Project:     project,
		Environment: environment,
		FilePath:    filePath,
	}

	if err := tagliatelle.Entrypoint(opts); err != nil {
		log.WithError(err).
			Fatal("tagliatelle failed to run")
	}
}

func invalid(str string) {
	log.Fatal("invalid parameter: " + str)
}
