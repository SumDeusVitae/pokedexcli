package main

import "fmt"

func callbackHistory(cfg *config, args ...string) error {
	history := cfg.history
	// index := cfg.current_history_index
	// fmt.Printf("Current index: %d\n", index)
	fmt.Println("Hisroty:")
	for _, command := range history {
		fmt.Printf("	%s\n", command)
	}

	return nil
}
