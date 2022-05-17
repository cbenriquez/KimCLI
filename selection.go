package main

import "errors"

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
