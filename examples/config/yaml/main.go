package main

import (
	"fmt"

	"github.com/zjllib/go-micro/config"
	"github.com/zjllib/go-micro/config/reader"
	"github.com/zjllib/go-micro/config/reader/json"
	"github.com/zjllib/go-micro/config/source/file"
	"github.com/zjllib/go-micro/plugins/config/encoder/yaml"
	"github.com/zjllib/go-micro/plugins/config/source/etcd"
)

func main() {
	// new yaml encoder
	enc := yaml.NewEncoder()

	// new config
	c, _ := config.NewConfig(
		config.WithReader(
			json.NewReader( // json reader for internal config merge
				reader.WithEncoder(enc),
			),
		),
	)

	// load the config from a file source
	if err := c.Load(file.NewSource(
		file.WithPath("./config.yaml"),
		etcd.WithAddress(),
	)); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("data", c.Map())

	// define our own host type
	type Host struct {
		Address string `json:"address"`
		Port    int    `json:"port"`
	}

	var host Host

	// read a database host
	if err := c.Get("hosts", "database").Scan(&host); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(host.Address, host.Port)
}
