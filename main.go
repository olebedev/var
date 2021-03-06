package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"io/ioutil"
	"os"
	"strings"

	"github.com/codegangsta/cli"
	"github.com/olebedev/config"
)

func main() {
	app := cli.NewApp()
	app.Usage = "tool to fill in json/yaml stdin stream from the environment variables"
	app.UsageText = `var [-a <FROM>=<TO> [-a <FROM>=<TO>]] < conf.yaml > env-conf.yaml

	 The tool reads stdin, lookups env variables, sets them up and
	 produce the result to stdout.

	 All dotted("some.path.to.the.value") paths in the input file
	 will be transformed to UPPERCASE and dot(".") will be replaced
	 to underscore("_") to lookup the env.
   For example, env lookup for this input:

	   {
	     "port": 5000,
		   "rabbitmq": {
			   "url": "..."
		   },
		   "proxies": [
			  	"...",
				  "..."
		   ]
	   }

	 will be:

	   port         -> PORT
	   rabbitmq.url -> RABBITMQ_URL
	   proxies      -> PROXIES_0, PROXIES_1
	 `
	app.Author = "olebedev <ole6edev@gmail.com>"
	app.Version = "0.1.1"
	app.Flags = []cli.Flag{
		cli.StringSliceFlag{
			Name:  "alias,a",
			Usage: "alias to be matched: <FROM>=<TO>",
		},
		cli.StringFlag{
			Name:  "indent,i",
			Usage: "indent string(only for json files)",
		},
	}
	app.Action = act
	app.Run(os.Args)
}

func act(ctx *cli.Context) error {
	for _, kv := range ctx.StringSlice("alias") {
		splitted := strings.Split(kv, "=")
		if len(splitted) == 2 {
			os.Setenv(splitted[1], os.Getenv(splitted[0]))
		}
	}

	var isYaml bool
	input, err := ioutil.ReadAll(bufio.NewReader(os.Stdin))
	if err != nil {
		return err
	}

	c, err := config.ParseJson(string(input))
	if err != nil {
		c, err = config.ParseYaml(string(input))
		if err != nil {
			return err
		}
		isYaml = true
	}
	c.Env()

	var j string
	if isYaml {
		j, err = config.RenderYaml(c.Root)
		if err != nil {
			return err
		}
	} else {
		j, err = config.RenderJson(c.Root)
		if err != nil {
			return err
		}

		if ctx.String("indent") != "" {
			var out bytes.Buffer
			json.Indent(&out, []byte(j), "", ctx.String("indent"))
			j = out.String()
		}
	}
	bytes.NewBuffer([]byte(j)).WriteTo(os.Stdout)
	return nil
}
