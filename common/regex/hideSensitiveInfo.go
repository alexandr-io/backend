package regex

import (
	"regexp"
	"strings"
)

var pattern = regexp.MustCompile(`(?im)(password|jwt|refreshtoken?\s*[:=]\s*)([^\s]+)`)

func Hide(toHide string) string {
	mask := "${1}:\"****\""
	if !strings.HasPrefix(mask, "${1}") {
		mask = "${1}" + mask
	}
	return pattern.ReplaceAllString(toHide, mask)
}
