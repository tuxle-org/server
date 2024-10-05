package main

import (
	"log/slog"
	"os"
	"strconv"

	"github.com/bbfh-dev/go-tools/tools/terr"
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

	port, err := strconv.Atoi(in.Default("port", "8080"))
	if err != nil {
		return err
	}
}

var CLI = parsex.New("example", Program, []parsex.Arg{
	{Name: "version", Match: "--AUTO,-v", Desc: "print version and exit"},
	{Name: "debug", Match: "--AUTO", Desc: "log verbose debug information"},
	{Name: "port", Match: "--AUTO,-p", Desc: "port to be used by HTTP server (Default: 8080)"},
})

func main() {
	if err := CLI.FromArgs().Run(); err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}
}
