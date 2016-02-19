package gokicadlib

import "strings"

func NewSExp(name string, newline bool, contents ...string) string {
	var output string
	if newline {
		output = "(" + name + "\n"
		for _, s := range contents {
			output = output + "    " + strings.Replace(s, "\n", "\n    ", -1) + "\n"
		}
		output = output + ")"
	} else {
		output = "(" + name + " "
		for _, s := range contents {
			output = output + " " + s
		}
		output = output + ")"
	}

	return output
}

func JoinStrings(s ...string) string {
	return strings.Join(s, " ")
}
