package core

import (
	"log"
	"regexp"
)

func getRegex() *regexp.Regexp {
	regex, err := regexp.Compile("#/[cC]omponents/[sS]chemas/(.*)")
	if err != nil {
		log.Fatal("Could not compile Regex")
	}
	return regex
}

var regex = getRegex()

func GetSchemaName(id string) string {
	match := regex.FindStringSubmatch(id)
	if len(match) < 2 {
		return ""
	}
	return match[1]
}
