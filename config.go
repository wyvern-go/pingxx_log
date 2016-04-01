package pingpp_log

type Place int
type ContentType string

const (
	ToConsole Place = 1 << iota
	ToFile
)

const (
	JSON ContentType = "JSON"
	STDOUTPUT ContentType = "STDOUTPUT"
)

type LogConfig struct {
	LogPlace Place
	Logfile  string
}

func NewConfig(p Place,filename string) *LogConfig{
	c:=new(LogConfig)
	c.LogPlace=p
	c.Logfile=filename
	return c
}