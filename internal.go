package termcol

import (
	"fmt"
	"strings"
)

func parse(f *Formatter, text string) []int {
	var p []int
	chars := []rune(text)

	for i := 0; i < len(chars); i++ {
		if chars[i] == f.key {
			if i+1 < len(chars) && chars[i+1] == f.key {
				i++
			} else {
				p = append(p, i)
			}
		}
	}

	return p
}

func replace(f *Formatter, text string, del int, keys []int, colors []colorKey) string {
	chars := []rune(text)

	for i := len(keys) - 1; i >= 0; i-- {
		pos := keys[i]
		colorCode := []rune(colorValues[colors[i]])
		chars = append(chars[:pos], append(colorCode, chars[pos+del:]...)...)

		if del == 1 && pos >= 2 && chars[pos-2] == '&' && chars[pos-1] == ' ' {
			chars = append(chars[:pos-1], chars[pos:]...)
		}
	}

	for i := len(chars) - 1; i >= 0; i-- {
		if chars[i] == f.resetKey {
			if i > 0 && chars[i-1] == f.resetKey {
				i--
			} else {
				resetRunes := []rune(colorValues[Reset])
				chars = append(chars[:i], append(resetRunes, chars[i+1:]...)...)
			}
		}
	}

	if f.resetAtEnd && len(colors) != 0 &&
		colors[len(colors)-1] != Reset &&
		strings.LastIndex(string(chars), colorValues[Reset]) != strings.LastIndex(string(chars), "\033[") {
		chars = append(chars, []rune(colorValues[Reset])...)
	}

	text = string(chars)
	k := string(f.key)
	resetK := string(f.resetKey)
	text = strings.ReplaceAll(text, k+k, k)
	text = strings.ReplaceAll(text, resetK+resetK, resetK)
	if f.resetBeforeNewline {
		text = strings.ReplaceAll(text, "\n", colorValues[Reset]+"\n")
	}

	return text
}

func colorize(f *Formatter, text string, colors []colorKey) (string, error) {
	keys := parse(f, text)
	if len(keys) != len(colors) {
		return "", ParseError{fmt.Sprintf("Number of colors (%d) does not match number of keys (%d)", len(colors), len(keys)), text, -1}
	}

	for i := 0; i < len(colors); i++ {
		if colors[i] < 0 || int(colors[i]) >= len(colorValues) {
			return "", ParseError{Err: fmt.Sprintf("Invalid color code %d at index %d", colors[i], i), Text: text, Pos: keys[i] + 1}
		}
	}

	text = replace(f, text, 1, keys, colors)
	return text, nil
}

func format(f *Formatter, text string) (string, error) {
	if len(text) == 0 {
		return text, nil
	}

	keys := parse(f, text)

	if len(keys) != 0 && keys[len(keys)-1]+1 >= len(text) {
		return "", ParseError{"Invalid colorKey code at string end - colorKey code without closing key", text, len(text) - 1}
	}

	var colors []colorKey
	chars := []rune(text)
	for i := 0; i < len(keys); i++ {
		color, ok := colorKeys[chars[keys[i]+1]]
		if !ok {
			return "", ParseError{Text: text, Pos: keys[i] + 1}
		}
		colors = append(colors, color)
	}

	text = replace(f, text, 2, keys, colors)

	return text, nil
}
