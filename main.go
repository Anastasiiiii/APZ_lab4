package main

import (
	"bufio"
	"os"
)

func main() {
	//eventLoop := new(engine.eventLoop)
	//eventLoop.Start()

	if input, err := os.Open("test.txt"); err == nil {
		defer input.Close()
		scanner := bufio.NewScanner(input)
		for scanner.Scan() {
			//commandLine := scanner.Text()
			//cmd := engine.Parse(commandLine)
			//eventLoop.Post(cmd)
		}
	}

	//eventLoop.AwaitFinish()
}
