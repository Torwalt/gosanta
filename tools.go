//go:build tools
// +build tools

package main

import (
	_ "github.com/golang/mock/mockgen"
	_ "gotest.tools/gotestsum"
)
