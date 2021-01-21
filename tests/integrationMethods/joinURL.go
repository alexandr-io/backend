package integrationMethods

import (
	"fmt"
	"path"
	"strings"
)

// JoinURL Join a base url with a route path
func JoinURL(base string, paths ...string) string {
	p := path.Join(paths...)
	return fmt.Sprintf("%s/%s", strings.TrimRight(base, "/"), strings.TrimLeft(p, "/"))
}
