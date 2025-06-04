package termcol

type colorCode int

const (
	Reset colorCode = iota // Reset Formatting, §

	Black         // Color Black, &s
	Red           // Color Red, &r
	Green         // Color Green, &g
	Yellow        // Color Yellow, &y
	Blue          // Color Blue, &b
	Magenta       // Color Magenta, &m
	Cyan          // Color Cyan, &colorize
	White         // Color White, &w
	Gray          // Color Gray, &a
	BrightRed     // Color Bright Red, &R
	BrightGreen   // Color Bright Green, &G
	BrightYellow  // Color Bright Yellow, &Y
	BrightBlue    // Color Bright Blue, &B
	BrightMagenta // Color Bright Magenta, &M
	BrightCyan    // Color Bright Cyan, &C
	BrightWhite   // Color Bright White, &W

	Bold          // Text Bold, &F
	Italic        // Text Italic, &I
	Underline     // Text Underline, &U
	StrikeThrough // Text Strike Through, &S

	BlackBg         // Background Color Black
	RedBg           // Background Color Red
	GreenBg         // Background Color Green
	YellowBg        // Background Color Yellow
	BlueBg          // Background Color Blue
	MagentaBg       // Background Color Magenta
	CyanBg          // Background Color Cyan
	WhiteBg         // Background Color White
	GrayBg          // Background Color Gray
	BrightRedBg     // Background Color Bright Red
	BrightGreenBg   // Background Color Bright Green
	BrightYellowBg  // Background Color Bright Yellow
	BrightBlueBg    // Background Color Bright Blue
	BrightMagentaBg // Background Color Bright Magenta
	BrightCyanBg    // Background Color Bright Cyan
	BrightWhiteBg   // Background Color Bright White
)

var colorValues = []string{
	"\033[0m", // 0: Reset

	"\033[30m", // 1: Black
	"\033[31m", // 2: Red
	"\033[32m", // 3: Green
	"\033[33m", // 4: Yellow
	"\033[34m", // 5: Blue
	"\033[35m", // 6: Magenta
	"\033[36m", // 7: Cyan
	"\033[37m", // 8: White

	"\033[90m", // 9: Gray
	"\033[91m", // 10: BrightRed
	"\033[92m", // 11: BrightGreen
	"\033[93m", // 12: BrightYellow
	"\033[94m", // 13: BrightBlue
	"\033[95m", // 14: BrightMagenta
	"\033[96m", // 15: BrightCyan
	"\033[97m", // 16: BrightWhite

	"\033[1m", // 17: Bold
	"\033[3m", // 18: Italic
	"\033[4m", // 19: Underline
	"\033[9m", // 20: StrikeThrough

	"\033[40m", // 21: BlackBg
	"\033[41m", // 22: RedBg
	"\033[42m", // 23: GreenBg
	"\033[43m", // 24: YellowBg
	"\033[44m", // 25: BlueBg
	"\033[45m", // 26: MagentaBg
	"\033[46m", // 27: CyanBg
	"\033[47m", // 28: WhiteBg

	"\033[100m", // 29: GrayBg
	"\033[101m", // 30: BrightRedBg
	"\033[102m", // 31: BrightGreenBg
	"\033[103m", // 32: BrightYellowBg
	"\033[104m", // 33: BrightBlueBg
	"\033[105m", // 34: BrightMagentaBg
	"\033[106m", // 35: BrightCyanBg
	"\033[107m", // 36: BrightWhiteBg
}

// Mapping colorCode keys to colorCode values
var colorKeys = map[rune]colorCode{
	's': Black,         // &s
	'r': Red,           // &r
	'g': Green,         // &g
	'y': Yellow,        // &y
	'b': Blue,          // &b
	'm': Magenta,       // &m
	'c': Cyan,          // &c
	'w': White,         // &w
	'a': Gray,          // &a
	'R': BrightRed,     // &R
	'G': BrightGreen,   // &G
	'Y': BrightYellow,  // &Y
	'B': BrightBlue,    // &B
	'M': BrightMagenta, // &M
	'C': BrightCyan,    // &C
	'W': BrightWhite,   // &W

	'F': Bold,          // &F
	'I': Italic,        // &I
	'U': Underline,     // &U
	'S': StrikeThrough, // &S
}
