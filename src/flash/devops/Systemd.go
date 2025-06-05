package devops

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"text/template"

	"github.com/craigmj/commander"
	"github.com/golang/glog"
)

func SystemdCommand() *commander.Command {
	fs := flag.NewFlagSet("systemd", flag.ExitOnError)

	bindir := filepath.Dir(os.Args[0])
	if !filepath.IsAbs(bindir) {
		cwd, err := os.Getwd()
		if nil!=err {
			glog.Fatal(err)
		}
		bindir = filepath.Clean(filepath.Join(cwd, bindir))
	}
	execName := filepath.Base(os.Args[0])

	wd := fs.String("dir", bindir, "Directory of flashcache executable")
	exc := fs.String("exec", execName, "Executable name of flashcache executable")
	user := fs.String("user", "flashcache", "User to run flashcache")
	group := fs.String("group", "flashcache", "Group to run flashcache")
	port := fs.Int("port", 16021, "Port to run flashcache on")
	bind := fs.String("bind", "127.0.0.1", "Address to which to bind webserver")

	return commander.NewCommand("systemd",
		"Output systemd script for flashcache",
		fs,
		func(args []string) error {
			return _systemdTemplate.Execute(os.Stdout, map[string]string{
				"Me": *exc,
				"Dir":   *wd,
				"User":  *user,
				"Group": *group,
				"Port":  fmt.Sprintf("%d", *port),
				"Bind":  *bind,
			})
		})
}

var _systemdTemplate = template.Must(template.New("").Parse(
	`
[Unit]
Description=flashcache

[Install]
WantedBy=multi-user.target

[Service]
Type=simple
User={{.User}}
Group={{.Group}}
WorkingDir="{{.Dir}}"
ExecStart="{{.Dir}}/{{.Me}}" -logtostderr web -port {{.Port}} -bind {{.Bind}}

`))
