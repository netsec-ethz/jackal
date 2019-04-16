/*
 * Copyright (c) 2018 Miguel Ángel Ortuño.
 * See the LICENSE file for more information.
 */

package log

import (
	"fmt"
	"io"
	"os"
	"path"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/ortuman/jackal/util"
)

const logChanBufferSize = 2048

const projectFolder = "jackal"

var exitHandler = func() { os.Exit(-1) }

// Level represents log level type.
type Level int

const (
	// DebugLevel represents DEBUG log level.
	DebugLevel Level = iota

	// InfoLevel represents INFO log level.
	InfoLevel

	// WarningLevel represents WARNING log level.
	WarningLevel

	// ErrorLevel represents ERROR log level.
	ErrorLevel

	// FatalLevel represents FATAL log level.
	FatalLevel

	// OffLevel represents a disabledLogger log level.
	OffLevel
)

// Logger represents a common logger interface.
type Logger interface {
	io.Closer

	Level() Level
	Log(level Level, logLine string)
}

// Debugf writes a 'debug' message to configured logger.
func Debugf(format string, args ...interface{}) {
	logLevel(DebugLevel, fmt.Sprintf(format, args...))
}

// Debug writes a set of arguments to the configured 'debug' logger.
func Debug(args ...interface{}) {
	logLevel(DebugLevel, fmt.Sprintf(util.StringRepeat("%v", " ", len(args)), args...))
}

// Infof writes a 'info' message to configured logger.
func Infof(format string, args ...interface{}) {
	logLevel(InfoLevel, fmt.Sprintf(format, args...))
}

// Info writes a set of arguments to the configured 'info' logger.
func Info(args ...interface{}) {
	logLevel(InfoLevel, fmt.Sprintf(util.StringRepeat("%v", " ", len(args)), args...))
}

// Warnf writes a 'warning' message to configured logger.
func Warnf(format string, args ...interface{}) {
	logLevel(WarningLevel, fmt.Sprintf(format, args...))
}

// Warn writes a set of arguments to the configured 'warning' logger.
func Warn(args ...interface{}) {
	logLevel(WarningLevel, fmt.Sprintf(util.StringRepeat("%v", " ", len(args)), args...))
}

// Errorf writes an 'error' message to configured logger.
func Errorf(format string, args ...interface{}) {
	logLevel(ErrorLevel, fmt.Sprintf(format, args...))
}

// Error writes a set of arguments to the configured 'error' logger.
func Error(args ...interface{}) {
	logLevel(ErrorLevel, fmt.Sprintf(util.StringRepeat("%v", " ", len(args)), args...))
}

// Fatalf writes a 'fatal' message to configured logger.
// Application should terminate after logging.
func Fatalf(format string, args ...interface{}) {
	logLevel(FatalLevel, fmt.Sprintf(format, args...))
}

// Fatal writes a set of arguments value to the configured logger.
// Application should terminate after logging.
func Fatal(args ...interface{}) {
	logLevel(FatalLevel, fmt.Sprintf(util.StringRepeat("%v", " ", len(args)), args...))
}

func logLevel(level Level, log string) {
	if inst := instance(); inst.Level() <= level {
		ci := getCallerInfo()
		logLine := formatLogLine(level, ci.pkg, ci.filename, ci.line, log)
		inst.Log(level, logLine)
	}
}

var (
	instMu sync.RWMutex
	inst   Logger
)

// Disabled stores a disabled logger instance.
var Disabled Logger = &disabledLogger{}

func init() {
	inst = Disabled
}

// Set sets the global logger.
func Set(logger Logger) {
	instMu.Lock()
	_ = inst.Close()
	inst = logger
	instMu.Unlock()
}

// Unset disables a previously set global logger.
func Unset() {
	Set(Disabled)
}

func instance() Logger {
	instMu.RLock()
	l := inst
	instMu.RUnlock()
	return l
}

type callerInfo struct {
	pkg      string
	filename string
	line     int
}

type record struct {
	logLine    string
	continueCh chan struct{}
}

type logger struct {
	level  Level
	output io.Writer
	files  []io.WriteCloser
	recCh  chan record
}

// New returns a default logger instance.
func New(level string, output io.Writer, files ...io.WriteCloser) (Logger, error) {
	lvl, err := levelFromString(level)
	if err != nil {
		return nil, err
	}
	l := &logger{
		level:  lvl,
		output: output,
		files:  files,
	}
	l.recCh = make(chan record, logChanBufferSize)
	go l.loop()
	return l, nil
}

func (l *logger) Level() Level {
	return l.level
}

func (l *logger) Log(level Level, logLine string) {
	entry := record{
		logLine:    logLine,
		continueCh: make(chan struct{}),
	}
	select {
	case l.recCh <- entry:
		if level == FatalLevel {
			<-entry.continueCh // wait until done
			exitHandler()
		}
	default:
		break // avoid blocking...
	}
}

func (l *logger) Close() error {
	close(l.recCh)
	return nil
}

func (l *logger) loop() {
	for {
		select {
		case rec, ok := <-l.recCh:
			if !ok {
				// close log files
				for _, w := range l.files {
					_ = w.Close()
				}
				return
			}

			_, _ = fmt.Fprintf(l.output, rec.logLine)
			for _, w := range l.files {
				_, _ = fmt.Fprintf(w, rec.logLine)
			}
			close(rec.continueCh)
		}
	}
}

func getCallerInfo() callerInfo {
	ci := callerInfo{}
	_, file, ln, ok := runtime.Caller(2)
	if ok {
		ci.pkg = filepath.Base(path.Dir(file))
		if ci.pkg == projectFolder {
			ci.pkg = ""
		}
		filename := filepath.Base(file)
		ci.filename = strings.TrimSuffix(filename, filepath.Ext(filename))
		ci.line = ln
	} else {
		ci.filename = "???"
		ci.pkg = "???"
	}
	return ci
}

func formatLogLine(level Level, pkg string, file string, line int, log string) string {
	var b strings.Builder

	b.WriteString(time.Now().Format("2006-01-02 15:04:05"))
	b.WriteString(" ")
	b.WriteString(logLevelGlyph(level))
	b.WriteString(" [")
	b.WriteString(logLevelAbbreviation(level))
	b.WriteString("] ")

	b.WriteString(pkg)
	if len(pkg) > 0 {
		b.WriteString("/")
	}
	b.WriteString(file)
	b.WriteString(":")
	b.WriteString(strconv.Itoa(line))
	b.WriteString(" - ")
	b.WriteString(log)
	b.WriteString("\n")

	return b.String()
}

func logLevelAbbreviation(level Level) string {
	switch level {
	case DebugLevel:
		return "DBG"
	case InfoLevel:
		return "INF"
	case WarningLevel:
		return "WRN"
	case ErrorLevel:
		return "ERR"
	case FatalLevel:
		return "FTL"
	default:
		return ""
	}
}

func logLevelGlyph(level Level) string {
	switch level {
	case DebugLevel:
		return "\U0001f50D"
	case InfoLevel:
		return "\u2139\ufe0f"
	case WarningLevel:
		return "\u26a0\ufe0f"
	case ErrorLevel:
		return "\U0001f4a5"
	case FatalLevel:
		return "\U0001f480"
	default:
		return ""
	}
}

func levelFromString(level string) (Level, error) {
	switch strings.ToLower(level) {
	case "debug":
		return DebugLevel, nil
	case "", "info":
		return InfoLevel, nil
	case "warning":
		return WarningLevel, nil
	case "error":
		return ErrorLevel, nil
	case "fatal":
		return FatalLevel, nil
	case "off":
		return OffLevel, nil
	}
	return Level(-1), fmt.Errorf("log: unrecognized level: %s", level)
}
