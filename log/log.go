package log

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

var LogLevel int = 0
var Cached bool = false

const (
	Reset  = "\033[0m"
	Gray   = "\033[90m"
	Blue   = "\033[34m"
	Yellow = "\033[33m"
	Red    = "\033[31m"
	Green  = "\033[32m"
)

func getLogLevel() int {
	if Cached {
		return LogLevel
	} else {
		level, done := os.LookupEnv("LOG_LEVEL")

		if done {
			val, err := strconv.Atoi(level)

			if err == nil {
				LogLevel = val
			} else {
				fmt.Println("Log level not set, defaulting to Info")
				LogLevel = 1
			}
		} else {
			fmt.Println("Couldn't find log level variable, defaulting to Debug")
			LogLevel = 0
		}

		Cached = true
		return LogLevel
	}
}

func WriteConsole(color string, tag string, format any, a ...any) {
	utc := time.Now().UTC()
	timeStamp := fmt.Sprintf("%s UTC", utc.Format(time.RFC3339))

	level := fmt.Sprintf("%s| %s |", color, tag)

	var message string
	switch v := format.(type) {
	case string:
		if len(a) > 0 {
			message = fmt.Sprintf(v, a...)
		} else {
			message = v
		}
	default:
		message = fmt.Sprint(format)
	}

	fmt.Println(timeStamp, level, message, Reset)
}

func Debug(format any, a ...any) {
	if getLogLevel() <= 0 {
		WriteConsole(Gray, "DEBUG", format, a...)
	}
}

func Info(format any, a ...any) {
	if getLogLevel() <= 1 {
		WriteConsole(Blue, "INFO", format, a...)
	}
}

func Warn(format any, a ...any) {
	if getLogLevel() <= 2 {
		WriteConsole(Yellow, "WARN", format, a...)
	}
}

func Error(format any, a ...any) {
	if getLogLevel() <= 3 {
		WriteConsole(Red, "ERROR", format, a...)
	}
}

func Done(format any, a ...any) {
	if getLogLevel() <= 4 {
		WriteConsole(Green, "DONE", format, a...)
	}
}

func Print(format any, a ...any) {
	if getLogLevel() <= 5 {
		WriteConsole(Reset, " LOG ", format, a...)
	}
}
