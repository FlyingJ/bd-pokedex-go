package repl

import (
    "testing"
)

func TestCleanInput(t *testing.T) {
    cases := []struct {
        input    string
        expectation []string
    }{
        {
            input:    "hello world",
            expectation: []string{"hello", "world"},
        },
        {
            input:    "  hello world  ",
            expectation: []string{"hello", "world"},
        },
        {
            input:    "Charmander Bulbasaur PIKACHU",
            expectation: []string{"charmander", "bulbasaur", "pikachu"},
        },
    }
    for _, c := range cases {
        input := c.input
        expectation := c.expectation
        result := cleanInput(input)
        // Check the length of the actual slice against the expectation slice
        // if they don't match, use t.Errorf to print an error message
        // and fail the test
        if len(result) != len(expectation) {
            t.Errorf(
                "Error: length of %v is %d while length of %v is %d",
                result,
                len(result),
                expectation,
                len(expectation),
            )
        }
        // Check each word in the slice
        // if they don't match, use t.Errorf to print an error message
        // and fail the test
        for i := range result {
            resultWord := result[i]
            expectedWord := expectation[i]
            if resultWord != expectedWord {
                t.Errorf(
                    "Error: %s and %s do not match",
                    resultWord,
                    expectedWord,
                )
            }

        }
    }
}