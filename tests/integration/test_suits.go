package main

import (
	authTests "github.com/alexandr-io/backend/auth/tests"
	libraryTests "github.com/alexandr-io/backend/library/tests"
	userTests "github.com/alexandr-io/backend/user/tests"
)

type testFunc func(string, string) error

type funcChannel struct {
	Func    testFunc
	Channel chan bool
}

type suiteStruct struct {
	Name     string
	FuncChan funcChannel
}

type testSuitesStruct struct {
	Microservice string
	Suites       []suiteStruct
}

var testSuites = []testSuitesStruct{
	{
		Microservice: "AUTH",
		Suites: []suiteStruct{
			{
				Name: "AUTH_WORKING",
				FuncChan: funcChannel{
					Func:    authTests.ExecAuthWorkingTests,
					Channel: make(chan bool, 1),
				},
			},
			{
				Name: "AUTH_BAD_REQUEST",
				FuncChan: funcChannel{
					Func:    authTests.ExecAuthBadRequestTests,
					Channel: make(chan bool, 1),
				},
			},
			{
				Name: "AUTH_INCORRECT",
				FuncChan: funcChannel{
					Func:    authTests.ExecAuthIncorrectTests,
					Channel: make(chan bool, 1),
				},
			},
			{
				Name: "AUTH_LOGOUT",
				FuncChan: funcChannel{
					Func:    authTests.ExecAuthLogoutTests,
					Channel: make(chan bool, 1),
				},
			},
		},
	},
	{
		Microservice: "LIBRARY",
		Suites: []suiteStruct{
			{
				Name: "LIBRARY_WORKING",
				FuncChan: funcChannel{
					Func:    libraryTests.ExecLibraryWorkingTests,
					Channel: make(chan bool, 1),
				},
			},
			{
				Name: "LIBRARY_BAD_REQUEST",
				FuncChan: funcChannel{
					Func:    libraryTests.ExecLibraryBadRequestTests,
					Channel: make(chan bool, 1),
				},
			},
		},
	},
	{
		Microservice: "USER",
		Suites: []suiteStruct{
			{
				Name: "USER_BAD_REQUEST",
				FuncChan: funcChannel{
					Func:    userTests.ExecUserBadRequestTests,
					Channel: make(chan bool, 1),
				},
			},
		},
	},
}
