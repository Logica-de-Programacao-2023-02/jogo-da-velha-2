package main

import "fmt"

type Color string

const (
	ColorReset  Color = "\033[0m"
	ColorRed    Color = "\033[31m"
	ColorGreen  Color = "\033[32m"
	ColorYellow Color = "\033[33m"
	ColorBlue   Color = "\033[34m"
	ColorPurple Color = "\033[35m"
	ColorCyan   Color = "\033[36m"
	ColorGray   Color = "\033[37m"
	ColorWhite  Color = "\033[97m"
)

func (c Color) Paint(msg string) string {
	return fmt.Sprintf("%s%s%s", c, msg, ColorReset)
}
