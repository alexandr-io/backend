# Test Guidelines

This document will guide you to write your test.

We are using the official `testing` go package to write our tests. You can take a look at the documentation [here](https://golang.org/pkg/testing/)

### What should be tested?  
Every single route of the API must have his own battery of tests.  
Every independent features must also be tested.  
Each complexe integration of dependencies must be tested.  

### When should we write tests?
When a new feature is added to a Pull Request, it must come with it's tests.  
When a feature with a breacking changes is added to a Pull Request, it must come with it's tests.
