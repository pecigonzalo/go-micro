package template

// Plugins is the plugins template used for new projects.
var Plugins = `package main

import (
	_ "go-micro.dev/plugins/registry/kubernetes/v4"
)
`
