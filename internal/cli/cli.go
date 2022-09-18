package cli

import (
  "flag"
  "fmt"
  "os"
  "marketingwizard/internal/logging"
)

func usage() {
  fmt.Print(`This program runs marketingwizard backend server.

Usage:

mw [arguments]

Supported arguments:

`)
  flag.PrintDefaults()
  os.Exit(1)
}

func Parse() string {
	flag.Usage = usage
	env := flag.String("env", "dev", `Sets run environment. Possible values are "dev" and "prod"`)
	flag.Parse()
	logging.ConfigureLogger(*env)
	if *env == "prod" {
		logging.SetGinLogToFile()
	}
	return *env
}