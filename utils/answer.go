package utils

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Answer string

func (a *Answer) IsYes() bool {
	lowercase := strings.ToLower(string(*a))
	return lowercase == "y" || lowercase == "yes"
}

func (a *Answer) IsNo() bool {
	lowercase := strings.ToLower(string(*a))
	return lowercase == "n" || lowercase == "no"
}

func AskQuestion(header string, question string) Answer {
	for true {
		if header != "" {
			fmt.Println(header)
		}
		fmt.Println(question)
		answer := ReadAnswer()
		if answer.IsYes() || answer.IsNo() {
			return answer
		}

		fmt.Println("Answer me property when I'm talking to you")
	}
	return "n"
}

func ReadAnswer() Answer {
	reader := bufio.NewReader(os.Stdin)
	line, _ := reader.ReadString('\n')
	return Answer(strings.TrimSuffix(line, "\n"))
}
