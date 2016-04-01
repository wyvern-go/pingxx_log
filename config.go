package pingxx_log

type Place int
type ContentType string
type Level int

const (
	ToConsole Place = 1 << iota
	ToFile
)

const (
	JSON ContentType = "JSON"
	STDOUTPUT ContentType = "STDOUTPUT"
)

const (
	Alert Level = iota
	Error
	Warn
	Info
	Debug
)

type LogConfig struct {
	LogPlace         Place
	Logfile          string
	level            Level
	PlaceContentType map[Place]ContentType
}

func NewConfig(filename string, p Place) *LogConfig {
	c := new(LogConfig)
	c.LogPlace = p
	c.Logfile = filename
	c.PlaceContentType = make(map[Place]ContentType)
	return c
}

func (c *LogConfig) SetCententType(place Place, contenttype ContentType) *LogConfig {
	c.PlaceContentType[place] = contenttype
	return c
}

func (c *LogConfig)SetLevel(level Level) *LogConfig {
	c.level = level
	return c
}