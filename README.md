# KimCLI

A command line interface for watching cartoons.

It scraps information from [KimCartoon](https://kimcartoon.li/) hence the name. It uses MPV to play videos. The program is developed in Go and distributed for Linux.

Builds for Windows and MacOS will come in later versions.

## Commands

When the program is executed, you are presented with a command prompt. Use these commands to navigate through the program.

| Name(s) | Parameter(s) | Description
| - | - | -
| exit, e, quit, q | *none* | Terminate the program.
| search, s, find, f | [keywords] | Select a cartoon.
| episodes, eps | *optional:* [cartoon] | Select an episode. If parameter is empty, use the current selection.
| watch, w, play, p | *optional:* [episode] | Select a playback quality and play the episode.


## Installation

Move the binary to any PATH directory.

## Dependencies
The only supported video player is MPV though more support will be added later.

- [mpv](https://mpv.io/)

## Planned Features
- download video
- play next video
- autoselect playback quality
- update mechanism