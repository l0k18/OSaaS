// Package log is a simple logger which prints to stderr, and will later integrate with the spore IPC. This library is
// intended to be dot-imported and names are meant to not clash with local namespace
package log

import (
	"fmt"
	"github.com/niubaoshu/gotiny"
	"io"
	"io/ioutil"
	"os"
	"runtime"
	"time"
)

type entry struct {
	time    time.Time
	level   string
	loc     string
	message string
}

var lvl = 4
var levels = map[string]int{
	"fatal": 0,
	"error": 1,
	"info":  2,
	"debug": 3,
	"trace": 4,
}

// SetLogLevel sets the logging level of the logger
func SetLogLevel(level string) {
	if lv, ok := levels[level]; ok {
		lvl = lv
	}
}

var tty io.Writer = os.Stderr
var pipe io.WriteCloser
var start = time.Now()

func SetPipeWriter(p io.WriteCloser) {
	pipe = p
}

// DisableTTYLogging disables logging to tty
func DisableTTYLogging() {
	tty = ioutil.Discard
}

// EnableTTYLogging sets the tty logging to go to stderr
func EnableTTYLogging() {
	tty = os.Stderr
}

// the following implement printing to stderr
func prtty(level, loc string, a ...interface{}) bool {
	if levels[level] > lvl {
		return false
	}
	_, _ = fmt.Fprintln(tty, time.Now().Sub(start), level, fmt.Sprintln(a...), loc)
	return true
}
func prttyf(level, loc string, format string, a ...interface{}) bool {
	if levels[level] > lvl {
		return false
	}
	_, _ = fmt.Fprintln(tty, time.Now().Sub(start), level, fmt.Sprintf(format, a...))
	return true
}

// the following print gotiny encoded binary for pipe and socket connections
func prpipe(level, loc string, a ...interface{}) {
	if pipe == nil {
		return
	}
	_, _ = pipe.Write(
		gotiny.Marshal(
			"log-v0.0.1",
			entry{
				time:    time.Now(),
				level:   level,
				loc:     loc,
				message: fmt.Sprintln(a...),
			},
		),
	)
}
func prpipef(level, loc string, format string, a ...interface{}) {
	if pipe == nil {
		return
	}
	_, _ = pipe.Write(
		gotiny.Marshal(
			"log-v0.0.1",
			entry{
				time:    time.Now(),
				level:   level,
				loc:     loc,
				message: fmt.Sprintf(format, a...),
			},
		),
	)
}

func getLoc() string {
	_, file, line, _ := runtime.Caller(2)
	return fmt.Sprintf("%s:%d", file, line)
}

func Fatal(a ...interface{}) {
	if prtty("fatal", getLoc(), a...) {
		prpipe("fatal", getLoc(), a...)
	}
}
func Error(a ...interface{}) {
	if prtty("error", getLoc(), a...) {
		prpipe("error", getLoc(), a...)
	}
}
func Info(a ...interface{}) {
	if prtty("info", getLoc(), a...) {
		prpipe("info", getLoc(), a...)
	}
}
func Debug(a ...interface{}) {
	if prtty("debug", getLoc(), a...) {
		prpipe("debug", getLoc(), a...)
	}
}
func Trace(a ...interface{}) {
	if prtty("trace", getLoc(), a...) {
		prpipe("trace", getLoc(), a...)
	}
}

func Fatalf(format string, a ...interface{}) {
	if prttyf("fatal", getLoc(), format, a...) {
		prpipef("fatal", getLoc(), format, a...)
	}
}
func Errorf(format string, a ...interface{}) {
	if prttyf("error", getLoc(), format, a...) {
		prpipef("error", getLoc(), format, a...)
	}
}
func Infof(format string, a ...interface{}) {
	if prttyf("info", getLoc(), format, a...) {
		prpipef("info", getLoc(), format, a...)
	}
}
func Debugf(format string, a ...interface{}) {
	if prttyf("debug", getLoc(), format, a...) {
		prpipef("debug", getLoc(), format, a...)
	}
}
func Tracef(format string, a ...interface{}) {
	if prttyf("trace", getLoc(), format, a...) {
		prpipef("trace", getLoc(), format, a...)
	}
}

func Fatalc(printer func() string) {
	pr := printer()
	if prtty("fatal", getLoc(), pr) {
		prpipe("fatal", getLoc(), pr)
	}
}
func Errorc(printer func() string) {
	pr := printer()
	if prtty("error", getLoc(), pr) {
		prpipe("error", getLoc(), pr)
	}
}
func Infoc(printer func() string) {
	pr := printer()
	if prtty("info", getLoc(), pr) {
		prpipe("info", getLoc(), pr)
	}
}
func Debugc(printer func() string) {
	pr := printer()
	if prtty("debug", getLoc(), pr) {
		prpipe("debug", getLoc(), pr)
	}
}
func Tracec(printer func() string) {
	pr := printer()
	if prtty("trace", getLoc(), pr) {
		prpipe("trace", getLoc(), pr)
	}
}

func Check(err error) bool {
	if err != nil {
		Error(err)
		return true
	}
	return false
}

func Caller(comment string, skip int) string {
	_, file, line, _ := runtime.Caller(skip + 1)
	o := fmt.Sprintf("%s: %s:%d", comment, file, line)
	return o
}
