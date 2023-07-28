package main

import (
	"os"

	"github.com/khaitranhq/ssh_manager/pkg/action"
	"github.com/sirupsen/logrus"
)

func main() {
	logrus.SetLevel(logrus.FatalLevel)

	args := os.Args[1:]
	logrus.Debug(args)
	action_type := args[0]

	switch action_type {
	case "save":
		action.SaveSSHConnection()
	}
}
