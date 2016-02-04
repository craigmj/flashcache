package devops

import (
	"flag"
	"fmt"
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
	user := fs.String("user", "flashcache", "User to run flashcache")
	group := fs.String("group", "flashcache", "Group to run flashcache")
	port := fs.Int("port", 16021, "Port to run flashcache on")
	bind := fs.String("bind", "127.0.0.1", "Address to which to bind webserver")

	return commander.NewCommand("upstart",
		"Output upstart script for flashcache",
		fs,
		func(args []string) error {
			return _upstartTemplate.Execute(os.Stdout, map[string]string{
				"Dir":   *wd,
				"User":  *user,
				"Group": *group,
				"Port":  fmt.Sprintf("%d", *port),
				"Bind":  *bind,
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
	setuid {{.User}}
	setgid {{.Group}}
	cd '{{.Dir}}'
	'{{.Dir}}/bin/flashcache' -logtostderr web -port {{.Port}} -bind {{.Bind}} &
end script

emits flashcache_starting
`))
