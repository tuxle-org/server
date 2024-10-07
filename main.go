package main

import (
	"log/slog"
	"os"
	"strconv"

	"github.com/bbfh-dev/go-tools/tools/terr"
	"github.com/bbfh-dev/parsex/parsex"
	"github.com/bbfh-dev/plog/plog"
	"github.com/tuxle-org/lib/tuxle/entities"
	"github.com/tuxle-org/server/tuxle"
	"github.com/tuxle-org/server/web"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
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
		return terr.Prefix("Converting port argument into number", err)
	}

	db, err := openDB()
	if err != nil {
		return err
	}

	if err = setupDB(db); err != nil {
		return err
	}

	return terr.Prefix("Running HTTP server", web.ServeHTTP(db, port))
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

func openDB() (*gorm.DB, error) {
	if err := tuxle.MakeDirs(); err != nil {
		return nil, terr.Prefix("Making tuxle dirs", err)
	}

	db, err := gorm.Open(sqlite.Open(tuxle.DbFile), new(gorm.Config))
	if err != nil {
		return nil, terr.Prefix("Opening database", err)
	}

	err = db.AutoMigrate(
		new(entities.PermissionMask),
		new(entities.Tag),
		new(entities.Role),
		new(entities.User),
		new(entities.Channel),
		new(entities.Server),
		new(entities.Directory),
		new(entities.MessageVote),
		new(entities.TextMessage),
	)
	if err != nil {
		return nil, terr.Prefix("Running migrations", err)
	}

	slog.Info("Opened database", "file", tuxle.DbFile)
	return db, nil
}

func setupDB(db *gorm.DB) error {
	slog.Debug("Ensuring default Server configuration exists...")

	var server = entities.Server{
		Name:        "Unnamed Server",
		Description: "Welcome!",
		Rules:       "",
		IconURI:     nil,
		BannerURI:   nil,
		OwnerId:     0,
		Region:      "US:en",
	}
	if err := db.FirstOrCreate(&server).Error; err != nil {
		return terr.Prefix("SELECT Server", err)
	}

	return nil
}
