package command

import (
	"fmt"
	"os"

	"github.com/codegangsta/cli"
	"github.com/ashcrow/consuloretcd"
)

func NewGetCommand() cli.Command {
	return cli.Command{
		Name:  "get",
		Usage: "retrieve the value of a key",
		Action: func(c *cli.Context) {
			rawhandle(c, getCommandFunc)
		},
	}
}

func getCommandFunc(c *cli.Context, client *consuloretcd.Client) {
	if len(c.Args()) == 0 {
		fmt.Fprintln(os.Stderr, "Error: Key required")
		os.Exit(1)
	}

	key := c.Args()[0]

	cc, _ := client.GetKey(key)
	println(cc.value)
}