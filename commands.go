package main

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

var commands []Command = []Command{
	// Selecting a Cartoon
	{
		Names: []string{"search", "s", "find", "f"},
		F:     CommandSearch,
	},
	// Selecting an Episode
	{
		Names: []string{"episodes", "eps"},
		F:     CommandEpisodes,
	},
	{
		Names: []string{"first-episode", "fstep"},
		F:     CommandFirstEpisode,
	},
	{
		Names: []string{"last-episode", "lstep"},
		F:     CommandLastEpisode,
	},
	{
		Names: []string{"next", "n"},
		F:     CommandNext,
	},
	{
		Names: []string{"back", "b"},
		F:     CommandBack,
	},
	// Playing an Episode
	{
		Names: []string{"play", "p", "watch", "w"},
		F:     CommandPlay,
	},
	{
		Names: []string{"play-highest", "ph", "watch-highest", "wh"},
		F:     CommandPlayHighest,
	},
	// Downloading an Episode
	{
		Names: []string{"download", "d"},
		F:     CommandDownload,
	},
	{
		Names: []string{"download-highest", "dh"},
		F:     CommandDownloadHighest,
	},
	// Others
	{
		Names: []string{"exit", "e", "quit", "q", "!"},
		F:     CommandExit,
	},
	{
		Names: []string{"list-episodes", "le"},
		F:     CommandListEpisodes,
	},
}

// Selecting a Cartoon

func CommandSearch(args []string) error {
	carts, err := SearchCartoons(strings.Join(args, " "))
	if err != nil {
		return err
	}
	if len(*carts) == 0 {
		return errors.New("empty result")
	}
	var ci int
	if len(*carts) > 1 {
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
		ci = *c
	}
	ResetSelection()
	selectedCartoon = &(*carts)[ci]
	return nil
}

// Selecting an Episode

func CommandEpisodes(args []string) error {
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

func CommandFirstEpisode(args []string) error {
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

func CommandLastEpisode(args []string) error {
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

func CommandNext(_ []string) error {
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

func CommandBack(_ []string) error {
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

// Playing an Episode

func CommandPlay(args []string) error {
	if err := IsEpisodeSelected(); err != nil {
		return err
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

func CommandPlayHighest(args []string) error {
	if err := IsEpisodeSelected(); err != nil {
		return err
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

// Downloading an Episode

func CommandDownload(args []string) error {
	if err := IsEpisodeSelected(); err != nil {
		return err
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
	fmt.Println("download started")
	p, err := vid.Download()
	if err != nil {
		return err
	}
	fmt.Println("downloaded to", *p)
	return nil
}

func CommandDownloadHighest(args []string) error {
	if err := IsEpisodeSelected(); err != nil {
		return err
	}
	vids, err := selectedEpisode.Videos()
	if err != nil {
		return err
	}
	if len(*vids) == 0 {
		return errors.New("no videos found")
	}
	vid := (*vids)[len(*vids)-1]
	fmt.Println("download started")
	p, err := vid.Download()
	if err != nil {
		return err
	}
	fmt.Println("downloaded to", *p)
	return nil
}

// Others

func CommandExit(args []string) error {
	os.Exit(0)
	return nil
}

func CommandListEpisodes(args []string) error {
	if err := IsCartoonSelected(); err != nil {
		return err
	}
	eps, err := selectedCartoon.Episodes()
	if err != nil {
		return err
	}
	for i, ep := range *eps {
		fmt.Printf("%d) %s\n", i, ep.Name)
	}
	return nil
}
