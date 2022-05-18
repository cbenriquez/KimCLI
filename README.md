# KimCLI

A command line interface for watching cartoons.

It scraps information from [KimCartoon](https://kimcartoon.li/) hence the name. It uses MPV to play videos. It is developed in Go.

## Commands

When the program is executed, you are presented with a command prompt. Use these commands to navigate through the program.

### Selecting a Cartoon

| Name(s) | Parameter(s) | Description
| - | - | -
| search, s, find, f | [keywords] | Select a cartoon from a search query.

### Selecting an Episode

| Name(s) | Parameter(s) | Description
| - | - | -
| episodes, eps | *none* | Select an episode from a list.
| first-episode, fstep | *none* | Select the first episode.
| last-episode, lstep | *none* | Select the last peisode.
| next, n | *none* | Select the next episode.
| back, b | *none* | Select the previous episode.

## Playing an Episode

| Name(s) | Parameter(s) | Description
| - | - | -
| watch, w, play, p | *none* | Select a playback quality from a list and play the episode.
| play-highest, ph | *none* | Play the episode with the highest playback quality.

## Downloading an Episode

| Name(s) | Parameter(s) | Description
| - | - | -
| download, d | *none* | Select a playback quality and download the episode.
| download-highest, dh | *none* | Download the episode with the highest playback quality.

## Others

| Name(s) | Parameter(s) | Description
| - | - | -
| exit, e, quit, q, ! | *none* | Terminate the program.
| list-episodes, le | *none* | Print a list of episodes.

## Installation

Move the binary to any PATH directory.

## Dependencies
The only supported video player is MPV though more support will be added later.

- [mpv](https://mpv.io/)