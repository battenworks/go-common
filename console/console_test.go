package console

import "testing"

func TestConsoleOut(t *testing.T) {
	Out("should be blue: ")
	Blue("blue")
	Outln("")

	Out("should be blue: ")
	Blueln("blue")

	Out("should be blue bold: ")
	BlueBold("blue bold")
	Outln("")

	Out("should be blue bold: ")
	BluelnBold("blue bold")

	Out("should be cyan: ")
	Cyan("cyan")
	Outln("")

	Out("should be cyan: ")
	Cyanln("cyan")

	Out("should be cyan bold: ")
	CyanBold("cyan bold")
	Outln("")

	Out("should be cyan bold: ")
	CyanlnBold("cyan bold")

	Out("should be green: ")
	Green("green")
	Outln("")

	Out("should be green: ")
	Greenln("green")

	Out("should be green bold: ")
	GreenBold("green bold")
	Outln("")

	Out("should be green bold: ")
	GreenlnBold("green bold")

	Out("should be magenta: ")
	Magenta("magenta")
	Outln("")

	Out("should be magenta: ")
	Magentaln("magenta")

	Out("should be magenta bold: ")
	MagentaBold("magenta bold")
	Outln("")

	Out("should be magenta bold: ")
	MagentalnBold("magenta bold")

	Out("should be red: ")
	Red("red")
	Outln("")

	Out("should be red: ")
	Redln("red")

	Out("should be red bold: ")
	RedBold("red bold")
	Outln("")

	Out("should be red bold: ")
	RedlnBold("red bold")

	Out("should be white: ")
	White("white")
	Outln("")

	Out("should be white: ")
	Whiteln("white")

	Out("should be white bold: ")
	WhiteBold("white bold")
	Outln("")

	Out("should be white bold: ")
	WhitelnBold("white bold")

	Out("should be yellow: ")
	Yellow("yellow")
	Outln("")

	Out("should be yellow: ")
	Yellowln("yellow")

	Out("should be yellow bold: ")
	YellowBold("yellow bold")
	Outln("")

	Out("should be yellow bold: ")
	YellowlnBold("yellow bold")
}