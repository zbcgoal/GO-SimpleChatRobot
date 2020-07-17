package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	inputbuffer := bufio.NewReader(os.Stdin)
	fmt.Println("please input your name:")
	input,err := inputbuffer.ReadString('\n')
	if err!=nil {
		fmt.Println(err)
		os.Exit(1)
	}else {
		username := input[:len(input)-1]
		fmt.Printf("hello %s what can i do for you?",username)
	}
	for {
		input, err = inputbuffer.ReadString('\n')
		if err!=nil {
			fmt.Println(err)
			os.Exit(1)
		}
		input = input[:len(input)-1]
		input = strings.ToLower(input)
		switch input {
		case "":
			continue
		case "bye","byebye","nothing":
			fmt.Println("ok,byebye")
			os.Exit(1)
		case "please give me a cake":
			fmt.Println("here is")
		default:
			fmt.Println("please say again")
		}
	}


}
