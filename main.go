package main

import (
	"marketingwizard/internal/cli"
	"marketingwizard/internal/conf"
	"marketingwizard/internal/server"
)

func main() {
	env := cli.Parse()
	server.Start(conf.NewConfig(env))
}
