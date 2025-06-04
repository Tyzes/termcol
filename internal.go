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

func replace(f *Formatter, text string, del int, keys []int, colors []colorCode) string {
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

func colorize(f *Formatter, text string, colors []colorCode) string {
	keys := parse(f, text)
	if len(keys) != len(colors) {
		b := strings.Builder{}
		b.WriteString("termcol: Number of colors (")
		b.WriteString(fmt.Sprint(len(colors)))
		b.WriteString(") does not match number of keys (")
		b.WriteString(fmt.Sprint(len(keys)))
		b.WriteRune(')')
		b.WriteRune('\n')
		b.WriteString(text)
		return b.String()
	}

	for i := 0; i < len(colors); i++ {
		if colors[i] < 0 || int(colors[i]) >= len(colorValues) {
			b := strings.Builder{}
			b.WriteString("termcol: Invalid color code ")
			b.WriteString(fmt.Sprint(colors[i]))
			b.WriteString(" as argument ")
			b.WriteString(fmt.Sprint(i + 2))
			b.WriteRune('\n')
			b.WriteString(text)
			return b.String()
		}
	}

	text = replace(f, text, 1, keys, colors)
	return text
}

func format(f *Formatter, text string) string {
	if len(text) == 0 {
		return text
	}

	keys := parse(f, text)

	if len(keys) != 0 && keys[len(keys)-1]+1 >= len(text) {
		return text[:len(text)-2] + "[termcol: Color key without value at end of text]"
	}

	var colors []colorCode
	chars := []rune(text)
	for _, key := range keys {
		color, ok := colorKeys[chars[key+1]]
		if !ok {
			b := strings.Builder{}
			b.WriteString(text[:key])
			b.WriteString("[termcol: Invalid color key '")
			b.WriteRune(chars[key+1])
			b.WriteString("']")
			b.WriteString(text[key+2:])
			return b.String()
		}
		colors = append(colors, color)
	}

	text = replace(f, text, 2, keys, colors)

	return text
}

func isColorCode(c colorCode) bool {
	return c >= 0 && int(c) < len(colorValues)
}
