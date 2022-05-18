package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Welcome to KimCLI v0.2.0.")
	for {
		var wd []string
		if selectedCartoon != nil {
			wd = append(wd, selectedCartoon.ID)
			if selectedEpisode != nil {
				wd = append(wd, selectedEpisode.ID)
			}
		}
		fmt.Print(strings.Join(wd, "/") + "> ")
		in, err := ReadString()
		if err != nil {
			panic(err)
		}
		name, args := ParseInput(*in)
		if name == nil {
			continue
		}
		var commandFound bool
		for _, c := range commands {
			if c.Matches(*name) {
				if err := c.F(*args); err != nil {
					fmt.Printf("%s: %s\n", *name, err.Error())
				}
				commandFound = true
				break
			}
		}
		if !commandFound {
			fmt.Printf("%s: %s\n", *name, "command not found")
		}
	}
}
