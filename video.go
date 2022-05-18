package main

import (
	"errors"
	"os/exec"
)

type Video struct {
	File    string `json:"file"`
	Label   string `json:"label"`
	Type    string `json:"type"`
	Episode *Episode
}

func (v *Video) Play() error {
	ct, err := v.Episode.Cartoon.Title()
	if err != nil {
		return err
	}
	var cmd *exec.Cmd
	title := *ct + " " + v.Episode.Name
	if p, err := exec.LookPath("mpv"); err == nil {
		cmd = &exec.Cmd{
			Path: p,
			Args: []string{"mpv", "--title=" + title, v.File},
		}
	} else if p, err := exec.LookPath("vlc"); err == nil {
		cmd = &exec.Cmd{
			Path: p,
			Args: []string{"vlc", "--meta-title=" + title, v.File},
		}
	} else {
		return errors.New("cannot find a supported media player")
	}
	if err := cmd.Run(); err != nil {
		return err
	}
	return nil
}
