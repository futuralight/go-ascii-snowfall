package main

import (
	"errors"
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"strings"
	"time"

	"github.com/buger/goterm"
	"github.com/fatih/color"
)

//DefaultSnowChar - default snow character
const DefaultSnowChar = "*"

//DefaultFlakesRatio - default flakes in string ratio
const DefaultFlakesRatio = 3

//DefaultPrintDelayMilliSeconds - default delay betwen printing
const DefaultPrintDelayMilliSeconds = 800

//snowStringGetter is func type for getting string of snow
type snowStringGetter func(int, string, int) string

//flakesRatio - flakes in string ratio
var flakesRatio int

//printDelayMilliSeconds - delay betwen printing
var printDelayMilliSeconds time.Duration

//snowChar - snow character
var snowChar string

var colorMap = map[string]color.Attribute{
	"white":   color.FgWhite,
	"black":   color.FgBlack,
	"red":     color.FgRed,
	"blue":    color.FgBlue,
	"magneta": color.FgMagenta,
	"cyan":    color.FgCyan,
	"green":   color.FgGreen,
	"yellow":  color.FgYellow,
}

var bgColorMap = map[string]color.Attribute{
	"white":   color.BgWhite,
	"black":   color.BgBlack,
	"red":     color.BgRed,
	"blue":    color.BgBlue,
	"magneta": color.BgMagenta,
	"cyan":    color.BgCyan,
	"green":   color.BgGreen,
	"yellow":  color.BgYellow,
}

func main() {
	flakesRatio = DefaultFlakesRatio
	printDelayMilliSeconds = DefaultPrintDelayMilliSeconds
	snowChar = DefaultSnowChar
	err := argsCheck()
	if err != nil {
		fmt.Println("Error: " + err.Error())
	} else {
		snowfall()
	}
}

func argsCheck() error {
	args := os.Args[1:]
	for i, v := range args {
		//Ratio
		if v == "-r" {
			if len(args)-1 < i+1 {
				return errors.New("The ratio is missing")
			}
			ratio, err := strconv.Atoi(strings.TrimSpace(args[i+1]))
			if err != nil {
				return err
			}
			if ratio > 100 {
				return errors.New("The ratio must be less than 100")
			}
			flakesRatio = ratio
		}
		//Flake
		if v == "-f" {
			if len(args)-1 < i+1 {
				return errors.New("The flake character is missing")
			}
			snowChar = strings.TrimSpace(args[i+1])
		}
		//Delay
		if v == "-d" {
			if len(args)-1 < i+1 {
				return errors.New("The delay is missing")
			}
			delay, err := strconv.Atoi(strings.TrimSpace(args[i+1]))
			if err != nil {
				return err
			}
			printDelayMilliSeconds = time.Duration(delay)
		}
		//Color
		if v == "-c" {
			if len(args)-1 < i+1 {
				return errors.New("The color is missing")
			}
			clr, ok := colorMap[args[i+1]]
			if !ok {
				return errors.New("No such color")
			}
			color.Set(clr)
		}
		//Background color
		if v == "-bc" {
			if len(args)-1 < i+1 {
				return errors.New("The background color is missing")
			}
			bgClr, ok := bgColorMap[args[i+1]]
			if !ok {
				return errors.New("No such color")
			}
			color.Set(bgClr)
		}
	}
	return nil
}

func snowfall() {
	height, width, err := getTerminalSize()
	if err != nil {
		panic(err)
	}
	for {
		fmt.Println(getScreen(height, width, flakesRatio, snowChar, getStringArray))
		time.Sleep(time.Millisecond * printDelayMilliSeconds)
	}
}

func getTerminalSize() (int, int, error) {
	switch runtime.GOOS {
	case "windows":
		return getTerminalSizeWindows()
	case "linux", "freebsd":
		return getTerminalSizeUnix()
	default:
		return 0, 0, errors.New("OS don't support")
	}
}

func getTerminalSizeUnix() (int, int, error) {
	cmd := exec.Command("stty", "size")
	cmd.Stdin = os.Stdin
	out, err := cmd.Output()
	if err != nil {
		return 0, 0, err
	}
	sizeSlice := strings.Split(strings.TrimSpace(string(out)), " ")
	if len(sizeSlice) != 2 {
		return 0, 0, errors.New("Wrong stty size output")
	}
	height, err := strconv.Atoi(sizeSlice[0])
	if err != nil {
		return 0, 0, err
	}
	width, err := strconv.Atoi(sizeSlice[1])
	if err != nil {
		return 0, 0, err
	}
	return height, width, nil
}

func getTerminalSizeWindows() (int, int, error) {
	height := goterm.Height()
	width := goterm.Width()
	return height, width, nil
}

func getScreen(height, width, ratio int, snowFlake string, fn snowStringGetter) string {
	screen := make([]string, height)
	for i := range screen {
		screen[i] = fn(width, snowFlake, ratio)
	}
	return strings.Join(screen, "\n")
}

func getStringConcat(width int, snowflake string, ratio int) string {
	snowString := ""
	for i := 0; i < width; i++ {
		r := rand.Intn(100)
		if r < ratio {
			snowString += snowflake
		} else {
			snowString += " "
		}
	}
	return snowString
}

func getStringArray(width int, snowflake string, ratio int) string {
	slice := make([]string, width)
	for i := range slice {
		r := rand.Intn(100)
		if r < ratio {
			slice[i] = snowflake
		} else {
			slice[i] = " "
		}
	}
	return strings.Join(slice, "")
}

func getStringAppend(width int, snowflake string, ratio int) string {
	snowSlice := make([]string, 0, width)
	for i := 0; i < width; i++ {
		r := rand.Intn(100)
		if r < ratio {
			snowSlice = append(snowSlice, snowflake)
		} else {
			snowSlice = append(snowSlice, " ")
		}
	}
	return strings.Join(snowSlice, "")
}
