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
)

//DefaultSnowChar - default snow character
const DefaultSnowChar = "*"

//DefaultFlakesRatio - default flakes in string ratio
const DefaultFlakesRatio = 3

//DefaultPrintDelayMilliSeconds - default delay betwen printing
const DefaultPrintDelayMilliSeconds = 800

//FlakesRatio - flakes in string ratio
var FlakesRatio int

//PrintDelayMilliSeconds - delay betwen printing
var PrintDelayMilliSeconds time.Duration

//SnowChar - snow character
var SnowChar string

func main() {
	FlakesRatio = DefaultFlakesRatio
	PrintDelayMilliSeconds = DefaultPrintDelayMilliSeconds
	SnowChar = DefaultSnowChar
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
		if v == "-r" { //ratio
			if len(args)-1 < i+1 {
				return errors.New("The ratio is missing")
			}
			ratio, err := strconv.Atoi(strings.TrimSpace(args[i+1]))
			if err != nil {
				return err
			}
			if ratio > 100 {
				return errors.New("The ratio is more than 100")
			}
			FlakesRatio = ratio
		}
		if v == "-f" { //flake
			if len(args)-1 < i+1 {
				return errors.New("The flake character is missing")
			}
			SnowChar = strings.TrimSpace(args[i+1])
		}
		if v == "-d" { //delay
			if len(args)-1 < i+1 {
				return errors.New("The delay is missing")
			}
			delay, err := strconv.Atoi(strings.TrimSpace(args[i+1]))
			if err != nil {
				return err
			}
			PrintDelayMilliSeconds = time.Duration(delay)
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
		fmt.Println(getScreen(height, width, FlakesRatio, SnowChar))
		time.Sleep(time.Millisecond * PrintDelayMilliSeconds)
	}
}

func getTerminalSize() (int, int, error) {
	switch os := runtime.GOOS; os {
	case "windows":
		return getTerminalSizeWindows()
	case "linux":
		return getTerminalSizeLinux()
	default:
		panic("OS don't support")
	}
	return 0, 0, nil
}

func getTerminalSizeLinux() (int, int, error) {
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

func getScreen(height, width, ratio int, snowFlake string) string {
	snowScreen := ""
	for i := 0; i < height; i++ {
		snowScreen += getString(width, snowFlake, ratio) + "\n"
	}
	return snowScreen
}

func getString(width int, snowflake string, ratio int) string {
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

func getScreenAppend(height, width, ratio int, snowFlake string) string {
	snowScreen := make([]string, 0, height)
	for i := 0; i < height; i++ {
		snowScreen = append(snowScreen, getStringAppend(width, snowFlake, ratio))
	}
	return strings.Join(snowScreen, "\n")
}
