package main

import (
	"github.com/AcalephStorage/consuloretcd-kv/command"
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
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:   "address, A",
			Value: "192.168.200.90",
			Usage:  "the remote endpoint for the cluster",
			EnvVar: "CONSULORETCD_KV_ADDRESS",
		},
		cli.IntFlag{
			Name:   "port, p",
			Value: "8500",
			Usage:  "the port of the remote endpoint for the cluster",
			EnvVar: "CONSULORETCD_KV_PORT",
		},
		cli.StringFlag{
			Name:   "ttl, t",
			Value: "",
			Usage:  "the time to live value of that is expected",
			EnvVar: "CONSULORETCD_KV_TTL",
		},
	}
	app.Commands = []cli.Command{
		command.NewGetCommand(),
		command.NewPutCommand(),
		command.NewDeleteCommand(),
		command.NewCASCommand(),
	}

	app.Run(os.Args)
}
