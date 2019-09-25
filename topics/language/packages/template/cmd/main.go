package main

import (
	"fmt"
	"miku/aboutgo/topics/language/packages/template"

	"github.com/fatih/color"
	// xcolor "github.com/x/color" // alias
)

func main() {
	color.Cyan("Prints text in cyan.")
	fmt.Println(template.Inc(10))
}
