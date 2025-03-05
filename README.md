# go-option

Functional options.

## Dependencies
None

## Installation
```shell
go get github.com/go-stdlib/go-option
```

## Usage

The package offers functional options for any type.

```go
package main

import (
    "github.com/go-stdlib/go-option"
)

type Config struct {
    Adapter string
    RPSLimit int
}

func WithAdapter(adapter string) option.Option[*Config] {
    return func(c *Config) error {
        c.Adapter = adapter
        return nil
    }
}

func WithRPSLimit(limit int) option.Option[*Config] {
    return func(c *Config) error {
        c.RPSLimit = limit
        return nil
    }
}

func main() {
    // Use `New` to create a zero value of the reference type with functional options.
    c1, err := option.New[*Config]()
    // &Config{Adapter: "", RPSLimit: 0}
    c2, err := option.New(WithRPSLimit(1024))
    // &Person{Adapter: "", RPSLimit: 1024}

    // Use `Apply` to mutate an existing reference type with functional options.
    c3, err := option.Apply(&Config{})
    // &Person{Name: "", Role: ""}
    c4, err := option.Apply(&Config{Adapter: "sqlite"}, WithRPSLimit(50))
    // &Person{Name: "sqlite", Role: 50}
}
```


## License

[Apache 2.0](../LICENSE)

