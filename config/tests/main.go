// Package main contains test
package main

import (
	"fmt"

	"github.com/alexadastra/ramme/config"
)

func main() {
	conf, start, stop, err := config.NewConfig("./test/config.yaml")
	if err != nil {
		panic(err)
	}

	go func() { _ = start() }()
	defer func() { _ = stop() }()

	fmt.Println(conf.Get("some_int"))
	fmt.Println(1 + config.ToInt(conf.Get("some_bool")))
}
