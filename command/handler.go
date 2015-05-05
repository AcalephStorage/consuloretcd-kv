package command

import (
	"net/http"
	"net/url"
	"time"

	"github.com/codegangsta/cli"
	"github.com/ashcrow/consuloretcd"
)

type handlerFunc func(*cli.Context, *consuloretcd.Client)

func rawhandle(c *cli.Context, fn handlerFunc) {
	clientone := http.Client{}
//	timeout := c.GlobalInt("wait")
	client, _ := consuloretcd.NewClient("consul",
		consuloretcd.Config{
            Endpoint: "http://127.0.0.1",
            Client: clientone,
            Port: 8500})

	if timeout != 0 {
		handleTimeout(timeout, consulConfig)
	}

	fn(c, client)
}
/*
func handleTimeout(timeout int, conf api.Config) {
	req := url.URL{Scheme: conf.Scheme, Host: conf.Address}

	for i := 1; i <= timeout; i++ {
		_, err := http.Get(req.String())
		if err == nil {
			break
		}

		time.Sleep(1 * time.Second)
	}
}*/