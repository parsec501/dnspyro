package cli

import (
	"os"
	"regexp"
)

const (
	REGEX = `^([a-zA-Z0-9]+\.){1,2}[a-zA-Z0-9]+$`
)

func GetUserInput() (string, bool) {
	if len(os.Args) < 2 || len(os.Args) > 3 {
		PrintUsage()
	}

	match, _ := regexp.MatchString(REGEX, os.Args[1])
	if !match {
		PrintUsage()
	}

	var checkTXT bool = false
	if len(os.Args) == 3 {
		switch os.Args[2] {
		case "--txt":
			checkTXT = true
		case "-txt":
			checkTXT = true
		default:
			PrintUsage()
		}
	}

	return os.Args[1], checkTXT
}
