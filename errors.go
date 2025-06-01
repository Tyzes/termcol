package termcol

import (
	"fmt"
	"strings"
)

type ParseError struct {
	Err  string
	Text string
	Pos  int
}

func (e ParseError) Error() string {
	if e.Pos < 0 {
		return "termcol parse error: " + e.Err
	} else if e.Pos != 0 && e.Pos >= len(e.Text) {
		return "termcol parse error: Illegal character at end of text"
	} else if e.Text == "" {
		return fmt.Sprintf("termcol parse error: %s", e.Err)
	}

	start := e.Pos - 15
	if start < 0 {
		start = 0
	}
	end := e.Pos + 16
	if end > len(e.Text) {
		end = len(e.Text)
	}

	text := e.Text[start:end]
	caretPos := e.Pos - start
	caretPos += strings.Count(text[:caretPos], "\n")
	text = strings.ReplaceAll(text, "\n", `\n`)

	if start > 0 {
		text = "..." + text
		caretPos += 3
	}
	if end < len(e.Text) {
		text = text + "..."
	}

	if e.Err == "" {
		return fmt.Sprintf(
			"termcol parse error: Illegal character %q at index %d\n%s\n%s^ HERE",
			e.Text[e.Pos], e.Pos, text, strings.Repeat(" ", caretPos),
		)
	} else {
		return fmt.Sprintf(
			"termcol parse error: %s\n%s\n%s^ HERE",
			e.Err, text, strings.Repeat(" ", caretPos),
		)
	}
}
