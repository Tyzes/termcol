package main

import (
	"github.com/tyzes/termcol"
	"os"
)

func main() {
	writer := os.Stdout

	// Print formatted text directly to stdout
	_, _ = termcol.Printc("&Red ", termcol.Red)

	// Return formatted text as a string
	str, _ := termcol.Sprintc("&Green ", termcol.Green)

	// Print formatted text to a specified io.Writer
	_, _ = termcol.Fprintc(writer, "&Yellow ", termcol.Yellow)

	// Print formatted text ending with a newline
	_, _ = termcol.Printlnc("&Blue ", termcol.Blue)

	// Print formatted text with arguments
	_, _ = termcol.Printf("&m%s ", "Magenta")

	// Return formatted text as a string with arguments
	str, _ = termcol.Sprintf("&c%s ", "Cyan")

	// Print formatted text to a specific writer
	_, _ = termcol.Fprintf(writer, "&F%s ", "Bold")

	// Print formatted text ending with a newline
	_, _ = termcol.Printlnf("&r%s ", "Red")

	// Using predefined helper functions for common log types:
	str = "something happened"

	// Success message (default: green + "Success: ")
	_, _ = termcol.Successf("%s", str)

	// Warning message (default: yellow + "Warning: ")
	_, _ = termcol.Warningf("%s", str)

	// Error message (default: red + "Error: ")
	_, _ = termcol.Errorf("%v", str)
}
