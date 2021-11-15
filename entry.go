package logrus

import (
	"github.com/sirupsen/logrus"
)

// An entry is the final or intermediate Logrus logging entry. It contains all
// the fields passed with WithField{,s}. It's finally logged when Trace, Debug,
// Info, Warn, Error, Fatal or Panic is called on it. These objects can be
// reused and passed around as much as you wish to avoid field duplication.
type Entry struct {
	*logrus.Entry
}

func (e *Entry) Debug(v ...interface{}) {
	e.Log(logrus.DebugLevel, v...)
}

func (e *Entry) Info(v ...interface{}) {
	e.Log(logrus.InfoLevel, v...)
}

func (e *Entry) Warning(v ...interface{}) {
	e.Log(logrus.WarnLevel, v...)
}

func (e *Entry) Error(v ...interface{}) {
	e.Log(logrus.ErrorLevel, v...)
}

func (e *Entry) Critical(v ...interface{}) {
	e.Log(logrus.ErrorLevel, v...)
}

func (e *Entry) Fatal(v ...interface{}) {
	e.Log(logrus.FatalLevel, v...)
}

func (entry *Entry) WithField(key string, value interface{}) *Entry {
	tmp := &Entry{entry.Entry.WithField(key, value)}
	return tmp
}