package main

import (
	"strings"
)

func splitBySpace(s string) []string {
	return strings.Split(s, " ")
}

func splitByComma(s string) []string {
	return strings.Split(s, ",")
}

