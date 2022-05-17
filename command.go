package main

import "strings"

type Command struct {
	Names []string
	F     func(args []string) error
}

func (c *Command) Matches(name string) bool {
	nameLower := strings.ToLower(name)
	for _, n := range c.Names {
		if strings.ToLower(n) == nameLower {
			return true
		}
	}
	return false
}
