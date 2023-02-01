package app

import "github.com/urfave/cli"

// Generate returns cli application ready to exec
func Generate() *cli.App {
	app := cli.NewApp()
	app.Name = "Simple CLI App"
	app.Usage = "Search IPs and Names of Internet Servers."

	return app
}
