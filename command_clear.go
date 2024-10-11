package main

import "fmt"

func callbackClear(cfg *config, args ...string) error {
	fmt.Print("\033[H\033[2J")
	return nil
}
