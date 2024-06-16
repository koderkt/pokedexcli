package main

import "os"

func cmdExit(config *Config) error {
	os.Exit(0)
	return nil
}
