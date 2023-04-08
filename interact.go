package jgcmd

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// AskForOption asks the user a question and returns the answer.
// The options are presented as a list of options, and the defaultAnswer
// is the default answer if the user just presses enter. if defaultAnswer is "",
// then continuse ask until the user enters a valid answer.
// The answer is returned as a string.
func AskForOption(question string, options []string, defaultAnswer string) string {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Printf("%s [%s]: ", question, strings.Join(options, "/"))
		answer, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading input: ", err)
			continue
		}

		answer = strings.TrimSpace(answer)
		if answer == "" {
			if defaultAnswer != "" {
				return defaultAnswer
			}
		} else {
			for _, option := range options {
				if answer == option {
					return answer
				}
			}
		}
	}
}

// AskForInput asks the user a question and returns the answer.
// The answer is returned as a string.
func AskForInput(question string, defaultAnswer string) string {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Printf("%s: ", question)
		answer, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading input: ", err)
			continue
		}

		answer = strings.TrimSpace(answer)
		if answer == "" {
			if defaultAnswer != "" {
				return defaultAnswer
			}
		} else {
			return answer
		}
	}
}
