package pingpp_log

import (
	"runtime"
	"path"
	"time"
	"fmt"
	"sync"
	"os"
	"io"
)

type Level int

const (
	Alert Level = iota
	Error
	Warn
	Info
	Debug
)

type Logger struct {
	config LogConfig
	info   LogInfo
	level  Level
	mu     sync.Mutex
}

func New(c LogConfig) *Logger {
	return &Logger{config: c}
}

func (l *Logger) SetModule(module string) *Logger {
	l.info.Module = module
	return l
}

func (l *Logger)SetInfo(info *LogInfo) *Logger {
	l.info = info
	return l
}

func (l *Logger) Output(level string, s string) {
	now := time.Now()
	var file string
	var line int
	var ok bool
	_, file, line, ok = runtime.Caller(1)
	if !ok {
		file = "???"
		line = 0
	}
	_, filename := path.Split(file)
	l.mu.Lock()
	defer l.mu.Unlock()
	l.info.Filename = filename
	l.info.Line = line
	l.info.LogLevel = level
	l.info.LogTime = now.Format("2006/01/02 15:04:05")
	l.info.Remark = s
	if l.config.LogPlace & ToConsole {
		l.ToConsole()
	}else if l.config.LogPlace & ToFile {
		l.ToFile()
	}else if l.config.LogPlace & ToFile | ToConsole {
		l.ToFileAndStdout()
	}
}

func (l *Logger)Debug(format string, a...interface{}) {
	if l.level >= Debug {
		l.Output("Debug", fmt.Sprintf(format, a))
	}
}

func (l *Logger)Warn(format string, a...interface{}) {
	if l.level >= Warn {
		l.Output("Warn", fmt.Sprintf(format, a))
	}
}

func (l *Logger)Alert(format string, a...interface{}) {
	if l.level >= Alert {
		l.Output("Alert", fmt.Sprintf(format, a))
	}
}

func (l *Logger)Info(format string, a...interface{}) {
	if l.level >= Info {
		l.Output("Info", fmt.Sprintf(format, a))
	}
}

func (l *Logger)Error(format string, a...interface{}) {
	if l.level >= Error {
		l.Output("Error", fmt.Sprintf(format, a))
	}
}

func (l *Logger)ToConsole() error {
	return l.Writeplace(os.Stdout)
}

func (l *Logger)ToFile() error {
	LogFile, err := os.OpenFile(l.config.Logfile, os.O_RDWR | os.O_CREATE, 0777)
	if err != nil {
		return err
	}
	return l.Writeplace(LogFile)
}

func (l *Logger)ToFileAndStdout() error {
	var writers []io.Writer
	LogFile, err := os.OpenFile(l.config.Logfile, os.O_RDWR | os.O_CREATE, 0777)
	if err != nil {
		return err
	}
	writers = append(writers, LogFile)
	writers = append(writers, os.Stdout)
	fileAndStdoutWriter := io.MultiWriter(writers...)
	return l.Writeplace(fileAndStdoutWriter)
}

func (l *Logger)Writeplace(iw io.Writer) error {
	var strByte []byte
	var err error
	if PlaceContentType[l.config.LogPlace] == JSON {
		strByte, err = l.info.ToJson()
		if err != nil {
			return err
		}

	}else if PlaceContentType[l.config.LogPlace] == STDOUTPUT {
		strByte = []byte(l.info.ToStd())
	}
	_, err = iw.Write(strByte)
	if err != nil {
		return err
	}
	return nil
}


