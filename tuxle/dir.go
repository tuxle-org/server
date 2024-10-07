package tuxle

import (
	"os"
	"path/filepath"

	"github.com/emersion/go-appdir"
)

var DataDir = filepath.Join(appdir.New("tuxle").UserData(), "server")

func MakeDirs() error {
	return os.MkdirAll(DataDir, os.ModePerm)
}

var DbFile = filepath.Join(DataDir, "server.db")
