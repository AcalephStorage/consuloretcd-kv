package command

import (
	"errors"
	"gopkg.in/ashcrow/consuloretcd.v1"
	"github.com/codegangsta/cli"
	"net/http"
)

// Makes a new client based on cli context
func makeClient(c *cli.Context) (consuloretcd.KeyValueStore, error) {
	//var client consuloretcd.KeyValueStore
	clientone := http.Client{}
	Address := c.GlobalString("address")
	Port := c.GlobalString("port")
	if c.Bool("consul") == true {
		return consuloretcd.NewClient(
			"consul",
			consuloretcd.Config{
				Endpoint: Address,
				Client:   clientone,
				Port:     Port,
			})
	}
	if c.Bool("etcd") == true {
		return consuloretcd.NewClient(
			"etcd",
			consuloretcd.Config{
				Endpoint: Address,
				Client:   clientone,
				Port:     Port,
			})
	}
	return nil, errors.New("Unknown client requested.")
}
