package command

import (
	"fmt"
	"os"

	"github.com/codegangsta/cli"
	"github.com/ashcrow/consuloretcd"
)



func NewPutCommand() cli.Command {
	return cli.Command{
		Name:  "put",
		Usage: "set the value of a key",
		Action: func(c *cli.Context) {
			rawhandle(c, putCommandFunc)
		},
	}
}

func putCommandFunc(c *cli.Context, client *api.Client) {
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

	cp, _ := client.PutKey("key","value")
	println("Added ",cp.key," ",cp.value)

}