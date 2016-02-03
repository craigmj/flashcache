package devops

import (
	"flag"
	"os"
	"text/template"

	"github.com/craigmj/commander"
	"github.com/golang/glog"
)

func UpstartCommand() *commander.Command {
	fs := flag.NewFlagSet("upstart", flag.ExitOnError)

	cwd, err := os.Getwd()
	if nil != err {
		glog.Fatal(err)
	}

	wd := fs.String("dir", cwd, "Directory of flashcache executable")

	return commander.NewCommand("upstart",
		"Output upstart script for flashcache",
		fs,
		func(args []string) error {
			return _upstartTemplate.Execute(os.Stdout, map[string]string{
				"Dir": *wd,
			})
		})
}

var _upstartTemplate = template.Must(template.New("").Parse(
	`
# flashcache
description "flashcache"

start on runlevel [2345]
stop on runlevel [!2345]

respawn
expect fork

script
	cd '{{.Dir}}'
	'{{.Dir}}/bin/flashcache' -logtostderr web &
end script

emits flashcache_starting
`))
