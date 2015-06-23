package command

import (
	"fmt"
	"os"
	"github.com/codegangsta/cli"
)

func NewCASCommand() cli.Command {
	return cli.Command{
		Name:  "CAS",
		Usage: "Check if the the key exists, if not, set the value of a key",
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
			client, _ := makeClient(c)

			if len(c.Args()) == 0 {
				fmt.Fprintln(os.Stderr, "Error: Key required")
				os.Exit(1)
			}
			key := c.Args()[0]

			if len(c.Args()) < 2 {
				fmt.Fprintln(os.Stderr, "Error: Value required")
				os.Exit(1)
			}
			value := c.Args()[1]
			client.PutKey(key, value, KeyOptions{CASet:0})
		},
	}
}
