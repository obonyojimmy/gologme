package main

import (
	"os"
	"os/user"
	"path"

	"github.com/codegangsta/cli"
	"github.com/erasche/gologme/client"
)

func main() {
	app := cli.NewApp()
	app.Name = "gologme"
	app.Usage = "local logging client"
	user, err := user.Current()
	var defaultDbPath string
	if err != nil {
		defaultDbPath = "gologme.db"
	} else {
		defaultDbPath = path.Join(user.HomeDir, ".gologme.db")
	}

	app.Flags = []cli.Flag{
		cli.IntFlag{
			Name:  "windowLogGranularity",
			Value: 500,
			Usage: "How often to poll window title in ms",
		},
		cli.IntFlag{
			Name:  "keyLogGranularity",
			Value: 1000,
			Usage: "How often to aggregate caught keypresses in ms",
		},
		cli.BoolFlag{
			Name:  "standalone",
			Usage: "Run in non-networked, standalone mode",
		},
		cli.StringFlag{
			Name:  "dbPath",
			Usage: "Path to the database",
			Value: defaultDbPath,
		},
		cli.StringFlag{
			Name:  "serverAddr",
			Usage: "Address to send logs to, defaults to localhost for --standalone mode.",
			Value: "http://127.0.0.1:10000",
		},
	}

	app.Action = func(c *cli.Context) {
		if c.Bool("standalone") {
			go client.Serve(
				c.String("dbPath"),
				c.String("serverAddr"),
			)
		}

		serverPath := c.String("serverAddr")

		client.Golog(
			c.Int("windowLogGranularity"),
			c.Int("keyLogGranularity"),
			c.Bool("standalone"),
			serverPath,
		)

	}
	app.Run(os.Args)
}
