package main

import (
	"flag"

	"github.com/craigmj/commander"
	"github.com/golang/glog"

	"flash"
	"flash/devops"
)

func main() {
	flag.Parse()

	if err := commander.Execute(flag.Args(),
		flash.WebCommand,
		devops.UpstartCommand,
		devops.SetEnvCommand,
		devops.WriteEnvCommand,
	); nil != err {
		glog.Fatal(err)
	}
}
