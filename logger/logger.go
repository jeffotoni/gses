package logger

import (
	"log"
	"os"
)

// escreve para o stdout.
func New() *log.Logger {
	return log.New(os.Stdout, "", log.LstdFlags)
}

func NewErr() *log.Logger {
	return log.New(os.Stderr, "", log.LstdFlags)
}

func Printf(format string, v ...interface{}) {
	logg := New()
	logg.Printf(format, v...)
}

func Printferr(format string, v ...interface{}) {
	logg := NewErr()
	logg.Printf(format, v...)
}
