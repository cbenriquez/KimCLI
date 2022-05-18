package main

import (
	"errors"
	"io"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
)

type Video struct {
	File    string `json:"file"`
	Label   string `json:"label"`
	Type    string `json:"type"`
	Episode *Episode
}

func (v *Video) Title() (*string, error) {
	ct, err := v.Episode.Cartoon.Title()
	if err != nil {
		return nil, err
	}
	title := *ct + " " + v.Episode.Name
	return &title, nil
}

func (v *Video) Play() error {
	title, err := v.Title()
	if err != nil {
		return err
	}
	var cmd *exec.Cmd
	if p, err := exec.LookPath("mpv"); err == nil {
		cmd = &exec.Cmd{
			Path: p,
			Args: []string{"mpv", "--title=" + *title, v.File},
		}
	} else if p, err := exec.LookPath("vlc"); err == nil {
		cmd = &exec.Cmd{
			Path: p,
			Args: []string{"vlc", "--meta-title=" + *title, v.File},
		}
	} else {
		return errors.New("cannot find a supported media player")
	}
	if err := cmd.Run(); err != nil {
		return err
	}
	return nil
}

func (v *Video) Download() (*string, error) {
	resp, err := http.Get(v.File)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	hp, err := os.UserHomeDir()
	if err != nil {
		return nil, err
	}
	df := filepath.Join(hp, "Downloads")
	if stat, err := os.Stat(df); err != nil {
		return nil, err
	} else if !stat.IsDir() {
		return nil, errors.New("cannot access downloads")
	}
	title, err := v.Title()
	if err != nil {
		return nil, err
	}
	out, err := os.Create(filepath.Join(df, *title+"."+v.Type))
	if err != nil {
		return nil, err
	}
	defer out.Close()
	if _, err := io.Copy(out, resp.Body); err != nil {
		return nil, err
	}
	return &df, nil
}
