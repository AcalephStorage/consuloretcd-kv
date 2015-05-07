package command

import (
	"errors"
	"github.com/ashcrow/consuloretcd"
	"github.com/codegangsta/cli"
	"net/http"
)

// Makes a new client based on cli context
func makeClient(c *cli.Context) (consuloretcd.KeyValueStore, error) {
	//var client consuloretcd.KeyValueStore
	clientone := http.Client{}

	if c.Bool("consul") == true {
		return consuloretcd.NewClient(
			"consul",
			consuloretcd.Config{
				Endpoint: "http://127.0.0.1",
				Client:   clientone,
				Port:     8500,
			})
	}
	if c.Bool("etcd") == true {
		return consuloretcd.NewClient(
			"etcd",
			consuloretcd.Config{
				Endpoint: "http://127.0.0.1",
				Client:   clientone,
				Port:     8500,
			})
	}
	return nil, errors.New("Unknown client requested.")
}
