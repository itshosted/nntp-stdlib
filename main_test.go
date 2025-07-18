package nntpstdlib

import (
	"strings"
	"testing"
)

// TestValidateGroup tests the ValidateGroup function for various cases.
func TestValidateGroup(t *testing.T) {
	cases := []struct {
		name, group string
		want        bool
	}{
		// valid
		{"valid simple", "news.announce", true},
		{"three parts", "news.announce.newgroups", true},
		{"digits+letters", "comp1.comp2", true},
		{"symbols in middle", "a_b+1.c-d_e", true},

		// invalid
		{"uppercase", "News.announce", false},
		{"single component", "onlyone", false},
		{"first not letter", "1news.announce", false},
		{"pure digits comp", "news.123", false},
		{"middle digits", "news.123.more", false},
		{"bad first name", "control.group", false},
		{"bad first name", "to.group", false},
		{"bad first name", "example.test", false},
		{"reserved comp", "news.all", false},
		{"reserved comp", "news.ctl.sub", false},
		{"leading symbol", "+foo.bar", false},
	}

	for _, tt := range cases {
		got := ValidateGroup(tt.group)
		if got != tt.want {
			t.Errorf("%s: ValidateGroup(%q) = %v; want %v",
				tt.name, tt.group, got, tt.want)
		}
	}

	// --- length boundary tests ---

	// build a 27‑component valid group: each "a1", separated by '.', length = 27*2 + 26 = 80
	parts80 := make([]string, 27)
	for i := range parts80 {
		parts80[i] = "a1"
	}
	longValid := strings.Join(parts80, ".")
	if len(longValid) != 80 {
		t.Fatalf("test setup error: expected longValid length 80, got %d", len(longValid))
	}
	if !ValidateGroup(longValid) {
		t.Errorf("max length valid: ValidateGroup(%q) = false; want true", longValid)
	}

	// build a 28‑component group: length = 28*2 + 27 = 83 > 80
	parts83 := make([]string, 28)
	for i := range parts83 {
		parts83[i] = "a1"
	}
	longInvalid := strings.Join(parts83, ".")
	if len(longInvalid) <= 80 {
		t.Fatalf("test setup error: expected longInvalid >80, got %d", len(longInvalid))
	}
	if ValidateGroup(longInvalid) {
		t.Errorf("over length invalid: ValidateGroup(%q) = true; want false", longInvalid)
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
