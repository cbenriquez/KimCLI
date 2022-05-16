package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

func main() {
	var p Prompt
	p.BufioReader = bufio.NewReader(os.Stdin)
	p.Report = func(name, message string) {
		fmt.Printf("%s: %s\n", name, message)
	}
	p.Commands = []Command{
		{
			Names: []string{"exit", "e", "quit", "q"},
			F:     CMDExit,
		},
		{
			Names: []string{"search", "s", "find", "f"},
			F:     CMDSearch,
		},
		{
			Names: []string{"episodes", "eps"},
			F:     CMDEpisodes,
		},
		{
			Names: []string{"watch", "w", "play", "p"},
			F:     CMDWatch,
		},
	}
	p.Run()
}

func CMDExit(p *Prompt, _ []string) error {
	p.Exit = true
	return nil
}

func CMDSearch(p *Prompt, args []string) error {
	carts, err := SearchCartoons(strings.Join(args, " "))
	if err != nil {
		return err
	}
	if len(*carts) == 0 {
		return errors.New("empty result")
	}
	for i, cart := range *carts {
		tt, err := cart.Title()
		if err != nil {
			return err
		}
		fmt.Printf("%d) %s\n", i, *tt)
	}
	c, err := p.ReadRange(0, len(*carts)-1, "pick a cartoon (number): ")
	if err != nil {
		return err
	}
	p.CurrentCartoon = &(*carts)[*c]
	p.CurrentEpisode = nil
	return nil
}

func CMDEpisodes(p *Prompt, args []string) error {
	if len(args) == 1 {
		p.CurrentCartoon = &Cartoon{args[0], nil, nil}
		p.CurrentEpisode = nil
	}
	if p.CurrentCartoon == nil {
		return errors.New("cartoon unspecified")
	}
	eps, err := p.CurrentCartoon.Episodes()
	if err != nil {
		return err
	}
	if len(*eps) == 0 {
		return errors.New("empty result")
	}
	for i, ep := range *eps {
		fmt.Printf("%d) %s\n", i, ep.Name)
	}
	c, err := p.ReadRange(0, len(*eps)-1, "pick an episode (number): ")
	if err != nil {
		return err
	}
	p.CurrentEpisode = &(*eps)[*c]
	return nil
}

func CMDWatch(p *Prompt, args []string) error {
	if p.CurrentCartoon == nil {
		return errors.New("cartoon ID not specified")
	}
	if len(args) == 1 {
		eps, err := p.CurrentCartoon.Episodes()
		if err != nil {
			return err
		}
		for _, ep := range *eps {
			if ep.ID == args[0] {
				p.CurrentEpisode = &ep
				break
			}
		}
		if p.CurrentEpisode == nil {
			return errors.New("invalid episode ID")
		}
	} else if p.CurrentEpisode == nil {
		return errors.New("episode ID not specified")
	}
	vids, err := p.CurrentEpisode.Videos()
	if err != nil {
		return err
	}
	for i, vid := range *vids {
		fmt.Printf("%d) %s - %s\n", i, vid.Label, vid.Type)
	}
	c, err := p.ReadRange(0, len(*vids)-1, "pick a video (number): ")
	if err != nil {
		return err
	}
	vid := (*vids)[*c]
	if err := vid.Play(); err != nil {
		return err
	}
	return nil
}
