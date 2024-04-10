package helpers

import (
	"regexp"
	"strconv"
)

func IsValidVersion(version string) bool {
	if version == "" {
		return true
	}

	// check if the version is valid (only numbers and dots) with regex
	reg := regexp.MustCompile(`^[0-9.]+$`)
	return reg.MatchString(version)
}

func VersionExists(version, tool string) bool {
	if version == "" {
		return true
	}

	intVersion, _ := strconv.Atoi(version[:2])
	switch tool {
	case "fnm", "nvm":
		if intVersion >= 22 || intVersion <= 0 {
			return false
		}
	case "sdkman":
		if intVersion >= 22 || intVersion <= 7 {
			return false
		}
	case "pyenv":
		if intVersion >= 4 || intVersion <= 2 {
			return false
		}
	}

	return true
}
