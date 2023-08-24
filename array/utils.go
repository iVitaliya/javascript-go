package array

import (
	"fmt"

	"github.com/iVitaliya/colors-go"
)

const (
	_LOG = iota
	INFO
	DEBUG
	WARNING
	ERROR
)

func print(state int, text ...string) {
	var (
		st    string
		open  = colors.BrightBlack("[")
		close = colors.BrightBlack("]")
	)

	go func(_state int) {
		switch _state {
		case INFO:
			st = open + colors.BrightBlue("INFO") + close
			break
		case DEBUG:
			st = open + colors.Green("DEBUG") + close
			break
		case WARNING:
			st = open + colors.Dim(colors.BrightYellow("WARNING")) + close
			break
		case ERROR:
			st = open + colors.Red("ERROR") + close
			break
		}
	}(state)
}

func appendValue[T any](arr []T, value T) []T {
	_arr := arr

	_arr = append(_arr, value)

	return _arr
}

func searchIndex[T any](arr []T, term T) int {
	var trm string = fmt.Sprint(term)

	for i, v := range arr {
		val := fmt.Sprint(v)

		if trm == val {
			return i
		}
	}

	return -1
}
