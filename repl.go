package main

import (
	"fmt"
	"strings"

	"github.com/chzyer/readline"
)

func startRepl(cfg *config) {
	// Create a readline instance
	rl, err := readline.NewEx(&readline.Config{
		Prompt: "Pokedex > ",
		// HistoryFile: "history.txt", // Save history to a file
	})
	if err != nil {
		fmt.Println("Error creating readline instance:", err)
		return
	}
	defer rl.Close()

	for {
		line, err := rl.Readline()
		if err != nil {
			fmt.Println("Error reading line:", err)
			continue
		}

		words := cleanInput(line)
		if len(words) == 0 {
			continue
		}

		call := strings.Join(words, " ")

		if len(cfg.history) < 1 || cfg.history[cfg.current_history_index] != call {
			for i, s := range cfg.history {
				if s == call {
					if i < len(cfg.history) {
						cfg.history = append(cfg.history[:i], cfg.history[i+1:]...)
					}
				}
			}
			cfg.history = append(cfg.history, call)
			cfg.current_history_index = len(cfg.history) - 1
		}

		commandName := words[0]
		args := []string{}
		if len(words) > 1 {
			args = words[1:]
		}

		command, exists := getCommands()[commandName]
		if exists {
			err := command.callback(cfg, args...)
			if err != nil {
				fmt.Println(err)
			}
		} else {
			fmt.Printf("Unknown command: %s\n", commandName)
		}
	}
}

func cleanInput(text string) []string {
	output := strings.ToLower(text)
	words := strings.Fields(output)
	return words
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config, ...string) error
}
