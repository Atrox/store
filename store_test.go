package store

import (
	"testing"
)

type MyCustomConfig struct {
	Name        string `yaml:"name"`
	OtherConfig string `yaml:"otherConfig"`
}

const name = "application-name"

func TestGet(t *testing.T) {
	storage, err := New(name)
	if err != nil {
		t.Fatal(err)
	}

	customConfig := &MyCustomConfig{
		Name: name,
	}

	err = storage.Save(customConfig)
	if err != nil {
		t.Error(err)
	}

	newCustomConfig := new(MyCustomConfig)
	err = storage.Get(newCustomConfig)
	if err != nil {
		t.Error(err)
	}

	if newCustomConfig.Name != name {
		t.Errorf("Name %s is not in retrieved config name %s", name, newCustomConfig.Name)
	}

	// cleanup
	storage.Remove(newCustomConfig)
}
