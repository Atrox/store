# Store

[![Build Status](https://img.shields.io/travis/Atrox/store.svg?style=flat-square)](https://travis-ci.org/Atrox/store)
[![Coverage Status](https://img.shields.io/coveralls/Atrox/store.svg?style=flat-square)](https://coveralls.io/r/Atrox/store)
[![Go Report Card](https://goreportcard.com/badge/github.com/Atrox/store?style=flat-square)](https://goreportcard.com/report/github.com/Atrox/store)
[![GoDoc](https://img.shields.io/badge/godoc-reference-5272B4.svg?style=flat-square)](https://godoc.org/github.com/Atrox/store)

> Simple, painless, zero-config configuration storage

## Installation

```sh
go get -u github.com/atrox/store
# or with dep
dep ensure -add github.com/atrox/store
```

## Example

```go
package main

import (
	"fmt"

	"github.com/atrox/store"
)

type Config struct {
	Name string `yaml:"name"`
	Age  int    `yaml:"age"`
}

func main() {
	storage, err := store.New("testapp")
	if err != nil {
		panic(err)
	}

	config := &Config{
		Name: "Dr. Robert Ford",
		Age:  70,
	}
	err = storage.Save(config)
	if err != nil {
		panic(err)
	}
	fmt.Printf("config is now saved at %s\n", storage.Path(config))
	// ~/.config/testapp/config.yaml

	config = &Config{}
	err = storage.Get(config)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Config retrieved:\n%+v\n", config)
	// &{Name:Dr. Robert Ford Age:70}
}
```

## Contributing

Everyone is encouraged to help improve this project. Here are a few ways you can help:

- [Report bugs](https://github.com/atrox/store/issues)
- Fix bugs and [submit pull requests](https://github.com/atrox/store/pulls)
- Write, clarify, or fix documentation
- Suggest or add new features
