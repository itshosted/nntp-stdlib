package nntpstdlib

import (
	"errors"
	"regexp"
	"strings"
)

const maxGroupLen = 80

var (
	overall   = regexp.MustCompile(`^[a-z][a-z0-9+_-]*\.[a-z0-9][a-z0-9+_-]*(?:\.[a-z0-9][a-z0-9+_-]*)*$`)
	compOK    = regexp.MustCompile(`^[a-z0-9+_-]+$`)
	allDigits = regexp.MustCompile(`^[0-9]+$`)
)

// ValidateGroup checks if the provided group name is valid according to NNTP standards.
// Based on the information from ftp://ftp.isc.org/pub/usenet/CONFIG/README
func ValidateGroup(group string) bool {
	// new: enforce total-length ≤ 80
	if len(group) > maxGroupLen {
		return false
	}
	if group != strings.ToLower(group) {
		return false
	}
	if !overall.MatchString(group) {
		return false
	}
	parts := strings.Split(group, ".")
	// first comp must start letter
	if parts[0][0] < 'a' || parts[0][0] > 'z' {
		return false
	}
	switch parts[0] {
	case "control", "to", "example":
		return false
	}
	for _, c := range parts {
		if !compOK.MatchString(c) {
			return false
		}
		if allDigits.MatchString(c) {
			return false
		}
		if c == "all" || c == "ctl" {
			return false
		}
	}
	return true
}

// IsMsgid will check if a valid message is given
func IsMsgid(arg string) (bool, error) {
	if len(arg) >= 3 && arg[0] == '<' && arg[len(arg)-1] == '>' {
		firstAt := strings.Index(arg, "@")

		if firstAt == -1 {
			// Missing @
			return false, errors.New("missing @ in msgid")
		}
		r, _ := regexp.Compile(`^[[:ascii:]]+$`)
		if !r.MatchString(arg) {
			return false, errors.New("none ASCII chars found in msgid")
		}
		// msgid
		return true, nil
	}
	// msgid
	return false, errors.New("invalid msgid")
}
