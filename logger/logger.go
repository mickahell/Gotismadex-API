package logger

import (
	"fmt"
	"os"
	"time"

	"github.com/mickahell/Gotismadex/helpers"
)

type Logger struct {
	loggerLevel int
	drawing     string
	timestamp   time.Time
	level       string
	endpoint    string
	message     string
	detailes    error
}

const DrawColor = "\033[96;1m"
const NoneColor = "\033[0m"

var mylog Logger

func (log *Logger) LogInit() {
	log.loggerLevel = helpers.TheAppConfig().Loglevel
}

func (log *Logger) LogDraw(drawing string) {
	log.drawing = drawing

	fmt.Fprintln(os.Stdout, DrawColor, string(log.drawing), NoneColor)
}

func (log *Logger) LogInfo(endpoint string, message string) {
	log.timestamp = time.Now()
	log.level = "INFO"
	log.endpoint = endpoint
	log.message = message

	if log.loggerLevel == 0 {
		fmt.Fprintln(
			os.Stdout,
			"timestamp=\""+log.timestamp.Format(
				"2006-01-02 15:04:05",
			)+"\" level="+log.level+" endpoint="+log.endpoint+" message=\""+log.message+"\"",
		)
	}
}

func (log *Logger) LogWarning(endpoint string, message string, detailes error) {
	log.timestamp = time.Now()
	log.level = "WARNING"
	log.endpoint = endpoint
	log.message = message
	log.detailes = detailes

	if log.loggerLevel <= 1 {
		fmt.Fprintln(
			os.Stderr,
			"timestamp=\""+log.timestamp.Format(
				"2006-01-02 15:04:05",
			)+"\" level="+log.level+" endpoint="+log.endpoint+" message=\""+log.message+
				"\" detailes=\""+log.detailes.Error()+"\"",
		)

	}
}

func (log *Logger) LogError(endpoint string, message string, detailes error) {
	log.timestamp = time.Now()
	log.level = "ERROR"
	log.endpoint = endpoint
	log.message = message
	log.detailes = detailes

	fmt.Fprintln(
		os.Stderr,
		"timestamp=\""+log.timestamp.Format(
			"2006-01-02 15:04:05",
		)+"\" level="+log.level+" endpoint="+log.endpoint+" message=\""+log.message+
			"\" detailes=\""+log.detailes.Error()+"\"",
	)
}

func (log *Logger) LogCritical(endpoint string, message string, detailes error, exit ...bool) {
	log.timestamp = time.Now()
	log.level = "CRITICAL"
	log.endpoint = endpoint
	log.message = message
	log.detailes = detailes

	fmt.Fprintln(
		os.Stderr,
		"timestamp=\""+log.timestamp.Format(
			"2006-01-02 15:04:05",
		)+"\" level="+log.level+" endpoint="+log.endpoint+" message=\""+log.message+
			"\" detailes=\""+log.detailes.Error()+"\"",
	)
	if len(exit) == 0 {
		os.Exit(1)
	}
}

func GetLogger() *Logger {
	return &mylog
}
