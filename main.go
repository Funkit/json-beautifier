package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"github.com/alexflint/go-arg"
	"io/ioutil"
	"os"
)

var args struct {
	JSONFilePath string `arg:"-f" help:"path to the JSON file"`
}

func main() {
	arg.MustParse(&args)

	var rawContent []byte

	if args.JSONFilePath != "" {
		content, err := ioutil.ReadFile(args.JSONFilePath)
		if err != nil {
			panic(fmt.Errorf("error when reading the JSON file: %w", err))
		}

		rawContent = content
	} else {
		reader := bufio.NewReader(os.Stdin)
		var err error
		var input byte
		for {
			input, err = reader.ReadByte()
			if err != nil {
				break
			} else {
				rawContent = append(rawContent, input)
			}
		}
	}

	var output any
	err := json.Unmarshal(rawContent, &output)
	if err != nil {
		panic(fmt.Errorf("error when unmarshalling JSON content: %w", err))
	}

	formattedOutput, err := json.MarshalIndent(output, "", " ")

	fmt.Println(string(formattedOutput))
}
