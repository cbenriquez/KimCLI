package main

import (
	"errors"
	"strings"
)

var selectedCartoon *Cartoon

var selectedEpisode *Episode

func ResetSelection() {
	selectedCartoon = nil
	selectedEpisode = nil
}

func IsCartoonSelected() error {
	if selectedCartoon == nil {
		return errors.New("cartoon not selected")
	}
	return nil
}

func IsEpisodeSelected() error {
	if selectedEpisode == nil {
		return errors.New("episode not selected")
	}
	return nil
}

func SelectEpisode(id string) error {
	eps, err := selectedCartoon.Episodes()
	if err != nil {
		return err
	}
	es := strings.ToLower(id)
	for _, ep := range *eps {
		if strings.ToLower(ep.ID) == es {
			selectedEpisode = &ep
			break
		}
	}
	if selectedEpisode == nil {
		return errors.New("invalid episode ID")
	}
	return nil
}
