package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/nikonhub/bashelp/internal/openai"
	"github.com/schachmat/ingo"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatalf("Usage: %s <input>\n", os.Args[0])
	}

	apiKey := flag.String("api-key", "", "`APIKEY` to use for OpenAI API")
	instructions := flag.String("instructions", "You are a CLI assistant that provides Linux commands. Respond only with the exact command, without any explanation or additional text. If multiple commands are needed, separate them with '&&' or ';' as appropriate.", "`INSTRUCTIONS` for the assistant")

	if err := ingo.Parse("bashelp"); err != nil {
		log.Fatal(err)
	}

	input := os.Args[1]

	client := openai.NewClient(*apiKey)
	output, err := client.Complete(*instructions, input)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(output)
}
