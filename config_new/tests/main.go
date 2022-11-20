package main

import (
	"fmt"

	"github.com/alexadastra/ramme/config_new"
)

func main() {
	conf, start, stop, err := config_new.NewConfig("./test/config.yaml")
	if err != nil {
		panic(err)
	}

	go func() { _ = start() }()
	defer func() { _ = stop() }()

	fmt.Println(conf.Get("some_int"))
	fmt.Println(1 + config_new.ToInt(conf.Get("some_bool")))
}
