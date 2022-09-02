package golor

import (
	"strconv"
	"strings"
)

type GolorMode int

const (
	Plain GolorMode = iota
	Bit8
	Bit256
	BitRGB
)

type GolorSound int

const (
	Active GolorSound = iota
	Mute
)

var Mode GolorMode
var Sound GolorSound

type GolorGhapMode string

const (
	Reset           GolorGhapMode = "\x1b[0m"
	Bold            GolorGhapMode = "\x1b[1m"
	Faint           GolorGhapMode = "\x1b[2m"
	Italic          GolorGhapMode = "\x1b[3m"
	Underline       GolorGhapMode = "\x1b[4m"
	Blinking        GolorGhapMode = "\x1b[5m"
	Reverse         GolorGhapMode = "\x1b[7m"
	Hidden          GolorGhapMode = "\x1b[8m"
	Strikerthrough  GolorGhapMode = "\x1b[9m"
	Boldr           GolorGhapMode = "\x1b[22m"
	Faintr          GolorGhapMode = "\x1b[22m"
	Italicr         GolorGhapMode = "\x1b[23m"
	Underliner      GolorGhapMode = "\x1b[24m"
	Blinkingr       GolorGhapMode = "\x1b[25m"
	Reverser        GolorGhapMode = "\x1b[27m"
	Hiddenr         GolorGhapMode = "\x1b[28m"
	Strikerthroughr GolorGhapMode = "\x1b[29m"
)

type GolorColor string

const (
	Black       GolorColor = "\x1b[30m"
	Red         GolorColor = "\x1b[31m"
	Green       GolorColor = "\x1b[32m"
	Yellow      GolorColor = "\x1b[33m"
	Blue        GolorColor = "\x1b[34m"
	Magenta     GolorColor = "\x1b[35m"
	Cyan        GolorColor = "\x1b[36m"
	White       GolorColor = "\x1b[37m"
	Default     GolorColor = "\x1b[39m"
	BlackBack   GolorColor = "\x1b[40m"
	RedBack     GolorColor = "\x1b[41m"
	GreenBack   GolorColor = "\x1b[42m"
	YellowBack  GolorColor = "\x1b[43m"
	BlueBack    GolorColor = "\x1b[44m"
	MagentaBack GolorColor = "\x1b[45m"
	CyanBack    GolorColor = "\x1b[46m"
	WhiteBack   GolorColor = "\x1b[47m"
	DefaultBack GolorColor = "\x1b[49m"
)

type Golor struct {
	len int
	sb  strings.Builder
}

func SetMode(s string) {
	switch s {
	case "bit8":
		Mode = Bit8
	case "bit256":
		Mode = Bit256
	case "bitrgb":
		Mode = BitRGB
	default:
		Mode = Plain
	}
}

func New() Golor {
	g := Golor{
		len: 0,
	}
	return g
}

func (o *Golor) Len() int {
	return o.len
}

func (o *Golor) String() string {
	return o.sb.String()
}

func (o *Golor) Add(s string) *Golor {
	o.len += len(s)
	o.sb.WriteString(s)
	return o
}

func (o *Golor) Reset() *Golor {
	if Mode != Plain {
		o.sb.WriteString(string(Reset))
	}
	return o
}

func (o *Golor) ResetAdd(s string) *Golor {
	o.Reset()
	o.Add(s)
	return o
}

func (o *Golor) Bell() *Golor {
	if Mode != Plain && Sound != Mute {
		o.sb.WriteString("\a")
	}
	return o
}

func (o *Golor) BellAdd(s string) *Golor {
	o.Bell()
	o.Add(s)
	return o
}

func (o *Golor) Bold() *Golor {
	if Mode != Plain {
		o.sb.WriteString(string(Bold))
	}
	return o
}

func (o *Golor) BoldAdd(s string) *Golor {
	o.Bold()
	o.Add(s)
	return o
}

func (o *Golor) BoldReset() *Golor {
	if Mode != Plain {
		o.sb.WriteString(string(Boldr))
	}
	return o
}

func (o *Golor) BoldResetAdd(s string) *Golor {
	o.BoldReset()
	o.Add(s)
	return o
}

func (o *Golor) Faint() *Golor {
	if Mode != Plain {
		o.sb.WriteString(string(Faint))
	}
	return o
}

func (o *Golor) FaintAdd(s string) *Golor {
	o.Faint()
	o.Add(s)
	return o
}

func (o *Golor) FaintReset() *Golor {
	if Mode != Plain {
		o.sb.WriteString(string(Faintr))
	}
	return o
}

func (o *Golor) FaintResetAdd(s string) *Golor {
	o.FaintReset()
	o.Add(s)
	return o
}

func (o *Golor) Italic() *Golor {
	if Mode != Plain {
		o.sb.WriteString(string(Italic))
	}
	return o
}

func (o *Golor) ItalicAdd(s string) *Golor {
	o.Italic()
	o.Add(s)
	return o
}

func (o *Golor) ItalicReset() *Golor {
	if Mode != Plain {
		o.sb.WriteString(string(Italicr))
	}
	return o
}

func (o *Golor) ItalicResetAdd(s string) *Golor {
	o.ItalicReset()
	o.Add(s)
	return o
}

func (o *Golor) Underline() *Golor {
	if Mode != Plain {
		o.sb.WriteString(string(Underline))
	}
	return o
}

func (o *Golor) UnderlineAdd(s string) *Golor {
	o.Underline()
	o.Add(s)
	return o
}

func (o *Golor) UnderlineReset() *Golor {
	if Mode != Plain {
		o.sb.WriteString(string(Underliner))
	}
	return o
}

func (o *Golor) UnderlineResetAdd(s string) *Golor {
	o.UnderlineReset()
	o.Add(s)
	return o
}

func (o *Golor) Blinking() *Golor {
	if Mode != Plain {
		o.sb.WriteString(string(Blinking))
	}
	return o
}

func (o *Golor) BlinkingAdd(s string) *Golor {
	o.Blinking()
	o.Add(s)
	return o
}

func (o *Golor) BlinkingReset() *Golor {
	if Mode != Plain {
		o.sb.WriteString(string(Blinkingr))
	}
	return o
}

func (o *Golor) BlinkingResetAdd(s string) *Golor {
	o.BlinkingReset()
	o.Add(s)
	return o
}

func (o *Golor) Reverse() *Golor {
	if Mode != Plain {
		o.sb.WriteString(string(Reverse))
	}
	return o
}

func (o *Golor) ReverseAdd(s string) *Golor {
	o.Reverse()
	o.Add(s)
	return o
}

func (o *Golor) ReverseReset() *Golor {
	if Mode != Plain {
		o.sb.WriteString(string(Reverser))
	}
	return o
}

func (o *Golor) ReverseResetAdd(s string) *Golor {
	o.ReverseReset()
	o.Add(s)
	return o
}

func (o *Golor) Hidden() *Golor {
	if Mode != Plain {
		o.sb.WriteString(string(Hidden))
	}
	return o
}

func (o *Golor) HiddenAdd(s string) *Golor {
	o.Hidden()
	o.Add(s)
	return o
}

func (o *Golor) HiddenReset() *Golor {
	if Mode != Plain {
		o.sb.WriteString(string(Hiddenr))
	}
	return o
}

func (o *Golor) HiddenResetAdd(s string) *Golor {
	o.HiddenReset()
	o.Add(s)
	return o
}

func (o *Golor) Strikerthrough() *Golor {
	if Mode != Plain {
		o.sb.WriteString(string(Strikerthrough))
	}
	return o
}

func (o *Golor) StrikerthroughAdd(s string) *Golor {
	o.Strikerthrough()
	o.Add(s)
	return o
}

func (o *Golor) StrikerthroughReset() *Golor {
	if Mode != Plain {
		o.sb.WriteString(string(Strikerthroughr))
	}
	return o
}

func (o *Golor) StrikerthroughResetAdd(s string) *Golor {
	o.StrikerthroughReset()
	o.Add(s)
	return o
}

func (o *Golor) ColorBit8(c GolorColor) *Golor {
	if Mode == Bit8 {
		o.sb.WriteString(string(c))
	}
	return o
}

func (o *Golor) ColorBit8Add(s string, c GolorColor) *Golor {
	o.ColorBit8(c)
	o.Add(s)
	return o
}

func (o *Golor) Black() *Golor {
	o.ColorBit8(Black)
	return o
}

func (o *Golor) BlackAdd(s string) *Golor {
	o.Black()
	o.Add(s)
	return o
}

func (o *Golor) BlackBack() *Golor {
	o.ColorBit8(BlackBack)
	return o
}

func (o *Golor) BlackBackAdd(s string) *Golor {
	o.BlackBack()
	o.Add(s)
	return o
}

func (o *Golor) Red() *Golor {
	o.ColorBit8(Red)
	return o
}

func (o *Golor) RedAdd(s string) *Golor {
	o.Red()
	o.Add(s)
	return o
}

func (o *Golor) RedBack() *Golor {
	o.ColorBit8(RedBack)
	return o
}

func (o *Golor) RedBackAdd(s string) *Golor {
	o.RedBack()
	o.Add(s)
	return o
}

func (o *Golor) Green() *Golor {
	o.ColorBit8(Green)
	return o
}

func (o *Golor) GreenAdd(s string) *Golor {
	o.Green()
	o.Add(s)
	return o
}

func (o *Golor) GreenBack() *Golor {
	o.ColorBit8(GreenBack)
	return o
}

func (o *Golor) GreenBackAdd(s string) *Golor {
	o.GreenBack()
	o.Add(s)
	return o
}

func (o *Golor) Yellow() *Golor {
	o.ColorBit8(Yellow)
	return o
}

func (o *Golor) YellowAdd(s string) *Golor {
	o.Yellow()
	o.Add(s)
	return o
}

func (o *Golor) YellowBack() *Golor {
	o.ColorBit8(YellowBack)
	return o
}

func (o *Golor) YellowBackAdd(s string) *Golor {
	o.YellowBack()
	o.Add(s)
	return o
}

func (o *Golor) Blue() *Golor {
	o.ColorBit8(Blue)
	return o
}

func (o *Golor) BlueAdd(s string) *Golor {
	o.Blue()
	o.Add(s)
	return o
}

func (o *Golor) BlueBack() *Golor {
	o.ColorBit8(BlueBack)
	return o
}

func (o *Golor) BlueBackAdd(s string) *Golor {
	o.BlueBack()
	o.Add(s)
	return o
}

func (o *Golor) Magenta() *Golor {
	o.ColorBit8(Magenta)
	return o
}

func (o *Golor) MagentaAdd(s string) *Golor {
	o.Magenta()
	o.Add(s)
	return o
}

func (o *Golor) MagentaBack() *Golor {
	o.ColorBit8(MagentaBack)
	return o
}

func (o *Golor) MagentaBackAdd(s string) *Golor {
	o.MagentaBack()
	o.Add(s)
	return o
}

func (o *Golor) Cyan() *Golor {
	o.ColorBit8(Cyan)
	return o
}

func (o *Golor) CyanAdd(s string) *Golor {
	o.Cyan()
	o.Add(s)
	return o
}

func (o *Golor) CyanBack() *Golor {
	o.ColorBit8(CyanBack)
	return o
}

func (o *Golor) CyanBackAdd(s string) *Golor {
	o.CyanBack()
	o.Add(s)
	return o
}

func (o *Golor) White() *Golor {
	o.ColorBit8(White)
	return o
}

func (o *Golor) WhiteAdd(s string) *Golor {
	o.White()
	o.Add(s)
	return o
}

func (o *Golor) WhiteBack() *Golor {
	o.ColorBit8(WhiteBack)
	return o
}

func (o *Golor) WhiteBackAdd(s string) *Golor {
	o.WhiteBack()
	o.Add(s)
	return o
}

func (o *Golor) Default() *Golor {
	o.ColorBit8(Default)
	return o
}

func (o *Golor) DefaultAdd(s string) *Golor {
	o.Default()
	o.Add(s)
	return o
}

func (o *Golor) DefaultBack() *Golor {
	o.ColorBit8(DefaultBack)
	return o
}

func (o *Golor) DefaultBackAdd(s string) *Golor {
	o.DefaultBack()
	o.Add(s)
	return o
}

func (o *Golor) ColorBit256(i int) *Golor {
	if Mode == Bit256 {
		o.sb.WriteString("\x1b[38;5;")
		o.sb.WriteString(strconv.Itoa(i))
		o.sb.WriteString("m")
	}
	return o
}

func (o *Golor) ColorBit256Add(s string, i int) *Golor {
	o.ColorBit256(i)
	o.Add(s)
	return o
}

func (o *Golor) ColorBit256Back(i int) *Golor {
	if Mode == Bit256 {
		o.sb.WriteString("\x1b[48;5;")
		o.sb.WriteString(strconv.Itoa(i))
		o.sb.WriteString("m")
	}
	return o
}

func (o *Golor) ColorBit256BackAdd(s string, i int) *Golor {
	o.ColorBit256Back(i)
	o.Add(s)
	return o
}

func (o *Golor) ColorBitRGB(r int, g int, b int) *Golor {
	if Mode == BitRGB {
		o.sb.WriteString("\x1b[38;2;")
		o.sb.WriteString(strconv.Itoa(r))
		o.sb.WriteString(";")
		o.sb.WriteString(strconv.Itoa(g))
		o.sb.WriteString(";")
		o.sb.WriteString(strconv.Itoa(b))
		o.sb.WriteString("m")
	}
	return o
}

func (o *Golor) ColorBitRGBAdd(s string, r int, g int, b int) *Golor {
	o.ColorBitRGB(r, g, b)
	o.Add(s)
	return o
}

func (o *Golor) ColorBitRGBBack(r int, g int, b int) *Golor {
	if Mode == BitRGB {
		o.sb.WriteString("\x1b[48;2;")
		o.sb.WriteString(strconv.Itoa(r))
		o.sb.WriteString(";")
		o.sb.WriteString(strconv.Itoa(g))
		o.sb.WriteString(";")
		o.sb.WriteString(strconv.Itoa(b))
		o.sb.WriteString("m")
	}
	return o
}

func (o *Golor) ColorBitRGBBackAdd(s string, r int, g int, b int) *Golor {
	o.ColorBitRGBBack(r, g, b)
	o.Add(s)
	return o
}
