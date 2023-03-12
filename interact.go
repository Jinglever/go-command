package jgcmd

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// AskForAnswer asks the user a question and returns the answer.
// The options are presented as a list of options, and the defaultAnswer
// is the default answer if the user just presses enter. if defaultAnswer is "",
// then continuse ask until the user enters a valid answer.
// The answer is returned as a string.
func AskForAnswer(question string, options []string, defaultAnswer string) string {
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
			} else {
				fmt.Println("Please enter a valid option")
			}
		} else {
			for _, option := range options {
				if answer == option {
					return answer
				}
			}
			fmt.Println("Please enter a valid option")
		}
	}
}
