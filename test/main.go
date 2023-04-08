package main

import (
	"fmt"

	jgcmd "github.com/Jinglever/go-command"
)

func main() {
	options := []string{"a", "b", "c"}
	answer := jgcmd.AskForOption("请选择一个选项", options, "a")
	fmt.Printf("你选择了：%s\n", answer)

	answer = jgcmd.AskForOption("请选择一个选项", options, "")
	fmt.Printf("你选择了：%s\n", answer)

	answer = jgcmd.AskForInput("请输入一个字符串", "default")
	fmt.Printf("你输入了：%s\n", answer)

	answer = jgcmd.AskForInput("请输入一个字符串", "")
	fmt.Printf("你输入了：%s\n", answer)
}
