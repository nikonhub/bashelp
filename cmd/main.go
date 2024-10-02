package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/nikonhub/bashelp/internal/config"
	"github.com/nikonhub/bashelp/internal/openai"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatalf("Usage: %s <input>\n", os.Args[0])
	}

	homeDir, err := os.UserHomeDir()

	if err != nil {
		log.Fatalf("Error gettings user home directory: %v", err)
	}

	configPath := filepath.Join(homeDir, ".bashelp", "config.yml")
	config, err := config.NewConfig(configPath)

	if err != nil {
		log.Fatalf("Error %v", err)
	}

	fmt.Println(config)
	os.Exit(0)

	input := os.Args[1]

	client := openai.NewClient(config.ApiKey)
	output, err := client.Complete(config.Instructions, input)

	if err != nil {
		panic(err)
	}

	fmt.Println(output)
}
