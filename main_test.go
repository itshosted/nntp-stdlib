package nntpstdlib

import (
	"testing"
)

func TestValidateGroup(t *testing.T) {
	for group, valid := range map[string]bool{
		"alt":                                                true,
		"alt.group":                                          true,
		"alt.group.test":                                     true,
		"nl.kunst.sf+fantasy":                                true,  // + sign allowed
		"comp.lang.c++":                                      true,  // + sign at the end
		"a-lt.bin-aries.dvd-r":                               true,  // Hypen allowed
		"a-lt.bin-aries.dvd-r-":                              true,  // Hypen allowed at the end
		".alt.binaires.exact":                                false, // Don't end with a dot
		"alt.binaires.exact.":                                false, // Don't end with a dot
		"NewsUP::Article=HASH(0x55e3ab065098)->newsgroups()": false, // Don't allow this
	} {
		res := ValidateGroup(group)
		if res != valid {
			t.Errorf("validateGroup failed for %s expected %v but received %v", group, valid, res)
		}
	}
}

func TestIsMsgid(t *testing.T) {
	for msgid, valid := range map[string]bool{
		"<valid@msgid>": true,
		"<msgid>":       false, // Missing @
		"notvalid":      false, // Missing start end <>
		"<notvalid":     false, // Missing end >
		"notvalid>":     false, // Missing start <
		"12345":         false, // article id
		"<part74of205.Ip6&7Agv&HC&oBhSlekz@�test>": false,
	} {
		res, _ := IsMsgid(msgid)

		if res != valid {
			t.Errorf("validateMsgid failed for %s expected '%v' but received '%v'", msgid, valid, res)

		}
	}
}
