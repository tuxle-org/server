package tuxle

import (
	"github.com/bbfh-dev/go-tools/tools/terr"
	"github.com/tuxle-org/lib/tuxle/entities"
	"gorm.io/gorm"
)

func ServerInfo(db *gorm.DB) (entities.Server, error) {
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
		return server, terr.Prefix("SELECT Server", err)
	}

	return server, nil
}
