package command

import (
	"fmt"
	"os"

	"github.com/codegangsta/cli"
	"github.com/ashcrow/consuloretcd"
)

// NewDeleteCommand prepares the "delete" command
func NewDeleteCommand() cli.Command {
	return cli.Command{
		Name:  "delete",
		Usage: "delete a key",
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
			if c.Bool("consul") == true {
				deleteCommandFuncConsul(c *cli.Context)
			}
			if c.Bool("etcd") == true {
				deleteCommandFuncEtcd(c *cli.Context)
			}
			
		},
	}
}

func deleteCommandFuncConsul(c *cli.Context) {

	clientone := http.Client{}
	client, _ := consuloretcd.NewClient("consul",
	consuloretcd.Config{
            Endpoint: "http://127.0.0.1",
            Client: clientone,
            Port: 8500})

	if len(c.Args()) == 0 {
		fmt.Fprintln(os.Stderr, "Error: Key required")
		os.Exit(1)
	}

	key := c.Args()[0]
	client.DeleteKey(key)
}

func deleteCommandFuncEtcd(c *cli.Context) {

	clientone := http.Client{}
	client, _ := consuloretcd.NewClient("etcd",
	consuloretcd.Config{
            Endpoint: "http://127.0.0.1",
            Client: clientone,
            Port: 8500})

	if len(c.Args()) == 0 {
		fmt.Fprintln(os.Stderr, "Error: Key required")
		os.Exit(1)
	}

	key := c.Args()[0]
	client.DeleteKey(key)
}