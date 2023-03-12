package main

import (
	"fmt"

	jgcmd "github.com/Jinglever/go-command"
)

func main() {
	options := []string{"a", "b", "c"}
	answer := jgcmd.AskForAnswer("请选择一个选项", options, "a")
	fmt.Printf("你选择了：%s\n", answer)
}
