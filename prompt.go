package main

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

type Prompt struct {
	CurrentCartoon *Cartoon
	CurrentEpisode *Episode
	Commands       []Command
	BufioReader    *bufio.Reader
	Report         func(name string, message string)
	Exit           bool
}

func (p *Prompt) Run() {
	for {
		var wd []string
		if p.CurrentCartoon != nil {
			wd = append(wd, p.CurrentCartoon.ID)
		}
		if p.CurrentEpisode != nil {
			wd = append(wd, p.CurrentEpisode.ID)
		}
		fmt.Printf("%s> ", strings.Join(wd, "/"))
		in, err := p.ReadString()
		if err != nil {
			panic(err)
		}
		args := strings.Split(*in, " ")
		p.Execute(args[0], args[1:])
		if p.Exit {
			break
		}
	}
}

func (p *Prompt) ReadString() (*string, error) {
	in, err := p.BufioReader.ReadString('\n')
	if err != nil {
		return nil, err
	}
	in = in[:strings.LastIndex(in, "\n")]
	return &in, nil
}

func (p *Prompt) ReadRange(min int, max int, prefix string) (*int, error) {
	for {
		fmt.Print(prefix)
		in, err := p.ReadString()
		if err != nil {
			return nil, err
		}
		c, err := strconv.Atoi(*in)
		if err != nil {
			fmt.Println("invalid number")
			continue
		}
		if c > max || c < min {
			fmt.Println("out of range")
			continue
		}
		return &c, nil
	}
}

func (p *Prompt) Execute(name string, args []string) {
	name = strings.ToLower(name)
	for _, c := range p.Commands {
		if c.Matches(name) {
			err := c.F(p, args)
			if err != nil {
				p.Report(name, err.Error())
			}
			return
		}
	}
	p.Report(name, "command not found")
}
