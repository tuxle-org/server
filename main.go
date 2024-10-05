package main

import (
	"log/slog"
	"os"

	"github.com/bbfh-dev/parsex/parsex"
	"github.com/bbfh-dev/plog/plog"
)

var Version string

func Program(in parsex.Input, args ...string) error {
	if in.Has("version") {
		println("Tuxle server " + Version)
		return nil
	}

	if in.Has("debug") {
		plog.SetupDefault(slog.LevelDebug)
	} else {
		plog.SetupDefault(slog.LevelInfo)
	}

	return nil
}

var CLI = parsex.New("example", Program, []parsex.Arg{
	{Name: "version", Match: "--AUTO,-v", Desc: "print version and exit"},
	{Name: "debug", Match: "--AUTO", Desc: "log verbose debug information"},
})

func main() {
	if err := CLI.FromArgs().Run(); err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}
}
