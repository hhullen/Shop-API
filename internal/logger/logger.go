package logger

import (
	"fmt"
	"io"
	"log"
	"sync"

	"shopapi/internal/supports"
)

const (
	messagesBuffer = 500
	LevelLabel     = "Level"
	infoTag        = "INFO"
	warnTag        = "WARNING"
	errorTag       = "ERROR"
	fatalTag       = "FATAL"
)

type Message struct {
	message string
	args    []any
}

type Logger struct {
	l     log.Logger
	logCh chan Message

	wg sync.WaitGroup
}

func NewLogger(out io.Writer, prefix string) *Logger {
	l := &Logger{
		logCh: make(chan Message, messagesBuffer),
	}

	writer := l.l.Println

	l.l.SetOutput(out)
	l.l.SetPrefix(prefix + " ")
	l.l.SetFlags(log.Ltime + log.Ldate)

	go l.listenCh(l.logCh, writer)

	return l
}

func (l *Logger) Stop() {
	l.wg.Wait()
	close(l.logCh)

}

func (l *Logger) Infof(template string, args ...any) {
	l.wg.Add(1)
	l.send(Message{message: fmt.Sprintf(template, args...), args: []any{LevelLabel, infoTag}})
}

func (l *Logger) Errorf(template string, args ...any) {
	l.wg.Add(1)
	l.send(Message{message: fmt.Sprintf(template, args...), args: []any{LevelLabel, errorTag}})
}

func (l *Logger) Warnf(template string, args ...any) {
	l.wg.Add(1)
	l.send(Message{message: fmt.Sprintf(template, args...), args: []any{LevelLabel, warnTag}})
}

func (l *Logger) Fatalf(template string, args ...any) {
	l.wg.Add(1)
	l.send(Message{message: fmt.Sprintf(template, args...), args: []any{LevelLabel, fatalTag}})
}

func (l *Logger) InfoKV(message string, argsKV ...any) {
	l.wg.Add(1)
	argsKV = append(argsKV, LevelLabel, infoTag)
	l.send(Message{message: message, args: argsKV})
}

func (l *Logger) WarnKV(message string, argsKV ...any) {
	l.wg.Add(1)
	argsKV = append(argsKV, LevelLabel, warnTag)
	l.send(Message{message: message, args: argsKV})
}

func (l *Logger) ErrorKV(message string, argsKV ...any) {
	l.wg.Add(1)
	argsKV = append(argsKV, LevelLabel, errorTag)
	l.send(Message{message: message, args: argsKV})
}

func (l *Logger) FatalKV(message string, argsKV ...any) {
	l.wg.Add(1)
	argsKV = append(argsKV, LevelLabel, fatalTag)
	l.send(Message{message: message, args: argsKV})
}

func (l *Logger) listenCh(ch chan Message, fn func(...any)) {
	for msg := range ch {
		bytes, err := supports.MakeKVMessagesJSON(msg.args...)
		if err != nil {
			fn("failed making log", err.Error())
		} else {
			fn(msg.message, string(bytes))
		}
		l.wg.Done()
	}
}

func (l *Logger) send(msg Message) {
	defer func() {
		if p := recover(); p != nil {
			fmt.Println(p)
		}
	}()

	l.logCh <- msg
}
