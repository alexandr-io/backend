package itgmtod

import (
	"fmt"

	"github.com/fatih/color"
)

var (
	green   = color.New(color.FgGreen).SprintFunc()
	red     = color.New(color.FgRed).SprintFunc()
	cyan    = color.New(color.FgCyan).SprintfFunc()
	magenta = color.New(color.FgHiMagenta).SprintfFunc()
	// BackCyan is used by AUTH
	BackCyan = color.New(color.BgCyan).Add(color.FgBlack).SprintfFunc()
	// BackBlue is used by USER
	BackBlue = color.New(color.BgBlue).Add(color.FgWhite).SprintfFunc()
)

func getHTTPVerbColor(verb string) string {
	coloredVerb := verb
	switch verb {
	case "GET":
		coloredVerb = green(verb)
		break
	case "POST":
		coloredVerb = cyan(verb)
		break
	case "PUT":
		coloredVerb = magenta(verb)
		break
	case "DELETE":
		coloredVerb = red(verb)
		break
	}
	return coloredVerb
}

// NewSuccessMessage print a success message for an integration Test
func NewSuccessMessage(MS string, verb string, route string, testSuit string) {
	fmt.Printf("%-25s%-20s%-55s%-15s%5s\n", MS, getHTTPVerbColor(verb), route, testSuit, green("✓"))
}

// NewFailureMessage print a failure message for an integration Test
func NewFailureMessage(MS string, verb string, route string, testSuit string, errMessage string) {
	fmt.Printf("%-25s%-20s%-55s%-15s%5s\t%s\n", MS, getHTTPVerbColor(verb), route, testSuit, red("✗"), errMessage)
}
