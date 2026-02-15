package repl

import (
    "strings"
)

func cleanInput(rawText string) []string {
    whiteSpace := " \t\n"
    wordSeperator := " "
    return strings.Split(
        strings.Trim(
            strings.ToLower(rawText),
            whiteSpace,
        ),
        wordSeperator,
    )
}
