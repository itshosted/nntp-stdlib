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
