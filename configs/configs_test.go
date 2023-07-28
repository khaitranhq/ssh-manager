package configs

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadConfigsWhenNo(t *testing.T) {
	/**
		1. When config file no exist	
	*/
	if _, err := os.Stat(CONFIG_FILE_PATH); err == nil {
		os.Remove(CONFIG_FILE_PATH)
	}

	
	configs := LoadConfigs()
	assert.Equal(t, []Group{}, configs)

	// Check file is created
	if _, err := 

	// Load successfully

}
