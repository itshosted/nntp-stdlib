package nntpstdlib

import (
	"errors"
	"regexp"
	"strings"
)

// ValidateGroup name will return true on valid or false when not valid
func ValidateGroup(group string) bool {
	r, _ := regexp.Compile("^[a-zA-Z0-9][a-zA-Z0-9-_+.]+[a-zA-Z0-9-+]$")
	if r.MatchString(group) == true {
		return true
	}
	return false
}

// IsMsgid will check if a valid message is given
func IsMsgid(arg string) (bool, error) {
	if len(arg) >= 3 && arg[0] == '<' && arg[len(arg)-1] == '>' {
		firstAt := strings.Index(arg, "@")

		if firstAt == -1 {
			// Missing @
			return false, errors.New("Missing @ in msgid")
		}
		r, _ := regexp.Compile(`^[[:ascii:]]+$`)
		if r.MatchString(arg) == false {
			return false, errors.New("none ASCII chars found in msgid")
		}
		// msgid
		return true, nil
	}
	// msgid
	return false, errors.New("Invalid msgid")
}
