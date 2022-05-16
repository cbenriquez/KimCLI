package main

import (
	"fmt"
	"os/exec"
)

type Video struct {
	File    string
	Label   string
	Type    string
	Episode *Episode
}

func (v *Video) Play() error {
	ct, err := v.Episode.Cartoon.Title()
	if err != nil {
		return err
	}
	cmd := exec.Command("mpv", fmt.Sprintf(`--title=%s %s`, *ct, v.Episode.Name), v.File)
	if err := cmd.Run(); err != nil {
		return err
	}
	return nil
}
