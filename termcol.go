// Package termcol provides easy terminal colorization in Go.
package termcol

import (
	"fmt"
	"io"
	"os"
	"strings"
)

// Default Formatter instance used for formatting.
var df = NewFormatter()

// Formatter configures the different formatting options for terminal output.
type Formatter struct {
	key                rune
	resetKey           rune
	resetAtEnd         bool
	resetBeforeNewline bool
	successText        string
	warningText        string
	errorText          string
	successColor       colorCode
	warningColor       colorCode
	errorColor         colorCode
}

// NewFormatter creates a new Formatter with the default settings and returns it.
func NewFormatter() *Formatter {
	return &Formatter{
		key:                '&',
		resetKey:           '§',
		resetAtEnd:         true,
		resetBeforeNewline: true,
		successText:        "Success: ",
		warningText:        "Warning: ",
		errorText:          "Error: ",
		warningColor:       Yellow,
		successColor:       Green,
		errorColor:         Red,
	}
}

// SetKey sets the key for colorization in the Formatter. (Default: '&')
func (f *Formatter) SetKey(k rune) {
	f.key = k
}

// SetResetKey sets the key for resetting colorization in the Formatter. (Default: '§')
func (f *Formatter) SetResetKey(k rune) {
	f.resetKey = k
}

// ResetAtEnd sets whether the reset key should be automatically applied at the end of the text. (Default: true)
func (f *Formatter) ResetAtEnd(b bool) {
	f.resetAtEnd = b
}

// ResetBeforeNewline sets whether the reset key should be automatically applied before newlines. (Default: true)
func (f *Formatter) ResetBeforeNewline(b bool) {
	f.resetBeforeNewline = b
}

// SetSuccessStyle sets the style for success messages in the Formatter. (Default: green "Success: ")
func (f *Formatter) SetSuccessStyle(color colorCode, text string) {
	if color < 0 || int(color) >= len(colorValues) {
		return
	}
	f.successText = text
	f.successColor = color
}

// SetWarningStyle sets the style for warning messages in the Formatter. (Default: yellow "Warning: ")
func (f *Formatter) SetWarningStyle(color colorCode, text string) {
	if color < 0 || int(color) >= len(colorValues) {
		return
	}
	f.warningText = text
	f.warningColor = color
}

// SetErrorStyle sets the style for error messages in the Formatter. // (Default: red "Error: ").
func (f *Formatter) SetErrorStyle(color colorCode, text string) {
	if color < 0 || int(color) >= len(colorValues) {
		return
	}
	f.errorText = text
	f.errorColor = color
}

/*
Sprintc returns a formatted string using the provided formatting options.
'&' is used as the formatting key, '§' resets the formatting.
Example: Sprintc("& &red-bold §text", termcol.Red, termcol.Bold) will render "red-bold" in red and bold and "text" normally.
*/
func (f *Formatter) Sprintc(text string, colors ...colorCode) string {
	text = colorize(f, text, colors)
	return text
}

// Printc formats the text using Sprintc and prints it to stdout.
func (f *Formatter) Printc(text string, colors ...colorCode) (int, error) {
	text = f.Sprintc(text, colors...)
	i, err := fmt.Print(text)
	return i, err
}

// Printlnc formats the text using Sprintc and prints it to stdout ending with a newline.
func (f *Formatter) Printlnc(text string, colors ...colorCode) (int, error) {
	text = f.Sprintc(text, colors...)
	i, err := fmt.Println(text)
	return i, err
}

// Fprintc formats the text using Sprintc and prints it to the provided io.Writer.
func (f *Formatter) Fprintc(w io.Writer, text string, colors ...colorCode) (int, error) {
	text = f.Sprintc(text, colors...)
	i, err := fmt.Fprint(w, text)
	return i, err
}

/*
Sprintf formats the text using placeholders and returns it as a string.
Placeholders are defined as '%X' for fmt placeholders, and '&X' for formatting placeholders.
For example, '&r&F%s' will format the following string, provided by the user, in red and bold.
The '§' character is used to reset the formatting.
*/
func (f *Formatter) Sprintf(text string, a ...any) string {
	text = format(f, text)
	if len(a) == 0 {
		return text
	}
	return fmt.Sprintf(text, a...)
}

// Printf formats the text using Sprintf and prints it to stdout.
func (f *Formatter) Printf(text string, a ...any) (int, error) {
	text = f.Sprintf(text, a...)
	i, err := fmt.Print(text)
	return i, err
}

// Printlnf formats the text using Sprintf and prints it to stdout ending with a newline.
func (f *Formatter) Printlnf(text string, a ...any) (int, error) {
	text = f.Sprintf(text, a...)
	i, err := fmt.Println(text)
	return i, err
}

// Fprintf formats the text using Sprintf and prints it to the provided io.Writer.
func (f *Formatter) Fprintf(w io.Writer, text string, a ...any) (int, error) {
	text = f.Sprintf(text, a...)
	i, err := fmt.Fprint(w, text)
	return i, err
}

// Successf prints the text to stdout as a success message in green ending with a newline.
func (f *Formatter) Successf(text string, a ...any) (int, error) {
	text = f.Sprintf(text, a...)
	if f.resetAtEnd && !strings.Contains(text, "\033[") {
		text = text + colorValues[Reset]
	}
	text = colorValues[f.successColor] + f.successText + text
	return fmt.Fprintln(os.Stdout, text)
}

// Warningf prints the text to stdout as a warning message in yellow ending with a newline.
func (f *Formatter) Warningf(text string, a ...any) (int, error) {
	text = f.Sprintf(text, a...)
	if f.resetAtEnd && !strings.Contains(text, "\033[") {
		text = text + colorValues[Reset]
	}
	text = colorValues[f.warningColor] + f.warningText + text
	return fmt.Fprintln(os.Stdout, text)
}

// Errorf prints the text to stdout as an error message in red ending with a newline.
func (f *Formatter) Errorf(text string, a ...any) (int, error) {
	text = f.Sprintf(text, a...)
	if f.resetAtEnd && !strings.Contains(text, "\033[") {
		text = text + colorValues[Reset]
	}
	text = colorValues[f.errorColor] + f.errorText + text
	return fmt.Fprintln(os.Stdout, text)
}

// Global functions

// Sprintc is a Wrapper for defaultFormatter.Sprintc (Further information in Formatter.Sprintc)
func Sprintc(text string, colors ...colorCode) string {
	return df.Sprintc(text, colors...)
}

// Printc is a Wrapper for defaultFormatter.Printc (Further information in Formatter.Printc)
func Printc(text string, colors ...colorCode) (int, error) {
	return df.Printc(text, colors...)
}

// Printlnc is a Wrapper for defaultFormatter.Printlnc (Further information in Formatter.Printlnc)
func Printlnc(text string, colors ...colorCode) (int, error) {
	return df.Printlnc(text, colors...)
}

// Fprintc is a Wrapper for defaultFormatter.Fprintc (Further information in Formatter.Fprintc)
func Fprintc(w io.Writer, text string, colors ...colorCode) (int, error) {
	return df.Fprintc(w, text, colors...)
}

// Sprintf is a Wrapper for defaultFormatter.Sprintf (Further information in Formatter.Sprintf)
func Sprintf(text string, a ...any) string {
	return df.Sprintf(text, a...)
}

// Printf is a Wrapper for defaultFormatter.Printf (Further information in Formatter.Printf)
func Printf(text string, a ...any) (int, error) {
	return df.Printf(text, a...)
}

// Printlnf is a Wrapper for defaultFormatter.Printlnf (Further information in Formatter.Printlnf)
func Printlnf(text string, a ...any) (int, error) {
	return df.Printlnf(text, a...)
}

// Fprintf is a Wrapper for defaultFormatter.Fprintf (Further information in Formatter.Fprintf)
func Fprintf(w io.Writer, text string, a ...any) (int, error) {
	return df.Fprintf(w, text, a...)
}

// Successf is a Wrapper for defaultFormatter.Successf (Further information in Formatter.Successf)
func Successf(text string, a ...any) (int, error) {
	return df.Successf(text, a...)
}

// Warningf is a Wrapper for defaultFormatter.Warningf (Further information in Formatter.Warningf)
func Warningf(text string, a ...any) (int, error) {
	return df.Warningf(text, a...)
}

// Errorf is a Wrapper for defaultFormatter.Errorf (Further information in Formatter.Errorf)
func Errorf(text string, a ...any) (int, error) {
	return df.Errorf(text, a...)
}

// Color returns the ANSI escape code for the given colorCode.
func Color(c colorCode) string {
	if c < 0 || int(c) >= len(colorValues) {
		return ""
	}
	return colorValues[c]
}

/*
Default returns a pointer to the default Formatter instance.
This is the instance used by the global functions.
Modifying the default Formatter will affect all global functions.
*/
func Default() *Formatter {
	return df
}
