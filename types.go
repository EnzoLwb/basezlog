package cuslog


type Level uint8

const (
	DebugLevel Level = iota
	InfoLevel
	WarnLevel
	ErrorLevel
	PanicLevel
	FatalLevel
)

type Formatter string

const  (
	EncodeJson Formatter = "JSON"
	EncodeText Formatter = "TEXT"
)

var LevelNameMapping = map[Level]string{
	DebugLevel: "DEBUG",
	InfoLevel:  "INFO",
	WarnLevel:  "WARN",
	ErrorLevel: "ERROR",
	PanicLevel: "PANIC",
	FatalLevel: "FATAL",
}

type OutputType string

const  (
	OutPutFile OutputType = "file"
	OutPutES OutputType = "es"
)