package logs

import (
	"github.com/natefinch/lumberjack"
	logr "github.com/sirupsen/logrus"
	"log"
	"os"
)

type LogConfig struct {
	LogFilename   string `json:"logFilename"`
	LogMaxSize    int    `json:"logMaxSize"`
	LogMaxBackups int    `json:"logMaxBackups"`
	LogMaxAge     int    `json:"logMaxAge"`
}

func InitLogger(filename string, maxSize, maxBackups, maxAge int) *LogConfig {
	return &LogConfig{
		filename,
		maxSize,
		maxBackups,
		maxAge,
	}
}

// A logger file in diff folder seemed cool!
func (l *LogConfig) LoadLogger() {
	//_, err := os.Create("logs/server.log")
	_, err := os.Create(l.LogFilename)

	if err != nil {
		log.Panic("Wohooo, another error", err.Error())
	}

	logr.SetOutput(&lumberjack.Logger{
		Filename:   l.LogFilename,
		MaxSize:    l.LogMaxSize,
		MaxAge:     l.LogMaxAge,
		MaxBackups: l.LogMaxBackups,
		LocalTime:  false,
		Compress:   false,
	})

	logr.SetLevel(logr.DebugLevel)

	logr.SetFormatter(&logr.JSONFormatter{
		TimestampFormat:   "",
		DisableTimestamp:  false,
		DisableHTMLEscape: false,
		DataKey:           "",
		FieldMap:          nil,
		CallerPrettyfier:  nil,
		PrettyPrint:       true,
	})

}
