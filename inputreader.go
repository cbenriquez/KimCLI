package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var inputReader *bufio.Reader = bufio.NewReader(os.Stdin)

func ReadString() (*string, error) {
	in, err := inputReader.ReadString('\n')
	if err != nil {
		return nil, err
	}
	in = strings.TrimRight(in, "\r\n")
	return &in, nil
}

func ReadRange(min int, max int) (*int, error) {
	for {
		fmt.Printf("Enter an integer (%d - %d): ", min, max)
		in, err := ReadString()
		if err != nil {
			return nil, err
		}
		if *in == "!" {
			return nil, errors.New("exit read")
		}
		c, err := strconv.Atoi(*in)
		if err != nil {
			fmt.Println("invalid integer")
			continue
		}
		if c < min || c > max {
			fmt.Println("out of range")
			continue
		}
		return &c, nil
	}
}

func ParseInput(input string) (*string, *[]string) {
	var sentences []string
	var sentence string
	var inQuote bool
	for i, w := range input {
		if w == '"' {
			if !inQuote {
				inQuote = true
			} else {
				inQuote = false
			}
		} else if (w == ' ' && !inQuote) || i == len(input)-1 {
			if i == len(input)-1 {
				sentence += string(w)
			}
			sentences = append(sentences, sentence)
			sentence = ""
		} else {
			sentence += string(w)
		}
	}
	if len(sentences) == 0 {
		return nil, nil
	}
	args := sentences[1:]
	return &sentences[0], &args
}
