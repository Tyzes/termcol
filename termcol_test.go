package termcol

import (
	"testing"
)

func TestParseError(t *testing.T) {
	type testParseError struct {
		text     string
		pos      int
		expected string
	}

	tests := []testParseError{
		{"Testing is the most important thing you can \n\ndo \n&xwhich is why you need to test absolutely everything", 51,
			"termcol parse error: Illegal character 'x' at index 51\n...you can \\n\\ndo \\n&xwhich is why yo...\n                     ^ HERE"},
		{"It is very important, to test your code &xproperly", 41,
			"termcol parse error: Illegal character 'x' at index 41\n...est your code &xproperly\n                  ^ HERE"},
		{"You &xsee, testing is great, as you find bugs you didn't know existed", 5,
			"termcol parse error: Illegal character 'x' at index 5\nYou &xsee, testing is...\n     ^ HERE"},
		{"And can &xfix them", 9, "termcol parse error: Illegal character 'x' at index 9\nAnd can &xfix them\n         ^ HERE"},
	}

	for _, v := range tests {
		err := &ParseError{Text: v.text, Pos: v.pos}
		if err.Error() != v.expected {
			t.Errorf("ParseError(%q, %d)\n%s\nwant\n%s", v.text, v.pos, err.Error(), v.expected)
		}
	}
}

func TestSprintc(t *testing.T) {
	type testSprintc struct {
		text     string
		colors   []colorKey
		expected string
	}

	tests := []testSprintc{
		{"i love &go", []colorKey{Red}, "i love \033[31mgo\033[0m"},
		{"i &love &go too", []colorKey{Red, Green}, "i \033[31mlove \033[32mgo too\033[0m"},
		{"go&& is &&&an awe&some language&&", []colorKey{Red, Green}, "go& is &\033[31man awe\033[32msome language&\033[0m"},
		{"&&&it's & &easy& to learn &&&&&&& and understand&&&", []colorKey{Red, Bold, Green, Yellow, Blue, Reset}, "&\033[31mit's \033[1m\033[32measy\033[33m to learn &&&\033[34m and understand&\033[0m"},
		{"&Testing is & & &important&!", []colorKey{Red, Underline, Bold, Green, Reset}, "\033[31mTesting is \033[4m\033[1m\033[32mimportant\033[0m!"},
		{"", []colorKey{}, ""},
		{"&&&&&&", []colorKey{}, "&&&"},
		{"&Hello\n&World", []colorKey{Red, Green}, "\033[31mHello\033[0m\n\033[32mWorld\033[0m"},
		{"&Bold text", []colorKey{Bold}, "\033[1mBold text\033[0m"},
		{"&Bold &Italic§", []colorKey{Bold, Italic}, "\033[1mBold \033[3mItalic\033[0m"},
		{"&Bold § normal", []colorKey{Bold}, "\033[1mBold \033[0m normal"},
		{"&你好 &世界", []colorKey{Red, Green}, "\033[31m你好 \033[32m世界\033[0m"},
		{"&Fg &Bg &Style§ & & & &Styles combined", []colorKey{Red, GreenBg, Bold, BrightBlue, GrayBg, Italic, Bold}, "\033[31mFg \033[42mBg \033[1mStyle\033[0m \033[94m\033[100m\033[3m\033[1mStyles combined\033[0m"},
		{"§", []colorKey{}, "\033[0m"},
	}

	for _, v := range tests {
		result, err := Sprintc(v.text, v.colors...)
		if err != nil && err.Error() != v.expected {
			t.Errorf("\nc(%s, %v)\ngot error\n%s\nexpected\n%s", v.text, v.colors, err, v.expected)
		}
		if err == nil && result != v.expected {
			t.Errorf("\nc(%s, %v)\ngot\n%s\nexpected\n%s", v.text, v.colors, result, v.expected)
		}
	}

	type testSprintcErr struct {
		text     string
		colors   []colorKey
		expected string
	}
	errTests := []testSprintcErr{
		{"&Hello &World", []colorKey{Red, Green, Blue}, "termcol parse error: Number of colors (3) does not match number of keys (2)\n&Hello &World"},
		{"&only one", []colorKey{}, "termcol parse error: Number of colors (0) does not match number of keys (1)\n&only one"},
		{"Hello World", []colorKey{Red}, "termcol parse error: Number of colors (1) does not match number of keys (0)\nHello World"},
	}

	for _, v := range errTests {
		if _, err := Sprintc(v.text, v.colors...); err == nil || err.Error() != v.expected {
			t.Errorf("\nc(%s, %v)\ngot\n%s\nexpected\n%s", v.text, v.colors, err, v.expected)
		}
	}
}

func TestSprintf(t *testing.T) {
	type testSprintc struct {
		text     string
		a        []any
		expected string
	}

	tests := []testSprintc{
		{"i love &r%s", []any{"go"}, "i love \033[31mgo\033[0m"},
		{"i &r%s &g%s %v", []any{"love", "go", "too"}, "i \033[31mlove \033[32mgo too\033[0m"},
		{"go&& is &&&ran awe&gsome language&&", []any{}, "go& is &\033[31man awe\033[32msome language&\033[0m"},
		{"&&&rit's &F&geasy&y to learn &&&&&&&b and understand&&§", []any{}, "&\033[31mit's \033[1m\033[32measy\033[33m to learn &&&\033[34m and understand&\033[0m"},
		{"&rTesting is &U&F&gimportant§!", []any{}, "\033[31mTesting is \033[4m\033[1m\033[32mimportant\033[0m!"},
		{"", []any{}, ""},
		{"&&&&&&", []any{}, "&&&"},
		{"&rHello\n&gWorld", []any{}, "\033[31mHello\033[0m\n\033[32mWorld\033[0m"},
		{"&FBold text", []any{}, "\033[1mBold text\033[0m"},
		{"&FBold &IItalic§", []any{}, "\033[1mBold \033[3mItalic\033[0m"},
		{"&FBold §normal", []any{}, "\033[1mBold \033[0mnormal"},
		{"&r你好 &g世界", []any{}, "\033[31m你好 \033[32m世界\033[0m"},
		{"&rFg &FStyle§ &B&I&FStyles combined", []any{}, "\033[31mFg \033[1mStyle\033[0m \033[94m\033[3m\033[1mStyles combined\033[0m"},
		{"§", []any{}, "\033[0m"},
	}

	for _, v := range tests {
		result, err := Sprintf(v.text, v.a...)
		if err != nil && err.Error() != v.expected {
			t.Errorf("\nc(%s, %v)\ngot error\n%s\nexpected\n%s", v.text, v.a, err, v.expected)
		}
		if err == nil && result != v.expected {
			t.Errorf("\nc(%s, %v)\ngot\n%s\nexpected\n%s", v.text, v.a, result, v.expected)
		}
	}

	type testSprintcErr struct {
		text     string
		a        []any
		expected string
	}
	errTests := []testSprintcErr{
		{"&Hello &World", []any{}, "termcol parse error: Illegal character 'H' at index 1\n&Hello &World\n ^ HERE"},
		{"&only one %s", []any{"str"}, "termcol parse error: Illegal character 'o' at index 1\n&only one %s\n ^ HERE"},
		{"Hello World", []any{}, ""},
	}

	for _, v := range errTests {
		if _, err := Sprintf(v.text, v.a...); err == nil || err.Error() != v.expected {
			if err == nil && v.expected != "" {
				t.Errorf("\nc(%s, %v)\ngot nil error\nexpected\n%s", v.text, v.a, v.expected)
			} else if err != nil {
				t.Errorf("\nc(%s, %v)\ngot\n%s\nexpected\n%s", v.text, v.a, err, v.expected)
			}
		}
	}

	f := NewFormatter()
	f.SetKey('N')
	f.SetResetKey('n')
	f.ResetAtEnd(false)
	f.ResetBeforeNewline(false)

	tests = []testSprintc{
		{"i love Nr%s", []any{"go"}, "i love \033[31mgo"},
		{"i Nr%s Ng%s %vn", []any{"love", "go", "too"}, "i \033[31mlove \033[32mgo too\033[0m"},
		{"NrHello\nNgWorld", []any{}, "\033[31mHello\n\033[32mWorld"},
	}

	for _, v := range tests {
		result, err := f.Sprintf(v.text, v.a...)
		if err != nil && err.Error() != v.expected {
			t.Errorf("\nc(%s, %v)\ngot error\n%s\nexpected\n%s", v.text, v.a, err, v.expected)
		}
		if err == nil && result != v.expected {
			t.Errorf("\nc(%s, %v)\ngot\n%s\nexpected\n%s", v.text, v.a, result, v.expected)
		}
	}
}
