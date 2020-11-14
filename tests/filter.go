package main

import "github.com/urfave/cli/v2"

type microservices string
type microservicesFunc func(string) error

const (
	auth microservices = "auth"
	user microservices = "user"
)

func (ms microservices) string() string {
	return string(ms)
}

var microservicesIncludeMap = map[microservices]microservicesFunc{
	auth: execAuth,
	user: execUser,
}

func filters(c *cli.Context) ([]microservicesFunc, error) {
	var includeFuncs []microservicesFunc

	if len(c.StringSlice("include")) != 0 {
		args := c.StringSlice("include")
		for _, arg := range args {
			if function, ok := microservicesIncludeMap[microservices(arg)]; ok {
				includeFuncs = append(includeFuncs, function)
			} else {
				return nil, cli.Exit(arg+" not recognized", 1)
			}
		}
	} else if len(c.StringSlice("exclude")) != 0 {
		args := c.StringSlice("exclude")
		excludeMap := microservicesIncludeMap
		for _, arg := range args {
			if _, ok := excludeMap[microservices(arg)]; ok {
				delete(excludeMap, microservices(arg))
			} else {
				return nil, cli.Exit(arg+" not recognized", 1)
			}
		}
		for _, element := range excludeMap {
			includeFuncs = append(includeFuncs, element)
		}
	}
	return includeFuncs, nil
}
