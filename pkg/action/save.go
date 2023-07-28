package action

import (
	"fmt"
)

func enterGroupIndex(groups []string) int {
	groupSize := len(groups) + 1

	for i, group := range groups {
		fmt.Printf("[%d] %s\n", i, group)
	}

	fmt.Printf("[%d] Create a new group\n", groupSize)

	var groupIndex int
	fmt.Println("Please choose group you want to save SSH connection on:")
	fmt.Scanf("%d", &groupIndex)

	if groupIndex == groupSize {
		fmt.Println("Enter the new group name: ")

		var newGroupName string
		fmt.Scanf("%s", &newGroupName)
	}

	return groupIndex
}

func SaveSSHConnection() {
	enterGroupIndex([]string{})
}
