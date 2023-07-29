package main

import (
	"os"

	"github.com/khaitranhq/ssh_manager/configs"
	"github.com/khaitranhq/ssh_manager/pkg/action"
	"github.com/sirupsen/logrus"
)

func main() {
	logrus.SetLevel(logrus.DebugLevel)

	args := os.Args[1:]
	logrus.Debug(args)
	action_type := args[0]

	groups := configs.LoadConfigs()

	switch action_type {
	case "save":
		action.SaveSSHConnection(groups)
	}
}
