package command

import (
	"fmt"
	"os"

	"github.com/codegangsta/cli"
	"github.com/ashcrow/consuloretcd"
	"net/http"
)

func NewGetCommand() cli.Command {
	return cli.Command{
		Name:  "get",
		Usage: "retrieve the value of a key",
		Flags: []cli.Flag{
			cli.BoolFlag{
				Name:   "consul, c",
				Usage:  "use consul",
				EnvVar: "",
			},
			cli.BoolFlag{
				Name:   "etcd, e",
				Usage:  "use etcd",
				EnvVar: "",
			},
		},
		Action: func(c *cli.Context) {
			var client interface{}
			clientone := http.Client{}
			if c.Bool("consul") == true {
				client, _ := consuloretcd.NewClient("consul",
				consuloretcd.Config{
            		Endpoint: "http://127.0.0.1",
            		Client: clientone,
            		Port: 8500})
			}
			if c.Bool("etcd") == true {
				client, _ := consuloretcd.NewClient("etcd",
				consuloretcd.Config{
            		Endpoint: "http://127.0.0.1",
            		Client: clientone,
            		Port: 8500})
			}
			if len(c.Args()) == 0 {
				fmt.Fprintln(os.Stderr, "Error: Key required")
				os.Exit(1)
			}

			key := c.Args()[0]

			cc, _ := client.GetKey(key)
		},
	}
}
