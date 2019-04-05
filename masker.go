// Package masker provides utility functions for masking passwords within strings based on common patterns
package masker

import (
	"regexp"
	"strings"
)

var pattern = regexp.MustCompile(`(?im)(pass(word)?\s*[:=]\s*)([^\s]+)`)

// Mask will mask any password within the given string and return it. Passwords are identified as following the phrase "password" or
// "pass" and an equal(=) or colon(:) sign, and are replaced with "****"
func Mask(toMask string) string {
	return MaskWith(toMask, "${1}****")
}

// Like masker.Mask, but you get to decide the replacing string
func MaskWith(toMask string, mask string) string {
	if !strings.HasPrefix(mask, "${1}") {
		mask = "${1}" + mask
	}
	return pattern.ReplaceAllString(toMask, mask)
}
