package storage

import (
	"github.com/Pantani/errors"
	"github.com/Pantani/replacer/pkg/replacer"
	"go.mongodb.org/mongo-driver/bson"
)

const (
	NameCollection = "names"
)

func (s *Storage) InsertName(name replacer.Name) (err error) {
	_, err = s.InsertMany(Database, NameCollection, []interface{}{name})
	return
}

func (s *Storage) UpdateName(name *replacer.Name) error {
	_, err := s.Update(Database, NameCollection, name, bson.D{{"name", name.Name}})
	if err != nil {
		return errors.E(err, "cannot update name", errors.Params{"name": name.Name})
	}
	return nil
}

func (s *Storage) GetName(name string) (result *replacer.Name, err error) {
	query := bson.M{"name": name}
	err = s.GetValue(Database, NameCollection, query, &result)
	return
}

func (s *Storage) GetAllNames() (result replacer.Names, err error) {
	err = s.GetValues(Database, NameCollection, nil, &result)
	return
}

func (s *Storage) InsertOrUpdateName(name replacer.Name) error {
	r, err := s.GetName(name.Name)
	if err == nil && r != nil {
		return s.UpdateName(&name)
	}
	return s.InsertName(name)
}

func (s *Storage) DeleteName(name string) error {
	query := bson.M{"name": name}
	_, err := s.DeleteOne(Database, NameCollection, query)
	return err
}

func (s *Storage) DeleteAllNames() error {
	_, err := s.DeleteMany(Database, NameCollection, bson.D{})
	return err
}
