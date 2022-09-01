package golor

import (
	"strings"
)

type GolorMode int

const (
	Plain GolorMode = iota
	Bit3_4
	Bit8
	Bit24
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

func New() Golor {
	g := Golor{
		len: 0,
	}
	return g
}

func (g *Golor) Len() int {
	return g.len
}

func (g *Golor) Build() string {
	return g.sb.String()
}

func (g *Golor) Write(s string) *Golor {
	g.len += len(s)
	g.sb.WriteString(s)
	return g
}

func (g *Golor) Reset() *Golor {
	if Mode != Plain {
		g.sb.WriteString(string(Reset))
	}
	return g
}

func (g *Golor) ResetAdd(s string) *Golor {
	g.Reset()
	g.Write(s)
	return g
}

func (g *Golor) Bell() *Golor {
	if Mode != Plain && Sound != Mute {
		g.sb.WriteString("\a")
	}
	return g
}

func (g *Golor) BellAdd(s string) *Golor {
	g.Bell()
	g.Write(s)
	return g
}

func (g *Golor) Bold() *Golor {
	if Mode != Plain {
		g.sb.WriteString(string(Bold))
	}
	return g
}

func (g *Golor) BoldAdd(s string) *Golor {
	g.Bold()
	g.Write(s)
	return g
}

func (g *Golor) BoldReset() *Golor {
	if Mode != Plain {
		g.sb.WriteString(string(Boldr))
	}
	return g
}

func (g *Golor) BoldResetAdd(s string) *Golor {
	g.BoldReset()
	g.Write(s)
	return g
}

func (g *Golor) Faint() *Golor {
	if Mode != Plain {
		g.sb.WriteString(string(Faint))
	}
	return g
}

func (g *Golor) FaintAdd(s string) *Golor {
	g.Faint()
	g.Write(s)
	return g
}

func (g *Golor) FaintReset() *Golor {
	if Mode != Plain {
		g.sb.WriteString(string(Faintr))
	}
	return g
}

func (g *Golor) FaintResetAdd(s string) *Golor {
	g.FaintReset()
	g.Write(s)
	return g
}

func (g *Golor) Italic() *Golor {
	if Mode != Plain {
		g.sb.WriteString(string(Italic))
	}
	return g
}

func (g *Golor) ItalicAdd(s string) *Golor {
	g.Italic()
	g.Write(s)
	return g
}

func (g *Golor) ItalicReset() *Golor {
	if Mode != Plain {
		g.sb.WriteString(string(Italicr))
	}
	return g
}

func (g *Golor) ItalicResetAdd(s string) *Golor {
	g.ItalicReset()
	g.Write(s)
	return g
}

func (g *Golor) Underline() *Golor {
	if Mode != Plain {
		g.sb.WriteString(string(Underline))
	}
	return g
}

func (g *Golor) UnderlineAdd(s string) *Golor {
	g.Underline()
	g.Write(s)
	return g
}

func (g *Golor) UnderlineReset() *Golor {
	if Mode != Plain {
		g.sb.WriteString(string(Underliner))
	}
	return g
}

func (g *Golor) UnderlineResetAdd(s string) *Golor {
	g.UnderlineReset()
	g.Write(s)
	return g
}

func (g *Golor) Blinking() *Golor {
	if Mode != Plain {
		g.sb.WriteString(string(Blinking))
	}
	return g
}

func (g *Golor) BlinkingAdd(s string) *Golor {
	g.Blinking()
	g.Write(s)
	return g
}

func (g *Golor) BlinkingReset() *Golor {
	if Mode != Plain {
		g.sb.WriteString(string(Blinkingr))
	}
	return g
}

func (g *Golor) BlinkingResetAdd(s string) *Golor {
	g.BlinkingReset()
	g.Write(s)
	return g
}

func (g *Golor) Reverse() *Golor {
	if Mode != Plain {
		g.sb.WriteString(string(Reverse))
	}
	return g
}

func (g *Golor) ReverseAdd(s string) *Golor {
	g.Reverse()
	g.Write(s)
	return g
}

func (g *Golor) ReverseReset() *Golor {
	if Mode != Plain {
		g.sb.WriteString(string(Reverser))
	}
	return g
}

func (g *Golor) ReverseResetAdd(s string) *Golor {
	g.ReverseReset()
	g.Write(s)
	return g
}

func (g *Golor) Hidden() *Golor {
	if Mode != Plain {
		g.sb.WriteString(string(Hidden))
	}
	return g
}

func (g *Golor) HiddenAdd(s string) *Golor {
	g.Hidden()
	g.Write(s)
	return g
}

func (g *Golor) HiddenReset() *Golor {
	if Mode != Plain {
		g.sb.WriteString(string(Hiddenr))
	}
	return g
}

func (g *Golor) HiddenResetAdd(s string) *Golor {
	g.HiddenReset()
	g.Write(s)
	return g
}

func (g *Golor) Strikerthrough() *Golor {
	if Mode != Plain {
		g.sb.WriteString(string(Strikerthrough))
	}
	return g
}

func (g *Golor) StrikerthroughAdd(s string) *Golor {
	g.Strikerthrough()
	g.Write(s)
	return g
}

func (g *Golor) StrikerthroughReset() *Golor {
	if Mode != Plain {
		g.sb.WriteString(string(Strikerthroughr))
	}
	return g
}

func (g *Golor) StrikerthroughResetAdd(s string) *Golor {
	g.StrikerthroughReset()
	g.Write(s)
	return g
}

//References https://en.wikipedia.org/wiki/ANSI_escape_code
