package regex

import (
	"regexp"
	"strings"
)

var pattern = regexp.MustCompile(`(?im)(password|jwt|refreshtoken?\s*[:=]\s*)([^\s]+)`)

// Hide hide password, jwt and refresh token from the given string. Replaced by ****
func Hide(toHide string) string {
	mask := "${1}:\"****\""
	if !strings.HasPrefix(mask, "${1}") {
		mask = "${1}" + mask
	}
	return pattern.ReplaceAllString(toHide, mask)
}
