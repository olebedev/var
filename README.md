# var

Tool to fill in json/yaml stdin stream from the environment variables.

### Usage

Install: `go get github.com/olebedev/var` or download compiled binaries at [releases](https://github.com/olebedev/var/releases) page.

```
$ var --help
NAME:
   var - tool to fill in json/yaml stdin stream from the environment variables

USAGE:
   var < conf.yaml > env-conf.yaml

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


VERSION:
   0.1.1

AUTHOR(S):
   olebedev <ole6edev@gmail.com>

COMMANDS:
     help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --alias value, -a value   alias to be matched: <FROM>=<TO>
   --indent value, -i value  indent string(only for json files)
   --help, -h                show help
   --version, -v             print the version
```

### License

Apache 2.0 - http://www.apache.org/licenses/LICENSE-2.0
