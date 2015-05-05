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
				Name:   "recursive, r",
				Usage:  "resursively delete keys",
				EnvVar: "",
			},
		},
		Action: func(c *cli.Context) {
			rawhandle(c, deleteCommandFunc)
		},
	}
}

func deleteCommandFunc(c *cli.Context, client *consuloretcd.Client) {
	var err error

	if len(c.Args()) == 0 {
		fmt.Fprintln(os.Stderr, "Error: Key required")
		os.Exit(1)
	}

	key := c.Args()[0]
	client.DeleteKey(key)
}