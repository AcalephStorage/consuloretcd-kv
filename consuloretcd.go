package main

import (
	"github.com/TelanJ/consul-kv-ceph/command"
	"github.com/codegangsta/cli"
	"os"
)

func main() {
	app := cli.NewApp()
	app.Name = "consuloretcd"
	app.Version = "0.1"
	app.Author = "Julio Telan"
	app.Email = "telanj@acale.ph"
	app.Usage = "A simple command line client for Consul/etcd k/v store."
	app.Commands = []cli.Command{
		command.NewGetCommand(),
		command.NewPutCommand(),
		command.NewDeleteCommand(),
	}

	app.Run(os.Args)

}
