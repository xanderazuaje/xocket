package flags

import (
	"log"
	"os"
)

func ValidateFlags() {
	if This.ConfigFile == "" {
		log.Fatal("config file is required")
	}
	if len(This.RunType) == 0 {
		This.RunType = []RunType{RunAll}
	}
	if This.logFile != "" {
		file, err := os.OpenFile(This.logFile, os.O_RDONLY|os.O_CREATE, 644)
		if err != nil {
			log.Fatal(err)
		}
		os.Stdout = file
	}
}
