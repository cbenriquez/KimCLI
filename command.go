package main

type Command struct {
	Names []string
	F     func(cp *Prompt, args []string) error
}

func (c Command) Matches(name string) bool {
	for _, n := range c.Names {
		if n == name {
			return true
		}
	}
	return false
}
