package helpers

import (
	"io/ioutil"
	"log"
	"oreo/config"
	"os"
)

func SetUpLogsFile(path string) {
	log.Printf("SetUpLogsFile started. path: %s", path)
	if !config.Config.LogsEnabled {
		log.Printf("SetUpLogsFile. Disabling logs for performance.")
		log.SetOutput(ioutil.Discard)
	} else {
		log.Printf("SetUpLogsFile. Redirecting logs into file.")
		os.Remove(path)
		f, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
		if err != nil {
			log.Fatal("Can`t open/create logs file.")
		}
		log.SetOutput(f)
	}
}
