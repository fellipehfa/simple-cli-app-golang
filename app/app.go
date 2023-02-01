package app

import (
	"fmt"
	"log"
	"net"

	"github.com/urfave/cli"
)

// Generate returns cli application ready to exec
func Generate() *cli.App {
	app := cli.NewApp()
	app.Name = "Simple CLI App"
	app.Usage = "Search IPs and Names of Internet Servers."

	app.Commands = []cli.Command{
		{
			Name:  "ip",
			Usage: "Search IPs of Internet Servers.",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "host",
					Value: "devbook.com.br",
				},
			},
			Action: searchIps,
		},
	}

	return app
}

func searchIps(c *cli.Context) {
	host := c.String("host")

	ips, err := net.LookupIP(host)
	if err != nil {
		log.Fatal(err)
	}

	for _, ip := range ips {
		fmt.Println(ip)
	}
}
