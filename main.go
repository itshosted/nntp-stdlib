package nntpstdlib

import "regexp"

// ValidateGroup name will return true on valid or false when not valid
func ValidateGroup(group string) bool {
	r, _ := regexp.Compile("^[a-zA-Z0-9][a-zA-Z0-9-_+.]+[a-zA-Z0-9-+]$")
	if r.MatchString(group) == true {
		return true
	}
	return false
}
