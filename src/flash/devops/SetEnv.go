package devops

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
	"text/template"

	"github.com/craigmj/commander"
)

func SetEnvScript(out io.Writer, server string, port int) error {
	return _setenvTemplate.Execute(out, map[string]interface{}{
		"Server": server,
		"Port":   port,
	})
}

func WriteEnvironment(server string, port int) error {
	buf, err := ioutil.ReadFile("/etc/environment")
	if nil != err && !os.IsNotExist(err) {
		return err
	}
	if nil != err {
		buf = []byte{}
	}
	foundLine := false
	in := bufio.NewScanner(bytes.NewReader(buf))
	out, err := os.Create("/etc/environment")
	if nil != err {
		return err
	}
	defer out.Close()

	cacheServer := fmt.Sprintf(`FLASHCACHE_SERVER="http://%s:%d/"`, server, port)

	for in.Scan() {
		l := in.Text()
		if strings.HasPrefix(l, "FLASHCACHE_SERVER=") {
			fmt.Fprintln(out, cacheServer)
			foundLine = true
		} else {
			fmt.Fprintln(out, l)
		}
	}
	if !foundLine {
		fmt.Fprintln(out, cacheServer)
	}
	return nil
}

func SetEnvCommand() *commander.Command {
	fs := flag.NewFlagSet("setenv", flag.ExitOnError)
	port := fs.Int("port", 16021, "Port on which flashcache runs")
	server := fs.String("server", "127.0.0.1", "Server location of flashcache (ip or fqdn)")

	return commander.NewCommand("setenv",
		"Return a script to set FlashCache environment values",
		fs,
		func(args []string) error {
			return SetEnvScript(os.Stdout, *server, *port)
		})
}

func WriteEnvCommand() *commander.Command {
	fs := flag.NewFlagSet("write-env", flag.ExitOnError)
	port := fs.Int("port", 16021, "Port on which flashcache runs")
	server := fs.String("server", "127.0.0.1", "Server location of flashcache (ip or fqdn)")

	return commander.NewCommand("write-env",
		"Write /etc/environment to set. FLASHCACHE_SERVER value",
		fs,
		func(args []string) error {
			return WriteEnvironment(*server, *port)
		})
}

var _setenvTemplate = template.Must(template.New("").Parse(
	`#!/bin/bash
export FLASHCACHE_SERVER=http://{{.Server}}:{{.Port}}/
`))
