package pingxx_log

import (
	"runtime"
	"path"
	"time"
	"fmt"
	"sync"
	"os"
	"io"
)

type Logger struct {
	config *LogConfig
	info   *LogInfo
	mu     sync.Mutex
}

func New(c *LogConfig) *Logger {
	return &Logger{config: c,info:new(LogInfo)}
}

func (l *Logger) GetConfig() *LogConfig {
	return l.config
}

func (l *Logger) GetLogInfo() *LogInfo {
	return l.info
}

func (l *Logger) SetModule(module string) *Logger {
	l.info.Module = module
	return l
}

func (l *Logger)SetLogInfo(info *LogInfo) *Logger {
	l.info = info
	return l
}

func (l *Logger) Output(level string, s string) {
	now := time.Now()
	var file string
	var line int
	var ok bool
	_, file, line, ok = runtime.Caller(2)
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
	if l.config.LogPlace & ToFile != 0&& l.config.LogPlace & ToConsole != 0 {
		l.ToFileAndStdout()
	}else if l.config.LogPlace & ToConsole != 0 {
		l.ToConsole()
	}else if l.config.LogPlace & ToFile != 0 {
		l.ToFile()
	}
}

func (l *Logger)Debug(format string, a ...interface{}) {
	if l.GetConfig().level >= Debug {
		l.Output("Debug", fmt.Sprintf(format, a...))
	}
}

func (l *Logger)Warn(format string, a...interface{}) {
	if l.GetConfig().level >= Warn {
		l.Output("Warn", fmt.Sprintf(format, a...))
	}
}

func (l *Logger)Alert(format string, a...interface{}) {
	if l.GetConfig().level >= Alert {
		l.Output("Alert", fmt.Sprintf(format, a...))
	}
}

func (l *Logger)Info(format string, a...interface{}) {
	if l.GetConfig().level >= Info {
		l.Output("Info", fmt.Sprintf(format, a...))
	}
}

func (l *Logger)Error(format string, a...interface{}) {
	if l.GetConfig().level >= Error {
		l.Output("Error", fmt.Sprintf(format, a...))
	}
}

func (l *Logger)ToConsole() error {
	if v, ok := l.GetConfig().PlaceContentType[ToConsole]; ok {
		return l.Writeplace(os.Stdout, v)
	}
	return fmt.Errorf("please set content type")
}

func (l *Logger)ToFile() error {
	LogFile, err := os.OpenFile(l.config.Logfile, os.O_RDWR | os.O_CREATE, 0777)
	if err != nil {
		return err
	}
	if v, ok := l.GetConfig().PlaceContentType[ToFile]; ok {
		return l.Writeplace(LogFile, v)
	}
	return fmt.Errorf("please set content type")
}

func (l *Logger)ToFileAndStdout() error {
	err := l.ToConsole()
	if err != nil {
		return err
	}
	return l.ToFile()
}

func (l *Logger)Writeplace(iw io.Writer, cType ContentType) error {
	var strByte []byte
	var s string
	var err error
	if cType == JSON {
		strByte, err = l.info.ToJson()
		if err != nil {
			return err
		}
		s=string(strByte)
	}else if cType == STDOUTPUT {
		strByte = []byte(l.info.ToStd())
		s=string(strByte)
	}else {
		return fmt.Errorf("Unknow ContentType")
	}
	if len(s) == 0 || s[len(s)-1] != '\n' {
		strByte = append(strByte, '\n')
	}
	_, err = iw.Write(strByte)
	if err != nil {
		return err
	}
	return nil
}


