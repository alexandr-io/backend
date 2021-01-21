package main

import (
	"fmt"
	"github.com/urfave/cli/v2"
)

func getTestsByName(name string) map[string]funcChannel {
	var m = make(map[string]funcChannel)
	for _, ms := range testSuits {
		if ms.Microservice == name {
			for _, suit := range ms.Suits {
				m[suit.Name] = suit.FuncChan
			}
			return m
		}
		for _, suit := range ms.Suits {
			if suit.Name == name {
				m[suit.Name] = suit.FuncChan
				return m
			}
		}
	}
	fmt.Println(name + " not found. Run tests to see the list of test suits and microservices")
	return m
}

func allTests() map[string]funcChannel {
	var m = make(map[string]funcChannel)
	for _, ms := range testSuits {
		for _, suit := range ms.Suits {
			m[suit.Name] = suit.FuncChan
		}
	}
	return m
}

func joinMaps(m1 map[string]funcChannel, m2 map[string]funcChannel) map[string]funcChannel {
	for key, value := range m2 {
		m1[key] = value
	}
	return m1
}

func filters(c *cli.Context) map[string]funcChannel {
	var includeFuncsMap = allTests()

	if c.Bool("asynchronous") {
		asynchronous = true
	}
	if len(c.StringSlice("include")) != 0 {
		args := c.StringSlice("include")
		includeFuncsMap = make(map[string]funcChannel)
		for _, arg := range args {
			joinMaps(includeFuncsMap, getTestsByName(arg))
		}
	}
	if len(c.StringSlice("exclude")) != 0 {
		args := c.StringSlice("exclude")
		for _, arg := range args {
			tests := getTestsByName(arg)
			for key := range tests {
				delete(includeFuncsMap, key)
			}
		}
	}
	return includeFuncsMap
}
