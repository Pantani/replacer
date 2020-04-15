package storage

import (
	"github.com/Pantani/replacer/internal/config"
	"github.com/Pantani/replacer/pkg/replacer"
	"github.com/Pantani/replacer/pkg/storage/mongodb"
	"net/url"
	"strings"
)

var (
	Database = "replacer"
)

type Storage struct {
	*mongodb.MongoDb
}

func New() (*Storage, error) {
	u, err := url.Parse(config.Configuration.Mongo.Uri)
	if err == nil {
		dbName := strings.Replace(u.Path, "/", "", -1)
		if len(dbName) > 0 {
			Database = dbName
		}
	}
	mongo, err := mongodb.NewMongoDbClient(config.Configuration.Mongo.Uri)
	if err != nil {
		return nil, err
	}
	return &Storage{mongo}, nil
}

type Name interface {
	InsertOrUpdateName(name replacer.Name) error
	GetName(name string) (result *replacer.Name, err error)
	GetAllNames() (result replacer.Names, err error)
	DeleteName(name string) error
	DeleteAllNames() error
}
