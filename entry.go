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
	e.Log(DEBUG, v...)
}

func (e *Entry) Info(v ...interface{}) {
	e.Log(INFO, v...)
}

func (e *Entry) Notice(v ...interface{}) {
	e.Log(NOTICE, v...)
}

func (e *Entry) Warning(v ...interface{}) {
	e.Log(WARNING, v...)
}

func (e *Entry) Error(v ...interface{}) {
	e.Log(ERROR, v...)
}

func (e *Entry) Critical(v ...interface{}) {
	e.Log(CRITICAL, v...)
}

func (e *Entry) Alert(v ...interface{}) {
	e.Log(ALERT, v...)
}

func (e *Entry) Emergency(v ...interface{}) {
	e.Log(EMERGENCY, v...)
}

func (entry *Entry) WithField(key string, value interface{}) *Entry {
	tmp := &Entry{entry.Entry.WithField(key, value)}
	return tmp
}