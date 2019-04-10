// Package mask provides utility functions for masking passwords within strings based on common patterns
package mask

import (
	"regexp"
	"strings"
)

var pattern = regexp.MustCompile(`(?im)(pass(word)?\s*[:=]\s*)([^\s]+)`)

// Password will mask any password within the given string and return it. Passwords are identified as following the phrase "password" or
// "pass" and an equal(=) or colon(:) sign, and are replaced with "****"
func Password(toMask string) string {
	return PasswordWith(toMask, "${1}****")
}

// Like mask.Password, but you get to decide the mask
func PasswordWith(toMask string, mask string) string {
	if !strings.HasPrefix(mask, "${1}") {
		mask = "${1}" + mask
	}
	return pattern.ReplaceAllString(toMask, mask)
}
