package main

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

var commands []Command = []Command{
	{
		Names: []string{"exit", "e", "quit", "q"},
		F:     Exit,
	},
	{
		Names: []string{"search", "s", "find", "f"},
		F:     Search,
	},
	{
		Names: []string{"episodes", "eps"},
		F:     Episodes,
	},
	{
		Names: []string{"play", "p", "watch", "w"},
		F:     Play,
	},
	{
		Names: []string{"play-highest", "ph", "watch-highest", "wh"},
		F:     PlayHighest,
	},
	{
		Names: []string{"next", "n"},
		F:     Next,
	},
	{
		Names: []string{"back", "b"},
		F:     Back,
	},
	{
		Names: []string{"list-episodes", "le"},
		F:     ListEpisodes,
	},
	{
		Names: []string{"first-episode", "fstep"},
		F:     FirstEpisode,
	},
	{
		Names: []string{"last-episode", "lstep"},
		F:     LastEpisode,
	},
}

func main() {
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

func Exit(_ []string) error {
	os.Exit(0)
	return nil
}

func Search(args []string) error {
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
	c, err := ReadRange(0, len(*carts)-1)
	if err != nil {
		return err
	}
	ResetSelection()
	selectedCartoon = &(*carts)[*c]
	return nil
}

func Episodes(args []string) error {
	if len(args) > 0 {
		ResetSelection()
		selectedCartoon = NewCartoon(strings.Join(args, "-"))
	}
	if err := IsCartoonSelected(); err != nil {
		return err
	}
	eps, err := selectedCartoon.Episodes()
	if err != nil {
		return err
	}
	if len(*eps) == 0 {
		return errors.New("empty result")
	}
	for i, ep := range *eps {
		fmt.Printf("%d) %s\n", i, ep.Name)
	}
	c, err := ReadRange(0, len(*eps)-1)
	if err != nil {
		return err
	}
	selectedEpisode = &(*eps)[*c]
	return nil
}

func Play(args []string) error {
	if err := IsCartoonSelected(); err != nil {
		return err
	}
	if len(args) > 0 {
		if err := SelectEpisode(strings.Join(args, "-")); err != nil {
			return err
		}
	} else if selectedEpisode == nil {
		return errors.New("episode ID not specified")
	}
	vids, err := selectedEpisode.Videos()
	if err != nil {
		return err
	}
	if len(*vids) == 0 {
		return errors.New("no videos found")
	}
	for i, vid := range *vids {
		fmt.Printf("%d) %s - %s\n", i, vid.Label, vid.Type)
	}
	c, err := ReadRange(0, len(*vids)-1)
	if err != nil {
		return err
	}
	vid := (*vids)[*c]
	if err := vid.Play(); err != nil {
		return err
	}
	return nil
}

func PlayHighest(args []string) error {
	if err := IsCartoonSelected(); err != nil {
		return err
	}
	if len(args) > 0 {
		if err := SelectEpisode(strings.Join(args, "-")); err != nil {
			return err
		}
	} else if selectedEpisode == nil {
		return errors.New("episode ID not specified")
	}
	vids, err := selectedEpisode.Videos()
	if err != nil {
		return err
	}
	if len(*vids) == 0 {
		return errors.New("no videos found")
	}
	vid := (*vids)[len(*vids)-1]
	if err := vid.Play(); err != nil {
		return err
	}
	return nil
}

func Next(_ []string) error {
	if err := IsEpisodeSelected(); err != nil {
		return err
	}
	eps, err := selectedCartoon.Episodes()
	if err != nil {
		return err
	}
	var ei int
	for i, ep := range *eps {
		if ep == *selectedEpisode {
			ei = i
			break
		}
	}
	if ei == 0 {
		return errors.New("reached last episode")
	}
	selectedEpisode = &(*eps)[ei-1]
	return nil
}

func Back(_ []string) error {
	if err := IsEpisodeSelected(); err != nil {
		return err
	}
	eps, err := selectedCartoon.Episodes()
	if err != nil {
		return err
	}
	var ei int
	for i, ep := range *eps {
		if ep == *selectedEpisode {
			ei = i
			break
		}
	}
	if ei == len(*eps)-1 {
		return errors.New("reached first episode")
	}
	selectedEpisode = &(*eps)[ei+1]
	return nil
}

func ListEpisodes(args []string) error {
	var c *Cartoon
	ai := strings.Join(args, "-")
	if len(args) > 1 && !strings.EqualFold(ai, selectedCartoon.ID) {
		c = NewCartoon(ai)
	} else if err := IsCartoonSelected(); err != nil {
		return err
	} else {
		c = selectedCartoon
	}
	eps, err := c.Episodes()
	if err != nil {
		return err
	}
	for i, ep := range *eps {
		fmt.Printf("%d) %s\n", i, ep.Name)
	}
	return nil
}

func FirstEpisode(args []string) error {
	if len(args) > 0 {
		ResetSelection()
		selectedCartoon = NewCartoon(strings.Join(args, "-"))
	}
	if err := IsCartoonSelected(); err != nil {
		return err
	}
	eps, err := selectedCartoon.Episodes()
	if err != nil {
		return err
	}
	if len(*eps) == 0 {
		return errors.New("empty result")
	}
	selectedEpisode = &(*eps)[len(*eps)-1]
	return nil
}

func LastEpisode(args []string) error {
	if len(args) > 0 {
		ResetSelection()
		selectedCartoon = NewCartoon(strings.Join(args, "-"))
	}
	if err := IsCartoonSelected(); err != nil {
		return err
	}
	eps, err := selectedCartoon.Episodes()
	if err != nil {
		return err
	}
	if len(*eps) == 0 {
		return errors.New("empty result")
	}
	selectedEpisode = &(*eps)[0]
	return nil
}
