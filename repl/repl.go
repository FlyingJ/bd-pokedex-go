package repl

import (
    "regexp"
    "strings"
)

func cleanInput(text string) []string {
    wsStr := " \t\n"
    wsRx := regexp.MustCompile(`[ \t\n]+`)

    text = strings.ToLower(text)
    text = strings.Trim(text, wsStr)
    return wsRx.Split(text, -1)
}
