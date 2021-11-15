//Package logrus provides a logger implementation based on the github.com/sirupsen/logrus pkg
package logrus

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"

	"github.com/devopsfaith/krakend/config"
	"github.com/sirupsen/logrus"
)

// Namespace is the key to look for extra configuration details
const Namespace = "github_com/devopsfaith/krakend-logrus"

// ErrWrongConfig is the error returned when there is no config under the namespace
var ErrWrongConfig = errors.New("getting the extra config for the krakend-logrus module")

type Level int

const (
	EMERGENCY = iota + 1
	ALERT     = iota
	CRITICAL  = iota
	ERROR     = iota
	WARNING   = iota
	NOTICE    = iota
	INFO      = iota
	DEBUG     = iota
)

// NewLogger returns a krakend logger wrapping a logrus logger
func NewLogger(cfg config.ExtraConfig, ws ...io.Writer) (*Logger, error) {
	logConfig, ok := ConfigGetter(cfg).(Config)
	if !ok {
		return nil, ErrWrongConfig
	}

	level, ok := logLevels[logConfig.Level]
	if !ok {
		return nil, fmt.Errorf("unknown log level: %s", logConfig.Level)
	}

	l := logrus.New()
	setFormatter(l, logConfig)
	setOutput(l, logConfig, ws...)
	l.Level = logrus.DebugLevel

	return &Logger{
		logger: l,
		level:  level,
	}, nil
}

// WrapLogger wraps already configured logrus instance
func WrapLogger(l *logrus.Logger, module string) *Logger {
	return &Logger{
		logger: l,
		level:  l.Level,
	}
}

func setFormatter(l *logrus.Logger, cfg Config) {
	switch {
	case cfg.JSONFormatter != nil:
		l.Formatter = cfg.JSONFormatter
	case cfg.TextFormatter != nil:
		l.Formatter = cfg.TextFormatter
	default:
		l.Formatter = &logrus.TextFormatter{}
	}
}

func setOutput(l *logrus.Logger, cfg Config, ws ...io.Writer) {
	if cfg.StdOut {
		ws = append(ws, os.Stdout)
	}
	if cfg.Syslog {
		// ws = append(ws, b)
	}

	if len(ws) == 1 {
		l.Out = ws[0]
		return
	}
	l.Out = io.MultiWriter(ws...)
}

// ConfigGetter implements the config.ConfigGetter interface
func ConfigGetter(e config.ExtraConfig) interface{} {
	v, ok := e[Namespace]
	if !ok {
		return nil
	}
	cfg := Config{}

	data, err := json.Marshal(v)
	if err != nil {
		return nil
	}
	if json.Unmarshal(data, &cfg); err != nil {
		return nil
	}

	return cfg
}

// Config is the custom config struct containing the params for the logger
type Config struct {
	Level         string                `json:"level"`
	StdOut        bool                  `json:"stdout"`
	Syslog        bool                  `json:"syslog"`
	Module        string                `json:"module"`
	TextFormatter *logrus.TextFormatter `json:"text"`
	JSONFormatter *logrus.JSONFormatter `json:"json"`
}

// Logger is a wrapper over a github.com/sirupsen/logrus logger
type Logger struct {
	logger *logrus.Logger
	level  logrus.Level
}

func (l *Logger) WithField(key string, value interface{}) *Entry {
	return &Entry{l.logger.WithField(key, value)}
}

func (l *Logger) Debug(v ...interface{}) {
	l.logger.Log(DEBUG, v...)
}

func (l *Logger) Info(v ...interface{}) {
	l.logger.Log(INFO, v...)
}

func (l *Logger) Notice(v ...interface{}) {
	l.logger.Log(NOTICE, v...)
}

func (l *Logger) Warning(v ...interface{}) {
	l.logger.Log(WARNING, v...)
}

func (l *Logger) Error(v ...interface{}) {
	l.logger.Log(ERROR, v...)
}

func (l *Logger) Critical(v ...interface{}) {
	l.logger.Log(CRITICAL, v...)
}

func (l *Logger) Alert(v ...interface{}) {
	l.logger.Log(ALERT, v...)
}

func (l *Logger) Emergency(v ...interface{}) {
	l.logger.Log(EMERGENCY, v...)
}

func (l *Logger) NewEntry() *Entry {
	return &Entry{logrus.NewEntry(l.logger)}
}

var logLevels = map[string]logrus.Level{
	"DEBUG":     DEBUG,
	"INFO":      INFO,
	"NOTICE":    NOTICE,
	"WARNING":   WARNING,
	"ERROR":     ERROR,
	"CRITICAL":  CRITICAL,
	"ALERT":     ALERT,
	"EMERGENCY": EMERGENCY,
}
