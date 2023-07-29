package action

import (
	"fmt"
	"os"

	"github.com/khaitranhq/ssh_manager/configs"
	"github.com/khaitranhq/ssh_manager/pkg/utils"
	"github.com/sirupsen/logrus"
)

func enterGroupName(groupNames []string) string {
	groupSize := len(groupNames) + 1

	for i, group := range groupNames {
		fmt.Printf("[%d] %s\n", i, group)
	}

	fmt.Printf("[%d] Create a new group\n", groupSize)

	var groupIndex int
	fmt.Print(
		"Please choose group index you want to save SSH connection on or choose the last option to create a new one: ",
	)
	fmt.Scanln(&groupIndex)

	var groupName string
	if groupIndex == groupSize {
		fmt.Print("Enter the new group name: ")
		fmt.Scanln(&groupName)
	}

	return groupName
}

func enterConnection() configs.Connection {
	var connection configs.Connection

	// Hostname
	fmt.Print("Enter hostname: ")
	fmt.Scanln(&connection.Host)

	// User
	fmt.Print("Enter SSH user: ")
	fmt.Scanf("%s", &connection.UserName)

	// Private key
	var keyPath string
	fmt.Print("Enter private key directory (leave blank if you use password): ")
	fmt.Scanf("%s", &keyPath)

	if keyPath != "" {
		if _, err := os.Stat(configs.KEYS_PATH); err != nil {
			os.MkdirAll(configs.KEYS_PATH, 0744)
		}
		_, err := utils.Copy(keyPath, configs.KEYS_PATH)
		if err != nil {
			logrus.Error(err)
		}
	}

	// Password
	fmt.Print("Enter password (leave blank if you use private key): ")
	fmt.Scanf("%s", &connection.Password)

	return connection
}

func SaveSSHConnection(groups []configs.Group) {
	groupName := enterGroupName([]string{})
	logrus.Debug(groupName)

	connection := enterConnection()
	logrus.Debug(connection)

	foundGroup := false
	for i, group := range groups {
		if group.GroupName == groupName {
			groups[i].Connections = append(groups[i].Connections, connection)
			foundGroup = true
		}
	}

	if !foundGroup {
		groups = append(groups, configs.Group{GroupName: groupName,
			Connections: []configs.Connection{connection},
		},
		)
	}

	logrus.Debug(groups)
}
