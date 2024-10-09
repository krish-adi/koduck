package paths

import (
	"log"
	"os"
	"path/filepath"
)

var HomeDir string
var KoduckDir string
var UserHomeDir string

func InitPaths() {
	var err error
	HomeDir, err = os.UserHomeDir()
	if err != nil {
		log.Fatal("Error getting user's home directory:", err)
	}

	KoduckDir = filepath.Join(HomeDir, ".koduck")
	if _, err := os.Stat(KoduckDir); os.IsNotExist(err) {
		err := os.Mkdir(KoduckDir, 0755)
		if err != nil {
			log.Fatalf("Failed to create $HOME/.koduck directory: %v", err)
		}
	} else if err != nil {
		log.Fatalf("Error checking $HOME/.koduck directory: %v", err)
	}

}
