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
		colors   []colorCode
		expected string
	}

	tests := []testSprintc{
		{"i love &go", []colorCode{Red}, "i love \033[31mgo\033[0m"},
		{"i &love &go too", []colorCode{Red, Green}, "i \033[31mlove \033[32mgo too\033[0m"},
		{"go&& is &&&an awe&some language&&", []colorCode{Red, Green}, "go& is &\033[31man awe\033[32msome language&\033[0m"},
		{"&&&it's & &easy& to learn &&&&&&& and understand&&&", []colorCode{Red, Bold, Green, Yellow, Blue, Reset}, "&\033[31mit's \033[1m\033[32measy\033[33m to learn &&&\033[34m and understand&\033[0m"},
		{"&Testing is & & &important&!", []colorCode{Red, Underline, Bold, Green, Reset}, "\033[31mTesting is \033[4m\033[1m\033[32mimportant\033[0m!"},
		{"", []colorCode{}, ""},
		{"&&&&&&", []colorCode{}, "&&&"},
		{"&Hello\n&World", []colorCode{Red, Green}, "\033[31mHello\033[0m\n\033[32mWorld\033[0m"},
		{"&Bold text", []colorCode{Bold}, "\033[1mBold text\033[0m"},
		{"&Bold &Italic§", []colorCode{Bold, Italic}, "\033[1mBold \033[3mItalic\033[0m"},
		{"&Bold § normal", []colorCode{Bold}, "\033[1mBold \033[0m normal"},
		{"&你好 &世界", []colorCode{Red, Green}, "\033[31m你好 \033[32m世界\033[0m"},
		{"&Fg &Bg &Style§ & & & &Styles combined", []colorCode{Red, GreenBg, Bold, BrightBlue, GrayBg, Italic, Bold}, "\033[31mFg \033[42mBg \033[1mStyle\033[0m \033[94m\033[100m\033[3m\033[1mStyles combined\033[0m"},
		{"§", []colorCode{}, "\033[0m"},
	}

	for _, v := range tests {
		if result := Sprintc(v.text, v.colors...); result != v.expected {
			t.Errorf("\nc(%s, %v)\ngot\n%s\nexpected\n%s", v.text, v.colors, result, v.expected)
		}
	}

	type testSprintcErr struct {
		text     string
		colors   []colorCode
		expected string
	}
	errTests := []testSprintcErr{
		{"&Hello &World", []colorCode{Red, Green, Blue}, "termcol: Number of colors (3) does not match number of keys (2)\n&Hello &World"},
		{"&only one", []colorCode{}, "termcol: Number of colors (0) does not match number of keys (1)\n&only one"},
		{"Hello World", []colorCode{Red}, "termcol: Number of colors (1) does not match number of keys (0)\nHello World"},
	}

	for _, v := range errTests {
		if res := Sprintc(v.text, v.colors...); res != v.expected {
			t.Errorf("\nc(%s, %v)\ngot\n%s\nexpected\n%s", v.text, v.colors, res, v.expected)
		}
	}
}

func TestSprintf(t *testing.T) {
	type testSprintf struct {
		text     string
		a        []any
		expected string
	}

	tests := []testSprintf{
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
		if result := Sprintf(v.text, v.a...); result != v.expected {
			t.Errorf("\nc(%s, %v)\ngot\n%s\nexpected\n%s", v.text, v.a, result, v.expected)
		}
	}

	type testSprintfErr struct {
		text     string
		a        []any
		expected string
	}
	errTests := []testSprintfErr{
		{"&Hello &World", []any{}, "[termcol: Invalid color key 'H']ello &World"},
		{"&only one %s", []any{"str"}, "[termcol: Invalid color key 'o']nly one str"},
		{"Hello World", []any{}, "Hello World"},
	}

	for _, v := range errTests {
		if res := Sprintf(v.text, v.a...); res != v.expected {
			t.Errorf("c(%s, %v)\ngot\n%s\nexpected\n%s", v.text, v.a, res, v.expected)
		}
	}

	f := NewFormatter()
	f.SetKey('N')
	f.SetResetKey('n')
	f.ResetAtEnd(false)
	f.ResetBeforeNewline(false)

	tests = []testSprintf{
		{"i love Nr%s", []any{"go"}, "i love \033[31mgo"},
		{"i Nr%s Ng%s %vn", []any{"love", "go", "too"}, "i \033[31mlove \033[32mgo too\033[0m"},
		{"NrHello\nNgWorld", []any{}, "\033[31mHello\n\033[32mWorld"},
		{"Nr你好 &g世界", []any{}, "\033[31m你好 &g世界"},
	}

	for _, v := range tests {
		if result := f.Sprintf(v.text, v.a...); result != v.expected {
			t.Errorf("\nc(%s, %v)\ngot\n%s\nexpected\n%s", v.text, v.a, result, v.expected)
		}
	}
}
