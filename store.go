// Package store provides a simple, painless configuration storage
package store

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"reflect"
	"strings"

	"github.com/atrox/homedir"
	yaml "gopkg.in/yaml.v2"
)

// Store encapsulates all information required to read and save files
type Store struct {
	Base string
}

// New initializes store with sane defaults
// Base will be set to ~/.config/<name>/
func New(name string) (*Store, error) {
	home, err := homedir.Dir()
	if err != nil {
		return nil, err
	}
	base := filepath.Join(home, ".config", name)

	// create directory
	err = os.MkdirAll(base, os.ModePerm)
	if err != nil {
		return nil, err
	}

	return &Store{Base: base}, nil
}

// Get retrieves the passed in interface from the filesystem
func (s *Store) Get(i interface{}) error {
	location := s.Path(i)

	file, err := ioutil.ReadFile(location)
	if err != nil {
		return err
	}

	err = yaml.Unmarshal(file, i)
	if err != nil {
		return err
	}

	return nil
}

// Save persists the passed in interface to the filesystem
func (s *Store) Save(i interface{}) error {
	location := s.Path(i)

	b, err := yaml.Marshal(i)
	if err != nil {
		return err
	}

	return ioutil.WriteFile(location, b, 0644)
}

// Remove removes the file from the passed in interface on the filesystem
func (s *Store) Remove(i interface{}) error {
	location := s.Path(i)
	return os.Remove(location)
}

// Path returns the file path for the specified interface
func (s *Store) Path(i interface{}) string {
	name := strings.ToLower(reflect.TypeOf(i).Elem().Name())
	return filepath.Join(s.Base, name+".yaml")
}
