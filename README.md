# golor

This is a libary to provide you the functions to create string with different colors and styles. It have four modes:

- Plain(Default)
- Bit8
- Bit256
- BitRGB

Change the golor.Mode to select the mode that you want to. The usage of them depend of the compatibility of the ANSI Codes and your shell.

You can use different kind of graphic modes, this will work until the mode is not plain, this is the list:

- Reset
- Bold
- Faint
- Italic
- Underline
- Blinking
- Reverse
- Hidden
- Strikerthrough

About the **Bell**, this is a mode that make a sound, you can mute it changing the golor.Sound value:

- Active(default)
- Mute

## Examples

### Simple

```Go
package main

import (
	"fmt"

	"github.com/RobertoRojas/golor"
)

func main() {
	a := golor.New()
	a.Add("Hello").Add(" world").Add("!")
	fmt.Println(a.String())
}
```

### Bit8

```Go
package main

import (
	"fmt"

	"github.com/RobertoRojas/golor"
)

func main() {
	golor.Mode = golor.Bit8
	a := golor.New()
	a.Blue().Add("Hello").Reset().Red().
		WhiteBackAdd(" world").Default().
		Add("!").Reset()
	fmt.Println(a.String())
}
```

### Bit256

```Go
package main

import (
	"fmt"

	"github.com/RobertoRojas/golor"
)

func main() {
	golor.Mode = golor.Bit256
	a := golor.New()
	a.Bold().Underline().ColorBit256Add("Hello", 21).
		ColorBit256(88).ColorBit256BackAdd(" world", 255).
		UnderlineReset().Add("!").Reset()
	fmt.Println(a.String())
}
```

### BitRGB

```Go
package main

import (
	"fmt"

	"github.com/RobertoRojas/golor"
)

func main() {
	golor.Mode = golor.BitRGB
	a := golor.New()
	a.Bell().ColorBitRGBAdd("Hello", 21, 21, 21).
		ColorBitRGBAdd(" world", 255, 0, 255).
		ColorBitRGBAdd("!", 255, 255, 0).Reset()
	fmt.Println(a.String())
}
```

### Program

```Go
package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/RobertoRojas/golor"
)

var verbose bool

func showMessage(s string) {
	fmt.Println(s)
}

func showError(s string) {
	g := golor.New()
	g.Bell().Red().ColorBit256(160).
		ColorBitRGB(255, 0, 0).Add(s).
		Reset()
	fmt.Fprintln(os.Stderr, g.String())
}

func showVerbose(s string) {
	if !verbose {
		return
	}
	g := golor.New()
	g.Bold().Cyan().ColorBit256(38).
		ColorBitRGB(0, 255, 255).
		Add("Verbose: ").BoldReset().
		Add(s).Reset()
	fmt.Println(g.String())
}

func main() {
	var name string
	var lastname string
	var color string
	flag.StringVar(&name, "name", "", "Add your name")
	flag.StringVar(&lastname, "lastname", "default", "Add your lastname")
	flag.StringVar(&color, "color", "plain", "Select the color mode[bit8, bit256, bitrgb, plain(default)]")
	flag.BoolVar(&verbose, "verbose", false, "Add your name")
	flag.Parse()
	golor.SetMode(color)
	showVerbose("Start of the program...")
	if name == "" {
		showError("The name parameters is mandatory")
		os.Exit(1)
	}
	showMessage("Hi " + name + ", nice to meet you.")
	m := golor.New()
	m.Add("Your last name is ").BoldAdd(lastname).Reset()
	showMessage(m.String())
	showVerbose("End of the program...")
	os.Exit(0)
}
```

### References

- [ANSI Escape Sequences](https://gist.github.com/fnky/458719343aabd01cfb17a3a4f7296797)